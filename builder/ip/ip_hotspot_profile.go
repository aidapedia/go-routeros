package ip

import (
	"github.com/aidapedia/airouteros/model"
)

// List of logged-in HotSpot users
// IPHotspotProfileBuilder is a builder for /ip/hotspot/profile.
type IPHotspotProfileBuilder struct {
	parent *IPHotspotBuilder
}

// NewIPHotspotProfileBuilder returns a new IPHotspotProfileBuilder.
func NewIPHotspotProfileBuilder(hotspot *IPHotspotBuilder) *IPHotspotProfileBuilder {
	return &IPHotspotProfileBuilder{
		parent: hotspot,
	}
}

// GetQuery returns the query of the builder.
func (b *IPHotspotProfileBuilder) GetQuery() string {
	return b.parent.GetQuery() + `profile/`
}

func (b *IPHotspotProfileBuilder) Print(queries model.PrintRequest) ([]model.HotspotProfile, error) {
	var results []model.HotspotProfile
	reply, err := b.parent.GetClient().Call(queries.BuildQuery(b.GetQuery() + `print`)...)
	if err != nil {
		return results, err
	}
	for _, re := range reply.Re {
		result := model.HotspotProfile(re.Map)
		results = append(results, result)
	}
	return results, nil
}
