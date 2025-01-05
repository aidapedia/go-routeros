package ip

import (
	"context"

	"github.com/aidapedia/airouteros/model"
	"github.com/aidapedia/airouteros/types"
)

// List of logged-in HotSpot users
// IPHotspotActiveBuilder is a builder for /ip/hotspot/active.
type IPHotspotActiveBuilder struct {
	path   string
	parent *IPHotspotBuilder
}

// NewIPHotspotActiveBuilder returns a new IPHotspotActiveBuilder.
func NewIPHotspotActiveBuilder(hotspot *IPHotspotBuilder) *IPHotspotActiveBuilder {
	return &IPHotspotActiveBuilder{
		path:   hotspot.GetPath() + `active/`,
		parent: hotspot,
	}
}

// GetPath returns the path of the builder.
func (b *IPHotspotActiveBuilder) GetPath() string {
	return b.path
}

// Print queries the hotspot active list based on the given queries.
func (b *IPHotspotActiveBuilder) Print(ctx context.Context, queries model.PrintRequest) ([]model.HotspotActive, error) {
	var (
		results []model.HotspotActive
		path    = b.path + string(types.ActionMapPrint)
	)
	reply, err := b.parent.GetClient().CallContext(ctx, queries.BuildQuery(path)...)
	if err != nil {
		return results, err
	}
	for _, re := range reply.Re {
		result := model.ParseHotspotActive(re.Map)
		results = append(results, result)
	}
	return results, nil
}

// Remove a hotspot active by the given id.
func (b *IPHotspotActiveBuilder) Remove(ctx context.Context, id string) error {
	var (
		path = b.path + string(types.ActionMapRemove)
	)
	_, err := b.parent.GetClient().CallContext(ctx, path, "=.id="+id)
	if err != nil {
		return err
	}
	return nil
}
