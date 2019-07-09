//go:generate ../.bin/terraform-routeros-binding-generator ../generator ../api .

package routeros

import wrapped "github.com/go-routeros/routeros"

type Client interface {
	Run(sentence ...string) (*wrapped.Reply, error)
	Close()
}

// routerOSClient provides a *synchronous* version of the original
// Router OS client. Since the original implementation is *not* thread-safe,
// we make sure that invocations of `Run` cannot happen simultaniously.
type routerOSClient struct {
	ready   chan interface{}
	wrapped wrapped.Client
}

func (r *routerOSClient) Run(sentence ...string) (*wrapped.Reply, error) {

	// Read/write ready "token" to ensure synchronous access of non-thread save API
	<-r.ready
	defer func() {
		r.ready <- nil
	}()

	return r.wrapped.Run(sentence...)
}

func (r *routerOSClient) Close() {
	r.wrapped.Close()
}
