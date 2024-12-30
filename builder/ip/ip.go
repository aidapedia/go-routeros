package ip

import (
	routeros "github.com/aidapedia/airouteros"
)

type IPBuilder struct {
	parent *routeros.RouterOS
}

func NewIPBuilder(builder *routeros.RouterOS) *IPBuilder {
	return &IPBuilder{
		parent: builder,
	}
}

func (b *IPBuilder) GetPath() string {
	return b.parent.GetPath() + `ip/`
}

func (b *IPBuilder) GetClient() *routeros.RouterOS {
	return b.parent
}
