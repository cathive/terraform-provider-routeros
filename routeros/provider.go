package routeros

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

const (
	// DefaultAddress specifies the default hostname / ip address
	// of the RouterOS target.
	DefaultAddress = "192.168.88.1"
	// DefaultPortInsecure specifies the default non-TLS port.
	DefaultPortInsecure = "8728"
	// DefaultPortSecure specifies the default TLS port.
	DefaultPortSecure = "8789"
)

// Provider creates a new RouterOS terraform provider instance.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Description: "Hostname or IP address of the RouterOS device to be managed",
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("ROUTEROS_ADDRESS", DefaultAddress),
				Optional:    true,
			},
			"port": {
				Description: "Port where the RouterOS device listens for incoming connections",
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("ROUTEROS_PORT", DefaultPortInsecure),
				Optional:    true,
				Default:     "",
			},
			"username": {
				Description: "Username to be used for accessing the RouterOS device",
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("ROUTEROS_USERNAME", "admin"),
				Optional:    true,
				Default:     "admin",
			},
			"password": {
				Description: "Password to be used for accessing the RouterOS device",
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("ROUTEROS_PASSWORD", ""),
				Optional:    true,
				Default:     "",
			},
			"tls": {
				Description: "Whether to use transport layer security or not",
				Type:        schema.TypeBool,
				DefaultFunc: schema.EnvDefaultFunc("ROUTEROS_TLS", false),
				Optional:    true,
				Default:     false,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"routeros_interface":  DataInterface(),
			"routeros_ip_address": DataIpAddress(),
			"routeros_user_group": DataUserGroup(),
		},
		ResourcesMap:  map[string]*schema.Resource{},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	address := d.Get("address").(string)
	if address == "" {
		address = DefaultAddress
	}
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	tls := d.Get("tls").(bool)
	port := d.Get("port").(string)
	if port == "" {
		if tls {
			port = DefaultPortSecure
		} else {
			port = DefaultPortInsecure
		}
	}

	return NewConfig(address, port, username, password, tls)
}
