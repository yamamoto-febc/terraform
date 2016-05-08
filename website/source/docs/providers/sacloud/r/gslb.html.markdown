---
layout: "sakuracloud"
page_title: "SakuraCloud: sacloud_gslb"
sidebar_current: "docs-sacloud-resource-gslb"
description: |-
  Provides a SakuraCloud GSLB resource. This can be used to create, modify, and delete GSLB.
---

# sakuracloud\_gslb

Provides a SakuraCloud GSLB(Global Site Load Balancing) resource. This can be used to create,
modify, and delete GSLB. 

## Example Usage

```
# Create a new GSLB and add two target server.
resource "sakuracloud_gslb" "mygslb" {
    name = "gslb_from_terraform"
    health_check = {
        protocol = "http"
        delay_loop = 10
        host_header = "example.com"
        path = "/"
        status = "200"
    }
    description = "GSLB from terraform for SAKURA CLOUD"
    tags = ["hoge1" , "hoge2" ]
    servers = {
      ipaddress = "192.0.2.1"
    }
    servers = {
      ipaddress = "192.0.2.2"
    }

}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of GSLB.
* `health_check` - (Required) The health_check rule of GSLB.
  * `protocol` - (Required) The protocol to use for health check. Must be in [`http`,`https`,`tcp`,`ping`]
  * `dalay_loop` - (Optional) The delay_loop of health check. Must be between `10` and `60`. default is `10`
  * `host_header` - (Only protocol is `http` or `https`) The host_header to use for health check.
  * `path` - (Only when protocol is `http` or `https`) The request path to use for health check.
  * `status` - (Only when protocol is `http` or `https`) The response code of health check request.
  * `port` - (Only when protocol is `tcp`) The port number to use for health check.
* `weighted` - (Optional)The flag of enabling to weighted balancing. default `false`
* `description` - (Optional) The description of GSLB.
* `tags` - (Required) The tags of GSLB.
* `servers` (Optional) The target servers of GSLB.
  * `ipaddress` - The IPAddress of target server.
  * `enabled` - The flag of enabling to target server.
  * `weight` - (Only when `weighted` is true)The weight of target server.


## Attributes Reference

The following attributes are exported:

* `id` - The ID of the GSLB.
* `name`- The name of GSLB.
* `health_check` - The health_check rule of GSLB.
  * `protocol` - The protocol to use for health check.
  * `dalay_loop` - The delay_loop of health check.
  * `host_header` - The host_header to use for health check.
  * `path` - The request path to use for health check.
  * `status` - The response code of health check request.
  * `port` - The port number to use for health check.
* `weighted` - The flag of enabling to weighted balancing.
* `description` - The description of GSLB.
* `tags` - The tags of GSLB.
* `servers` - The target servers of GSLB.
  * `ipaddress` - The IPAddress of target server.
  * `enabled` - The flag of enabling to target server.
  * `weight` - The weight of target server.
* `FQDN` - The FQDN of GSLB.