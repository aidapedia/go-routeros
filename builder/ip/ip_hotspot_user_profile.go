package ip

import (
	"github.com/aidapedia/airouteros/model"
)

// List of logged-in HotSpot users
// IPHotspotUserProfileBuilder is a builder for /ip/hotspot/user/profile.
type IPHotspotUserProfileBuilder struct {
	parent *IPHotspotBuilder
}

// NewIPHotspotUserProfileBuilder returns a new IPHotspotUserProfileBuilder.
func NewIPHotspotUserProfileBuilder(hotspot *IPHotspotBuilder) *IPHotspotUserProfileBuilder {
	return &IPHotspotUserProfileBuilder{
		parent: hotspot,
	}
}

// GetQuery returns the query of the builder.
func (b *IPHotspotUserProfileBuilder) GetQuery() string {
	return b.parent.GetQuery() + `user/profile/`
}

func (b *IPHotspotUserProfileBuilder) Print(queries model.PrintRequest) ([]model.HotspotUserProfile, error) {
	var results []model.HotspotUserProfile
	reply, err := b.parent.GetClient().Call(queries.BuildQuery(b.GetQuery() + `print`)...)
	if err != nil {
		return results, err
	}
	for _, re := range reply.Re {
		result := model.ParseHotspotUserProfile(re.Map)
		results = append(results, result)
	}
	return results, nil
}
