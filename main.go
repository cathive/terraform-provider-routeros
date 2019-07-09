//go:generate ./.bin/terraform-routeros-binding-generator ./generator/ ./routeros/

package main

import (
	"github.com/cathive/terraform-provider-routeros/routeros"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: routeros.Provider,
	})
}
