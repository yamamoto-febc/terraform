---
layout: "sakuracloud"
page_title: "SakuraCloud: sacloud_server"
sidebar_current: "docs-sacloud-resource-server"
description: |-
  Provides a SakuraCloud Server resource. This can be used to create, modify, and delete Disk records.
---

# sakuracloud\_server

Provides a SakuraCloud Server resource. This can be used to create, modify,
and delete Server.

## Example Usage

```
# Create a new Server"
resource "sakuracloud_server" "myserver" {
    name = "myserver"
    disks = ["${sakuracloud_disk.mydisk.id}"]
    switched_interfaces = [""]
    description = "Server from TerraForm for SAKURA CLOUD"
    tags = ["@virtio-net-pci"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the server
* `disks` - (Required) The ID list of the disk to connect server
* `core` - (Optional) The number of CPU core. default `1`
* `memory` - (Optional) The size of memory(MB). default `1024`
* `shared_interface` - (Optional) The flag of to create a NIC to connect to a shared segment.
* `switched_interfaces` - (Optional) The ID list of to create a NIC to connect to switch.
   If `""` is specified , it creates a NIC unconnected.
* `description` - (Optional) The description of the server
* `tags` - (Optional) The tags of the server
* `zone` - (Optional) The zone of to create server



## Attributes Reference

The following attributes are exported:

* `id` - The ID of the server.
* `name` - The name of the server.
* `disks`- The ID list of the disks.
* `core` - The number of the CPU core.
* `memory` - The size(MB) of the memory.
* `shared_interface` - The flag of has NIC to connect to a shared segment.
* `switched_interfaces` - The ID list of the connected switch.
* `description` - The description of the server.
* `tags` - The tags of the server.
* `zone` - The zone of the server.
