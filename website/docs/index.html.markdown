---
layout: "routeros"
page_title: "Provider: RouterOS"
sidebar_current: "docs-routeos-index"
description: |-
  The RouterOS provider allows Terraform to read from, write to, and configure devices
  and virtual machines powered by RouterOS.
---

# RouterOS Provider

The RouterOS provider allows Terraform to read from, write to, and configure
devices and virtual machines powered by [RouterOS](https://mikrotik.com/software).

## Provider Arguments

The provider configuration block accepts the following arguments.
In most cases it is recommended to set them via the indicated environment
variables in order to keep credential information out of the configuration.

* `address` - (Optional) Hostname or IP address of the RouterOS device to be managed.
  May be set via the `ROUTEROS_ADDRESS` environment variable.

* `port` - (Optional) Management port to be used.
  May be set via the `ROUTEROS_PORT` environment variable.

* `username` - (Optional) Username to be used.
  May be set via the `ROUTEROS_USERNAME` environment variable.

* `password` - (Optional) Passsword to be used.
  May be set via the `ROUTEROS_PASSWORD` environment variable.

* `tls` - (Optional) Whether to enable transport layer security or not.
  May be set via the `ROUTEROS_TLS` environment variable.

## Example Usage

```hcl
provider "routeros" {
  # It is strongly recommended to configure this provider through the
  # environment variables described above, so that each user can have
  # separate credentials set in the environment.
  #
  # This will default to using $ROUTEROS_ADDRESS
  # But can be set explicitly
  # address = "192.168.88.1"
}

data "routeros_interface" "ether1" {
  name = "ether1"
}
```
