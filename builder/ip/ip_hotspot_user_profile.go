package ip

import (
	"github.com/aidapedia/airouteros/model"
	"github.com/aidapedia/airouteros/types"
	"github.com/aidapedia/airouteros/util"
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
	var (
		results []model.HotspotUserProfile
		path    = b.GetQuery() + string(types.ActionMapPrint)
	)
	reply, err := b.parent.GetClient().Call(queries.BuildQuery(path)...)
	if err != nil {
		return results, err
	}
	for _, re := range reply.Re {
		result := model.ParseHotspotUserProfile(re.Map)
		results = append(results, result)
	}
	return results, nil
}

func (b *IPHotspotUserProfileBuilder) Add(request model.HotspotUserProfile) error {
	var (
		path = b.GetQuery() + string(types.ActionMapAdd)
	)
	_, err := b.parent.GetClient().Call(util.ToQuery(path, request.ToMap(types.ActionMapAdd))...)
	if err != nil {
		return err
	}
	return nil
}

func (b *IPHotspotUserProfileBuilder) Set(request model.HotspotUserProfile) error {
	var (
		path = b.GetQuery() + string(types.ActionMapSet)
	)
	_, err := b.parent.GetClient().Call(util.ToQuery(path, request.ToMap(types.ActionMapSet))...)
	if err != nil {
		return err
	}
	return nil
}

func (b *IPHotspotUserProfileBuilder) Remove(id string) error {
	var (
		path = b.GetQuery() + string(types.ActionMapRemove)
	)
	_, err := b.parent.GetClient().Call(path, "=.id="+id)
	if err != nil {
		return err
	}
	return nil
}
