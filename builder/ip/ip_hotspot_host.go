package ip

import (
	"github.com/aidapedia/airouteros/model"
)

// List of logged-in HotSpot users
// IPHotspotHostBuilder is a builder for /ip/hotspot/host.
type IPHotspotHostBuilder struct {
	parent *IPHotspotBuilder
}

// NewIPHotspotHostBuilder returns a new IPHotspotHostBuilder.
func NewIPHotspotHostBuilder(hotspot *IPHotspotBuilder) *IPHotspotHostBuilder {
	return &IPHotspotHostBuilder{
		parent: hotspot,
	}
}

// GetQuery returns the query of the builder.
func (b *IPHotspotHostBuilder) GetQuery() string {
	return b.parent.GetQuery() + `host/`
}

func (b *IPHotspotHostBuilder) Print(queries model.PrintRequest) ([]model.HotspotHost, error) {
	var results []model.HotspotHost
	reply, err := b.parent.GetClient().Call(queries.BuildQuery(b.GetQuery() + `print`)...)
	if err != nil {
		return results, err
	}
	for _, re := range reply.Re {
		result := model.HotspotHost(re.Map)
		results = append(results, result)
	}
	return results, nil
}
