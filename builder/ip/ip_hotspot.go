package ip

import (
	routeros "github.com/aidapedia/airouteros"
	"github.com/aidapedia/airouteros/model"
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

func (b *IPHotspotBuilder) GetQuery() string {
	return b.parent.GetQuery() + `hotspot/`
}

func (b *IPHotspotBuilder) Print(queries model.PrintRequest) ([]model.Hotspot, error) {
	var results []model.Hotspot
	reply, err := b.parent.GetClient().Call(queries.BuildQuery(b.GetQuery() + `print`)...)
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
