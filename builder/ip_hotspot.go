package builder

import (
	routeros "github.com/aoida/router-os"
	"github.com/aoida/router-os/builder/entity"
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

func (b *IPHotspotBuilder) Print(queries entity.PrintRequest) ([]entity.Hotspot, error) {
	var results []entity.Hotspot
	reply, err := b.parent.GetClient().Run(queries.BuildQuery(b.GetQuery() + `print`)...)
	if err != nil {
		return results, err
	}
	for _, re := range reply.Re {
		result := entity.Hotspot(re.Map)
		results = append(results, result)
	}
	return results, nil
}

func (b *IPHotspotBuilder) GetClient() *routeros.RouterOSBuilder {
	return b.parent.GetClient()
}
