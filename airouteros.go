package airouteros

import (
	"context"

	"github.com/go-routeros/routeros/v3"
)

type RouterOS struct {
	client *routeros.Client
	option *Options
}

func NewRouterOS(opt *Options) *RouterOS {
	return &RouterOS{
		option: opt,
	}
}

// Connect connects to the router.
func (r *RouterOS) Connect() error {
	var err error
	if r.option.Timeout > 0 {
		r.client, err = routeros.DialTimeout(r.option.Address, r.option.Username, r.option.Password, r.option.Timeout)
		return err
	}
	r.client, err = routeros.Dial(r.option.Address, r.option.Username, r.option.Password)
	return err
}

// Close the connection to the router.
func (r *RouterOS) Close() {
	r.client.Close()
}

// Async starts asynchronous mode and returns immediately..
func (r *RouterOS) Async() {
	r.client.Async()
}

// Run runs a sentence on the router.
func (b *RouterOS) CallContext(ctx context.Context, sentence ...string) (*routeros.Reply, error) {
	return b.client.RunContext(ctx, sentence...)
}

// GetQuery returns the query of the builder.
func (b *RouterOS) GetPath() string {
	return `/`
}
