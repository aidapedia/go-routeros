package ip

import (
	"github.com/aidapedia/airouteros/model"
	"github.com/aidapedia/airouteros/types"
)

// List of logged-in HotSpot users
// IPHotspotHostBuilder is a builder for /ip/hotspot/host.
type IPHotspotHostBuilder struct {
	path   string
	parent *IPHotspotBuilder
}

// NewIPHotspotHostBuilder returns a new IPHotspotHostBuilder.
func NewIPHotspotHostBuilder(hotspot *IPHotspotBuilder) *IPHotspotHostBuilder {
	return &IPHotspotHostBuilder{
		path:   hotspot.GetPath() + `host/`,
		parent: hotspot,
	}
}

// GetPath returns the path of the builder.
func (b *IPHotspotHostBuilder) GetPath() string {
	return b.path
}

// Print queries the hotspot host list based on the given queries.
func (b *IPHotspotHostBuilder) Print(queries model.PrintRequest) ([]model.HotspotHost, error) {
	var (
		result []model.HotspotHost
		path   = b.path + string(types.ActionMapPrint)
	)
	reply, err := b.parent.GetClient().Call(queries.BuildQuery(path)...)
	if err != nil {
		return result, err
	}
	for _, re := range reply.Re {
		result = append(result, model.ParseHotspotHost(re.Map))
	}
	return result, nil
}

// Remove a hotspot host by the given id.
func (b *IPHotspotHostBuilder) Remove(id string) error {
	var (
		path = b.path + string(types.ActionMapRemove)
	)
	_, err := b.parent.GetClient().Call(path, "=.id="+id)
	if err != nil {
		return err
	}
	return nil
}
