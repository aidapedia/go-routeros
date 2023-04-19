package routeros

import (
	"gopkg.in/routeros.v2"
)

type RouterOSBuilder struct {
	client *routeros.Client
	option *Option
}

func NewRouterOSBuilder(address, username, password string) *RouterOSBuilder {
	return &RouterOSBuilder{
		option: NewOption(address, username, password),
	}
}

// Connect connects to the router.
func (r *RouterOSBuilder) Connect() error {
	var err error
	if r.option.GetTimeout() > 0 {
		r.client, err = routeros.DialTimeout(r.option.GetAddress(), r.option.GetUsername(), r.option.GetPassword(), r.option.GetTimeout())
		return err
	}
	r.client, err = routeros.Dial(r.option.GetAddress(), r.option.GetUsername(), r.option.GetPassword())
	return err
}

// Close closes the connection to the router.
func (r *RouterOSBuilder) Close() {
	r.client.Close()
}

// Close closes the connection to the router.
func (r *RouterOSBuilder) Async() {
	r.client.Async()
}

// Run runs a sentence on the router.
func (b *RouterOSBuilder) Run(sentence ...string) (*routeros.Reply, error) {
	return b.client.Run(sentence...)
}

// GetQuery returns the query of the builder.
func (b *RouterOSBuilder) GetQuery() string {
	return `/`
}
