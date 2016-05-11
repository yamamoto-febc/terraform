---
layout: "sakuracloud"
page_title: "Provider: SakuraCloud"
sidebar_current: "docs-sacloud-index"
description: |-
  The SakuraCloud(sacloud) provider is used to interact with the resources supported by SakuraCloud. The provider needs to be configured with the proper credentials before it can be used.
---

# Sakura Cloud Provider

The SakuraCloud provider is used to interact with the
resources supported by SakuraCloud. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```

# Configure the SakuraCloud Provider
provider "sakuracloud" {
    token = "your token"    
    secret = "your secret"
    zone = "target zone"
}

# Create a web server
resource "sakuracloud_server" "web" {
    ...
}
```

## Argument Reference

The following arguments are supported:

* `token` - (Required) This is the SakuraCloud API token. This can also be specified
  with the `SAKURACLOUD_ACCESS_TOKEN` shell environment variable.

* `secret` - (Required) This is the SakuraCloud API secret. This can also be specified
  with the `SAKURACLOUD_ACCESS_TOKEN_SECRET` shell environment variable.
  
* `zone` - (Required) This is the SakuraCloud zone. This can also be specified
  with the `SAKURACLOUD_ZONE` shell environment variable.