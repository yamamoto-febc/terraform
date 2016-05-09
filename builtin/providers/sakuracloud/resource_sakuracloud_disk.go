package sakuracloud

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yamamoto-febc/libsacloud/api"
	"github.com/yamamoto-febc/libsacloud/sacloud"
	"time"
)

func resourceSakuraCloudDisk() *schema.Resource {
	return &schema.Resource{
		Create: resourceSakuraCloudDiskCreate,
		Read:   resourceSakuraCloudDiskRead,
		//Update: resourceSakuraCloudDiskUpdate, //!Not Support!
		Delete: resourceSakuraCloudDiskDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"plan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  sacloud.DiskPlanSSD.ID.String(),
				ValidateFunc: validateStringInWord([]string{
					sacloud.DiskPlanSSD.ID.String(),
					sacloud.DiskPlanHDD.ID.String(),
				}),
			},
			"connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  sacloud.DiskConnectionVirtio,
				ValidateFunc: validateStringInWord([]string{
					fmt.Sprintf("%s", sacloud.DiskConnectionVirtio),
					fmt.Sprintf("%s", sacloud.DiskConnectionIDE),
				}),
			},
			"source_archive_id": &schema.Schema{
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				ConflictsWith: []string{"source_archive_name", "source_disk_name", "source_disk_id"},
			},
			"source_archive_name": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"source_archive_id", "source_disk_name", "source_disk_id"},
			},
			"source_disk_id": &schema.Schema{
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				ConflictsWith: []string{"source_disk_name", "source_archive_name", "source_archive_id"},
			},
			"source_disk_name": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"source_disk_id", "source_archive_name", "source_archive_id"},
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  20480,
			},
			"server_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
			"zone": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "target SakuraCloud zone",
				ValidateFunc: validateStringInWord([]string{"is1a", "is1b", "tk1a", "tk1v"}),
			},
		},
	}
}

func resourceSakuraCloudDiskCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)

	zone, ok := d.GetOk("zone")
	if ok {
		originalZone := client.Zone
		client.Zone = zone
		defer func() { client.Zone = originalZone }()
	}

	opts := client.Disk.New()

	//name
	opts.Name = d.Get("name").(string)

	//plan
	opts.Plan.SetIDByString(d.Get("plan").(string))

	//connection
	opts.Connection = sacloud.EDiskConnection(d.Get("connection").(string))

	//source archieve/disk
	archiveID, ok := d.GetOk("source_archive_id")
	if ok {
		opts.SetSourceArchive(archiveID.(string))
	}
	//search archive
	archiveName, ok := d.GetOk("source_archive_name")
	if ok {
		res, err := client.Archive.WithNameLike(archiveName.(string)).Include("ID").Limit(1).Find()
		if err != nil {
			return fmt.Errorf("Failed to create SakuraCloud Disk resource - Source archive not found: %s", err)
		}
		opts.SetSourceArchive(res.Archives[0].ID)
	}

	diskID, ok := d.GetOk("source_disk_id")
	if ok {
		opts.SetSourceDisk(diskID.(string))
	}
	//search archive
	diskName, ok := d.GetOk("source_disk_name")
	if ok {
		res, err := client.Disk.WithNameLike(diskName.(string)).Include("ID").Limit(1).Find()
		if err != nil {
			return fmt.Errorf("Failed to create SakuraCloud Disk resource - Source disk not found: %s", err)
		}
		opts.SetSourceDisk(res.Disks[0].ID)
	}

	//size
	opts.SizeMB = d.Get("size").(int)

	//description
	opts.Description = d.Get("description").(string)

	//tags
	rawTags := d.Get("tags").([]interface{})
	if rawTags != nil {
		opts.Tags = expandStringList(rawTags)
	}

	disk, err := client.Disk.Create(opts)
	if err != nil {
		return fmt.Errorf("Failed to create SakuraCloud Disk resource: %s", err)
	}

	err = client.Disk.SleepWhileCopying(disk.ID, 10*time.Minute)
	if err != nil {
		return fmt.Errorf("Failed to create SakuraCloud Disk resource: %s", err)
	}

	//server_id
	server_id, ok := d.GetOk("server_id")
	if ok {
		_, err = client.Disk.ConnectToServer(disk.ID, server_id.(string))

		if err != nil {
			return fmt.Errorf("Failed to connect SakuraCloud Disk resource: %s", err)
		}
	}

	d.SetId(disk.ID)
	return resourceSakuraCloudDiskRead(d, meta)
}

func resourceSakuraCloudDiskRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	zone, ok := d.GetOk("zone")
	if ok {
		originalZone := client.Zone
		client.Zone = zone
		defer func() { client.Zone = originalZone }()
	}

	disk, err := client.Disk.Read(d.Id())
	if err != nil {
		return fmt.Errorf("Couldn't find SakuraCloud Disk resource: %s", err)
	}

	d.Set("name", disk.Name)
	d.Set("plan", disk.Plan.ID.String())
	d.Set("connection", fmt.Sprintf("%s", disk.Connection))
	//if disk.SourceArchive != nil {
	//	d.Set("source_arvhice_id", disk.SourceArchive.ID)
	//}
	//if disk.SourceDisk != nil {
	//	d.Set("source_disk_id", disk.SourceDisk.ID)
	//}
	d.Set("size", disk.SizeMB)
	d.Set("description", disk.Description)
	d.Set("tags", disk.Tags)

	if disk.Server != nil {
		d.Set("server_id", disk.Server.ID)
	}

	return nil
}

func resourceSakuraCloudDiskDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	zone, ok := d.GetOk("zone")
	if ok {
		originalZone := client.Zone
		client.Zone = zone
		defer func() { client.Zone = originalZone }()
	}

	_, err := client.Disk.Delete(d.Id())

	if err != nil {
		return fmt.Errorf("Error deleting SakuraCloud SimpleMonitor resource: %s", err)
	}

	return nil
}
