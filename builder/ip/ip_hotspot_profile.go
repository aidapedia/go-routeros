package ip

import (
	"context"

	"github.com/aidapedia/go-routeros/model"
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
func (b *IPHotspotProfileBuilder) GetPath() string {
	return b.parent.GetPath() + `profile/`
}

func (b *IPHotspotProfileBuilder) Print(ctx context.Context, queries model.PrintRequest) ([]model.HotspotProfile, error) {
	var results []model.HotspotProfile
	reply, err := b.parent.GetClient().CallContext(ctx, queries.BuildQuery(b.GetPath()+`print`)...)
	if err != nil {
		return results, err
	}
	for _, re := range reply.Re {
		result := model.HotspotProfile(re.Map)
		results = append(results, result)
	}
	return results, nil
}
