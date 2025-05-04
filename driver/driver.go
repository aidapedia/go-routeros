package driver

import (
	"context"
	"strings"

	"github.com/aidapedia/go-routeros"
	"github.com/aidapedia/go-routeros/model"
	"github.com/aidapedia/go-routeros/module"
	"github.com/aidapedia/go-routeros/types"
	"github.com/aidapedia/go-routeros/util"
)

// Driver is a driver for the RouterOS API.
type Driver struct {
	client *routeros.RouterOS
	module module.ModuleInterface
}

// New creates a new driver.
func New(c *routeros.RouterOS, m module.Module) (*Driver, error) {
	mod, err := module.Get(m)
	if err != nil {
		return nil, err
	}
	return &Driver{
		client: c,
		module: mod,
	}, nil
}

// Print queries the data list based on the given queries.
func (d *Driver) Print(ctx context.Context, queries model.PrintRequest) ([]module.ModuleDataInterface, error) {
	var (
		result []module.ModuleDataInterface
		path   = strings.Join([]string{d.module.GetQueryPath(), string(types.ActionMapPrint)}, "/")
	)
	if err := d.module.CheckAction(types.ActionMapPrint); err != nil {
		return nil, err
	}
	reply, err := d.client.CallContext(ctx, queries.BuildQuery(path)...)
	if err != nil {
		return nil, err
	}
	for _, re := range reply.Re {
		data := d.module.GetData()
		err := data.UnmarshalMap(re.Map)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}
	return result, nil
}

// Add is used to add a new data.
func (d *Driver) Add(ctx context.Context, request module.ModuleDataInterface) error {
	var (
		path = strings.Join([]string{d.module.GetQueryPath(), string(types.ActionMapAdd)}, "/")
	)
	if err := d.module.CheckAction(types.ActionMapAdd); err != nil {
		return err
	}
	req, err := request.MarshalMap(types.ActionMapAdd)
	if err != nil {
		return err
	}
	_, err = d.client.CallContext(ctx, util.ToQuery(path, req)...)
	if err != nil {
		return err
	}
	return nil
}

// Set is used to update an existing data.
func (d *Driver) Set(ctx context.Context, request module.ModuleDataInterface) error {
	var (
		path = strings.Join([]string{d.module.GetQueryPath(), string(types.ActionMapSet)}, "/")
	)
	if err := d.module.CheckAction(types.ActionMapSet); err != nil {
		return err
	}
	req, err := request.MarshalMap(types.ActionMapSet)
	if err != nil {
		return err
	}
	_, err = d.client.CallContext(ctx, util.ToQuery(path, req)...)
	if err != nil {
		return err
	}
	return nil
}

// Remove a data that matches the given id.
func (d *Driver) Remove(ctx context.Context, id string) error {
	var (
		path = strings.Join([]string{d.module.GetQueryPath(), string(types.ActionMapRemove)}, "/")
	)
	if err := d.module.CheckAction(types.ActionMapPrint); err != nil {
		return err
	}
	_, err := d.client.CallContext(ctx, path, "=.id="+id)
	if err != nil {
		return err
	}
	return nil
}
