---
layout: "sakuracloud"
page_title: "SakuraCloud: sacloud_ssh_key"
sidebar_current: "docs-sacloud-resource-ssh-key"
description: |-
  Provides a SakuraCloud SSHKey resource. This can be used to create, modify, and delete SSHKey.
---

# sakuracloud\_ssh\_key

Provides a SakuraCloud SSHKey resource. This can be used to create, modify,
and delete SSHKey.

## Example Usage

```
resource "sakuracloud_ssh_key" "mykey" {
    name = "mykey"
    public_key = "ssh-rsa XXXXXXXXX....."
    # or
    #public_key = "${file("./id_rsa.pub")}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the SSHKey.
* `public_key` - (Required) The value of the SSHKey.
* `description` - (Optional) The description of the SSHKey.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the SSHKey.
* `name`- The name of the SSHKey.
* `public_key` - The value of the SSHKey.
* `description` - The description of the SSHKey.
* `fingerprint` - The FingerPrint of the SSHKey.
