package ip

import (
	"context"

	"github.com/aidapedia/go-routeros/model"
	"github.com/aidapedia/go-routeros/types"
	"github.com/aidapedia/go-routeros/util"
)

// List of logged-in HotSpot users
// IPHotspotUserProfileBuilder is a builder for /ip/hotspot/user/profile.
type IPHotspotUserProfileBuilder struct {
	path   string
	parent *IPHotspotBuilder
}

// NewIPHotspotUserProfileBuilder returns a new IPHotspotUserProfileBuilder.
func NewIPHotspotUserProfileBuilder(hotspot *IPHotspotBuilder) *IPHotspotUserProfileBuilder {
	return &IPHotspotUserProfileBuilder{
		path:   hotspot.GetPath() + `user/profile/`,
		parent: hotspot,
	}
}

// GetPath returns the path of the builder.
func (b *IPHotspotUserProfileBuilder) GetPath() string {
	return b.path
}

// Print queries the hotspot user profile list based on the given queries.
func (b *IPHotspotUserProfileBuilder) Print(ctx context.Context, queries model.PrintRequest) ([]model.HotspotUserProfile, error) {
	var (
		results []model.HotspotUserProfile
		path    = b.path + string(types.ActionMapPrint)
	)
	reply, err := b.parent.GetClient().CallContext(ctx, queries.BuildQuery(path)...)
	if err != nil {
		return results, err
	}
	for _, re := range reply.Re {
		result := model.ParseHotspotUserProfile(re.Map)
		results = append(results, result)
	}
	return results, nil
}

// Add a new hotspot user profile.
func (b *IPHotspotUserProfileBuilder) Add(ctx context.Context, request model.HotspotUserProfile) error {
	var (
		path = b.path + string(types.ActionMapAdd)
	)
	_, err := b.parent.GetClient().CallContext(ctx, util.ToQuery(path, request.ToMap(types.ActionMapAdd))...)
	if err != nil {
		return err
	}
	return nil
}

// Set the hotspot user profile.
func (b *IPHotspotUserProfileBuilder) Set(ctx context.Context, request model.HotspotUserProfile) error {
	var (
		path = b.path + string(types.ActionMapSet)
	)
	_, err := b.parent.GetClient().CallContext(ctx, util.ToQuery(path, request.ToMap(types.ActionMapSet))...)
	if err != nil {
		return err
	}
	return nil
}

// Remove the hotspot user profile based on the given id.
func (b *IPHotspotUserProfileBuilder) Remove(ctx context.Context, id string) error {
	var (
		path = b.path + string(types.ActionMapRemove)
	)
	_, err := b.parent.GetClient().CallContext(ctx, path, "=.id="+id)
	if err != nil {
		return err
	}
	return nil
}
