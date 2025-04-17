package ip

import (
	"context"

	routeros "github.com/aidapedia/go-routeros"
	"github.com/aidapedia/go-routeros/model"
)

type IPHotspotBuilder struct {
	parent *IPBuilder
}

// NewIPHotspotBuilder returns a new IPHotspotBuilder
func NewIPHotspotBuilder(ip *IPBuilder) *IPHotspotBuilder {
	return &IPHotspotBuilder{
		parent: ip,
	}
}

func (b *IPHotspotBuilder) GetPath() string {
	return b.parent.GetPath() + `hotspot/`
}

func (b *IPHotspotBuilder) Print(ctx context.Context, queries model.PrintRequest) ([]model.Hotspot, error) {
	var results []model.Hotspot
	reply, err := b.parent.GetClient().CallContext(ctx, queries.BuildQuery(b.GetPath()+`print`)...)
	if err != nil {
		return results, err
	}
	for _, re := range reply.Re {
		result := model.Hotspot(re.Map)
		results = append(results, result)
	}
	return results, nil
}

func (b *IPHotspotBuilder) GetClient() *routeros.RouterOS {
	return b.parent.GetClient()
}
