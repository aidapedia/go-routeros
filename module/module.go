package module

import (
	"fmt"

	"github.com/aidapedia/go-routeros/types"
)

// ModuleDataInterface is an interface for all modules.
type ModuleDataInterface interface {
	UnmarshalMap(map[string]string) error
	MarshalMap(action types.ActionMap) (map[string]string, error)
}

type ModuleInterface interface {
	GetQueryPath() string
	GetData() ModuleDataInterface
	CheckAction(action types.ActionMap) error
}

type Module string

var modules = map[Module]ModuleInterface{}

func Get(name Module) (ModuleInterface, error) {
	module, exist := modules[name]
	if !exist {
		return nil, fmt.Errorf("unknown module: %s", name)
	}

	return module, nil
}

func register(name Module, module ModuleInterface) {
	if module == nil {
		panic("Module is nil")
	}
	if _, exist := modules[name]; exist {
		panic("Module has been registered with same name: " + name)
	}
	modules[name] = module
}
