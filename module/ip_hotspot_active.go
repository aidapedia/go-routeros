package module

import (
	"errors"

	"github.com/aidapedia/go-routeros/types"
	"github.com/aidapedia/go-routeros/util"
)

const (
	HotspotActiveModule           Module = "/ip/hotspot/active"
	ErrInvalidActionHotspotActive        = "invalid action for HotspotActive"
)

func init() {
	register(HotspotActiveModule, &HotspotActive{})
}

type HotspotActive struct{}

// GetData returns the data of the module.
func (h *HotspotActive) GetData() ModuleDataInterface {
	return &HotspotActiveData{}
}

// GetQueryPath returns the path of the module.
func (h *HotspotActive) GetQueryPath() string {
	return string(HotspotActiveModule)
}

// CheckAction checks if the action is valid for the module.
func (h *HotspotActive) CheckAction(action types.ActionMap) error {
	switch action {
	case types.ActionMapRemove:
	case types.ActionMapPrint:
		return nil
	default:
		return errors.New(ErrInvalidActionHotspotActive)
	}
	return nil
}

// HotspotActiveData is a representation of a logged-in HotSpot user.
type HotspotActiveData struct {
	ID         string
	MacAddress string
	Address    string
	User       string
	Server     string
	Uptime     *types.AiTimeDuration
	BytesIn    *types.AiInt64
	BytesOut   *types.AiInt64
	PacketsIn  *types.AiInt64
	PacketsOut *types.AiInt64
	Comment    string
	IdleTime   *types.AiTimeDuration
}

func (h *HotspotActiveData) UnmarshalMap(m map[string]string) error {
	h.ID = m[".id"]
	h.MacAddress = m["mac-address"]
	h.Address = m["address"]
	h.User = m["user"]
	h.Server = m["server"]
	h.Uptime = util.FindKeyToDuration(m, "uptime")
	h.BytesIn = util.FindKeyToInt64(m, "bytes-in")
	h.BytesOut = util.FindKeyToInt64(m, "bytes-out")
	h.PacketsIn = util.FindKeyToInt64(m, "packets-in")
	h.PacketsOut = util.FindKeyToInt64(m, "packets-out")
	h.Comment = m["comment"]
	h.IdleTime = util.FindKeyToDuration(m, "idle-time")
	return nil
}

func (h *HotspotActiveData) MarshalMap(action types.ActionMap) (map[string]string, error) {
	switch action {
	case types.ActionMapRemove:
		result := map[string]string{}
		result[".id"] = h.ID
		return result, nil
	case types.ActionMapPrint:
		result := map[string]string{}
		result[".id"] = h.ID
		result["mac-address"] = h.MacAddress
		result["address"] = h.Address
		result["user"] = h.User
		result["server"] = h.Server
		result["uptime"] = h.Uptime.String()
		result["bytes-in"] = h.BytesIn.String()
		result["bytes-out"] = h.BytesOut.String()
		result["packets-in"] = h.PacketsIn.String()
		result["packets-out"] = h.PacketsOut.String()
		result["comment"] = h.Comment
		result["idle-time"] = h.IdleTime.String()
		return result, nil
	default:
		return map[string]string{}, errors.New(ErrInvalidActionHotspotActive)
	}
}
