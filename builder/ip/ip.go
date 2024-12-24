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

func (b *IPBuilder) GetQuery() string {
	return b.parent.GetQuery() + `ip/`
}

func (b *IPBuilder) GetClient() *routeros.RouterOS {
	return b.parent
}
