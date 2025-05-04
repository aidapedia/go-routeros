package routeros

import (
	"context"
	"fmt"

	"github.com/aidapedia/go-routeros/types"
	"github.com/go-routeros/routeros/v3"
)

type RouterOS struct {
	client *routeros.Client
	option *Options

	isConnected bool
}

func NewRouterOS(opt *Options) *RouterOS {
	return &RouterOS{
		option: opt,
	}
}

// NewAndConnctRouterOS creates a new RouterOS client and connects to the router.
func NewAndConnctRouterOS(opt *Options) (*RouterOS, error) {
	router := &RouterOS{
		option: opt,
	}
	return router, router.Connect()
}

// Connect connects to the router.
func (r *RouterOS) Connect() error {
	var err error
	if r.option.Timeout > 0 {
		r.client, err = routeros.DialTimeout(r.option.Address, r.option.Username, r.option.Password, r.option.Timeout)
		if err != nil {
			return err
		}
		r.isConnected = true
		return nil
	}
	r.client, err = routeros.Dial(r.option.Address, r.option.Username, r.option.Password)
	if err != nil {
		return err
	}
	r.isConnected = true
	return nil
}

// Close the connection to the router.
func (r *RouterOS) Close() error {
	if !r.isConnected {
		return fmt.Errorf("router is not connected")
	}
	r.client.Close()
	return nil
}

// Async starts asynchronous mode and returns immediately..
func (r *RouterOS) Async() {
	r.client.Async()
}

// Run runs a sentence on the router.
func (r *RouterOS) CallContext(ctx context.Context, sentence ...string) (*routeros.Reply, error) {
	if !r.isConnected {
		return nil, fmt.Errorf("router is not connected")
	}
	re, err := r.client.RunContext(ctx, sentence...)
	if err != nil {
		if err.Error() == "EOF" {
			var reconnect bool
			if ctx.Value(types.ContextReconnect) != nil {
				reconnect = ctx.Value(types.ContextReconnect).(bool)
			}
			if !reconnect && r.option.AutoReconnect {
				ctx = context.WithValue(ctx, types.ContextReconnect, true)
				if err := r.Connect(); err != nil {
					return nil, err
				}
				return r.CallContext(ctx, sentence...)
			}
			r.isConnected = false
		}
		return nil, err
	}
	return re, nil
}

// GetQuery returns the query of the builder.
func (r *RouterOS) GetPath() string {
	return `/`
}
