package builder

import (
	"github.com/aoida/router-os/builder/entity"
)

// List of logged-in HotSpot users
// IPHotspotActiveBuilder is a builder for /ip/hotspot/active.
type IPHotspotActiveBuilder struct {
	parent *IPHotspotBuilder
}

// NewIPHotspotActiveBuilder returns a new IPHotspotActiveBuilder.
func NewIPHotspotActiveBuilder(hotspot *IPHotspotBuilder) *IPHotspotActiveBuilder {
	return &IPHotspotActiveBuilder{
		parent: hotspot,
	}
}

// GetQuery returns the query of the builder.
func (b *IPHotspotActiveBuilder) GetQuery() string {
	return b.parent.GetQuery() + `active/`
}

func (b *IPHotspotActiveBuilder) Print(queries entity.PrintRequest) ([]entity.HotspotActive, error) {
	var results []entity.HotspotActive
	reply, err := b.parent.GetClient().Run(queries.BuildQuery(b.GetQuery() + `print`)...)
	if err != nil {
		return results, err
	}
	for _, re := range reply.Re {
		result := entity.HotspotActive(re.Map)
		results = append(results, result)
	}
	return results, nil
}
