---
layout: "sakuracloud"
page_title: "SakuraCloud: sacloud_gslb"
sidebar_current: "docs-sacloud-resource-gslb"
description: |-
  Provides a SakuraCloud GSLB resource. This can be used to create, modify, and delete GSLB.
---

# sakuracloud\_gslb

Provides a SakuraCloud GSLB(Global Site Load Balancer) resource. This can be used to create,
modify, and delete GSLB. 

## Example Usage

```
# Create a new GSLB and add two target server.
resource "sakuracloud_gslb" "mygslb" {
    name = "gslb_from_terraform"
    health_check = {
        protocol = "tcp"
        delay_loop = 10
        port = 80
#        host_header = "libsacloud.com"
#        path = "/index.html"
#        status = "200"
    }
    description = "GSLB from terraform for SAKURA CLOUD"
    tags = ["hoge1" , "hoge2" ]
    servers = {
      ipaddress = "133.242.1.1"
    }
    servers = {
      ipaddress = "133.242.1.2"
    }

}
```

## Argument Reference

The following arguments are supported:

* `zone` - (Required) The DNS target zone name
* `description` - (Required) The region to start in
* `tags` - (Required) The instance size to start
* `records` - (Optional) The records of target zone
  * `name` - (Required) The name of the record
  * `type` - (Required) The type of the record
  * `value` - (Required) The value of the record
  * `ttl` - (Optional) The TTL of the record . default `3600`
  * `priority` - (Optional) (Only type is MX) The priority of the record



## Attributes Reference

The following attributes are exported:

* `id` - The ID of the DNS
* `name`- The domain name of the target zone
* `FQDN` - The name servers of the target zone
* `health_check` - The records of target zone
  * `protocol` - The name of the record
  * `dalay_loop` - The type of the record
  * `host_header` - The value of the record
  * `path` - The TTL of the record
  * `status` - (Only type is MX) The priority of the record
  * `port` - (Only type is MX) The priority of the record
* `weighted` - The name servers of the target zone
* `description` - The description of target zone
* `tags` - The description of target zone
* `servers` - The records of target zone
  * `ipaddress` - The name of the record
  * `enabled` - The type of the record
  * `weight` - The value of the record
