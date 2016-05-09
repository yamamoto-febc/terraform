---
layout: "sakuracloud"
page_title: "Provider: SakuraCloud"
sidebar_current: "docs-sacloud-index"
description: |-
  The SakuraCloud(sacloud) provider is used to interact with the resources supported by SakuraCloud. The provider needs to be configured with the proper credentials before it can be used.
---

# Sakura Cloud Provider

The SakuraCloud(sacloud) provider is used to interact with the
resources supported by SakuraCloud. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```
# Set the variable value in *.tfvars file
# or using -var="sacloud_token=..." CLI option
# or using environment value
variable "sacloud_token" {}
variable "sacloud_secret" {}
variable "sacloud_zone" {}

# Configure the SakuraCloud Provider
provider "sacloud" {
    token = "${var.sacloud_token}"
    secret = "${var.sacloud_secret}"
    zone = "${var.sacloud_zone}"
}

# Create a web server
resource "sacloud_server" "web" {
    ...
}
```

## Argument Reference

The following arguments are supported:

* `token` - (Required) This is the sacloud APIKey(token). This can also be specified
  with the `SAKURACLOUD_ACCESS_TOKEN` shell environment variable.

* `secret` - (Required) This is the sacloud API(secret). This can also be specified
  with the `SAKURACLOUD_ACCESS_TOKEN_SECRET` shell environment variable.
  
* `zone` - (Required) This is the sacloud zone. This can also be specified
  with the `SAKURACLOUD_ZONE` shell environment variable.