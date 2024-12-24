package ip

import (
	"github.com/aidapedia/airouteros/model"
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

func (b *IPHotspotActiveBuilder) Print(queries model.PrintRequest) ([]model.HotspotActive, error) {
	var results []model.HotspotActive
	reply, err := b.parent.GetClient().Call(queries.BuildQuery(b.GetQuery() + `print`)...)
	if err != nil {
		return results, err
	}
	for _, re := range reply.Re {
		result := model.HotspotActive(re.Map)
		results = append(results, result)
	}
	return results, nil
}
