---
layout: "routeros"
page_title: "RouterOS: routeros_interface data source"
sidebar_current: "docs-routeros-datasource-interface"
description: |-
  Reads information of a given interface.
---

# routeros\_interface

This is a data source which can be used to retrieve the configuration parameters
of any given named interface.

## Example usage

```hcl
data "routeros_interface" "ether1" {
  name = "ether1"
}
```

## Argument Reference

Only one argument is supported:

* `name` - (Required) Name of the interface to be examined.

## Attributes


* `comment` - User comment provided for the interface.

* `type` - Type of the interface ("ether", "bridge",...).

* `mtu` - Maximum transmission unit of the interfae in bytes or "auto".

* `l2mtu` - The maximum size of the frame without MAC header that can be sent by this
  interface.

* `default_name` - The default name of the interface.

* `mac_address` - MAC address of the interface.
