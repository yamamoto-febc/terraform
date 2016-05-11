---
layout: "sakuracloud"
page_title: "SakuraCloud: sacloud_disk"
sidebar_current: "docs-sacloud-resource-disk"
description: |-
  Provides a SakuraCloud Disk resource. This can be used to create, and delete Disk .
---

# sakuracloud\_disk

Provides a SakuraCloud Disk resource. This can be used to create, 
and delete Disk .

## Example Usage

```
# Create a new Disk with source archive named "Ubuntu"
resource "sakuracloud_disk" "mydisk"{
    name = "mydisk"
    size = 20480
    source_archive_name = "Ubuntu Server 14.04"
    description = "Disk from terraform for SAKURA CLOUD"
    tags = ["hoge1" , "hoge2"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) The name of the disk
* `plan` - (Optional) The plan of the disk
* `connection` - (Optional) The connection of the disk
* `size` - (Optional) The size(MB) of the disk
* `source_archive_id` - (Optional) The id of source archive.
  Conflicts with `source_archive_name , source_disk_id , source_disk_name`
* `source_archive_name` - (Optional) The name of source archive.
  Conflicts with `source_archive_id , source_disk_id , source_disk_name`
* `source_disk_id` - (Optional) The id of source disk.
  Conflicts with `source_archive_id , source_archive_name , source_disk_name`
* `source_disk_name` - (Optional) The name of source disk.
  Conflicts with `source_archive_id , source_archive_name , source_disk_id`
* `description` - (Optional) The description of the disk
* `tags` - (Optional) The tags of the disk
* `zone` - (Optional) The zone of to create disk
* `password` - (Optional) The password of the disk
* `ssh_key_ids` - (Optional) The ID list of SSHKey.
* `disable_pw_auth` - (Optional) The flag that to disable SSH login with password authentication / challenge-response. default id `false`.
* `note_ids` - (Optional) The ID list of Note.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the disk
* `name`- The name of the disk
* `plan` - The plan of the disk
* `connection` - The connection of the disk
* `size` - The size(MB) of the disk
* `source_archive_id` - The id of source archive
* `source_archive_name` - The name of source archive
* `source_disk_id` - The id of source disk
* `source_disk_name` - The name of source disk
* `description` - The description of the disk
* `tags` - The tags of the disk
* `zone` - The zone of the disk
* `password` - The password of the disk/
* `ssh_key_ids` The ID list of SSHKey.
* `disable_pw_auth` - The flag that to disable SSH login with password authentication / challenge-response.
* `note_ids` - The ID list of Note.
