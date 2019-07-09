# Terraform Provider for RouterOS

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x
- [Go](https://golang.org/doc/install) 1.12.x (to build the provider plugin)

## Building the Provider

Clone repository outside of your GOPATH.

```sh
# Clone the provider plugin sources
$ mkdir -p ~/Projects; cd ~/Projects
$ git clone git@github.com:cathive/terraform-provider-routeros.git
```

Enter the provider directory and build the provider

```sh
# Build the provider plugin
$ cd ~/Projects/terraform-provider-routeros
$ make build
# Now make the plugin locally available
# OS represents your operating system (e.g. windows, darwin, linux) and
# ARCH represents your operating system's architecture (most likely: amd64)
$ cp $GOPATH/bin/terraform-provider-routeros ~/.terraform.d/plugins/${OS}_${ARCH}/
```

## Using the provider

If you're building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin).
After placing it into your plugins directory,  run `terraform init` to initialize it.
