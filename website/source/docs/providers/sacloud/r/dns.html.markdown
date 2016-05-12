---
layout: "sakuracloud"
page_title: "SakuraCloud: sacloud_dns"
sidebar_current: "docs-sacloud-resource-dns"
description: |-
  Provides a SakuraCloud DNS resource. This can be used to create, modify, and delete DNS records.
---

# sakuracloud\_dns

Provides a SakuraCloud DNS resource. This can be used to create,
modify, and delete DNS records. 

## Example Usage

```
# Create a new DNS zone and add two A records.
resource "sakuracloud_dns" "dns" {
    zone = "example.com"
    records = {
        name = "test1"
        type = "A"
        value = "192.168.0.1"
    } 
    records = {
        name = "test2"
        type = "A"
        value = "192.168.0.2"
    } 
}
```

## Argument Reference

The following arguments are supported:

* `zone` - (Required) The DNS target zone name.
* `description` - (Required) The description of DNS.
* `tags` - (Required) The tags of DNS.
* `records` - (Optional) The records of target zone.
  * `name` - (Required) The name of the record.
  * `type` - (Required) The type of the record.
  * `value` - (Required) The value of the record.
  * `ttl` - (Optional) The TTL of the record . default `3600`.
  * `priority` - (Optional) (Only type is MX) The priority of the record.



## Attributes Reference

The following attributes are exported:

* `id` - The ID of the DNS.
* `zone`- The DNS target zone name.
* `dns_servers` - The name servers of the target zone.
* `description` - The description of target zone.
* `tags` - The description of target zone.
* `records` - The records of target zone.
  * `name` - The name of the record.
  * `type` - The type of the record.
  * `value` - The value of the record.
  * `ttl` - The TTL of the record.
  * `priority` - (Only type is MX) The priority of the record.
