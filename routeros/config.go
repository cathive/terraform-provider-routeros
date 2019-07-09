package routeros

import (
	"crypto/tls"
	"fmt"

	"github.com/go-routeros/routeros"
)

// Config provides a way to wrap/encapsulate all the stuff that is necessary
// to interact with the RouterOS API.
type Config interface {
	Client() (Client, error)
}

type config struct {
	address  string
	port     string
	username string
	password string
	tls      bool
	client   Client
}

func (cfg *config) Client() (Client, error) {
	if cfg.client == nil {
		var client *routeros.Client
		var err error
		address := fmt.Sprintf("%s:%s", cfg.address, cfg.port)

		// Dial in to the RouterOS device.
		if cfg.tls {
			client, err = routeros.DialTLS(address, cfg.username, cfg.password, &tls.Config{
				// TODO a bit more configurability would be nice.
			})
		} else {
			client, err = routeros.Dial(address, cfg.username, cfg.password)
		}

		if err != nil {
			return nil, fmt.Errorf(`error while connecting to RouterOS device "%s:%s": %v`, cfg.address, cfg.port, err)
		}

		ready := make(chan interface{}, 1)
		ready <- nil
		cfg.client = &routerOSClient{
			ready:   ready,
			wrapped: *client,
		}
	}
	return cfg.client, nil
}

// NewConfig creates a new configuration structure to be used provider-internally.
func NewConfig(address string, port string, username string, password string, tlsEnabled bool) (Config, error) {
	cfg := config{
		address:  address,
		port:     port,
		username: username,
		password: password,
		tls:      tlsEnabled,
	}

	return &cfg, nil
}
