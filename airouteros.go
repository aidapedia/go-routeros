package routeros

import (
	"github.com/go-routeros/routeros/v3"
)

type RouterOS struct {
	client *routeros.Client
	option *Option
}

func NewRouterOS(address, username, password string) *RouterOS {
	return &RouterOS{
		option: NewOption(address, username, password),
	}
}

// Connect connects to the router.
func (r *RouterOS) Connect() error {
	var err error
	if r.option.GetTimeout() > 0 {
		r.client, err = routeros.DialTimeout(r.option.GetAddress(), r.option.GetUsername(), r.option.GetPassword(), r.option.GetTimeout())
		return err
	}
	r.client, err = routeros.Dial(r.option.GetAddress(), r.option.GetUsername(), r.option.GetPassword())
	return err
}

// Close closes the connection to the router.
func (r *RouterOS) Close() {
	r.client.Close()
}

// Close closes the connection to the router.
func (r *RouterOS) Async() {
	r.client.Async()
}

// Run runs a sentence on the router.
func (b *RouterOS) Call(sentence ...string) (*routeros.Reply, error) {
	return b.client.Run(sentence...)
}

// GetQuery returns the query of the builder.
func (b *RouterOS) GetPath() string {
	return `/`
}
