package builder

import routeros "github.com/aoida/router-os"

type IPBuilder struct {
	parent *routeros.RouterOSBuilder
}

func NewIPBuilder(builder *routeros.RouterOSBuilder) *IPBuilder {
	return &IPBuilder{
		parent: builder,
	}
}

func (b *IPBuilder) GetQuery() string {
	return b.parent.GetQuery() + `ip/`
}

func (b *IPBuilder) GetClient() *routeros.RouterOSBuilder {
	return b.parent
}
