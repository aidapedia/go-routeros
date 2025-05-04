package module

import (
	"errors"
	"strconv"

	"github.com/aidapedia/go-routeros/types"
	"github.com/aidapedia/go-routeros/util"
)

const (
	HotspotHostModule           Module = "/ip/hotspot/host"
	ErrInvalidActionHotspotHost        = "invalid action for HotspotHost"
)

func init() {
	register(HotspotHostModule, &HotspotHost{})
}

type HotspotHost struct{}

// GetData returns the data of the module.
func (h *HotspotHost) GetData() ModuleDataInterface {
	return &HotspotHostData{}
}

// GetQueryPath returns the path of the module.
func (h *HotspotHost) GetQueryPath() string {
	return string(HotspotHostModule)
}

// CheckAction checks if the action is valid for the module.
func (h *HotspotHost) CheckAction(action types.ActionMap) error {
	switch action {
	case types.ActionMapRemove:
	case types.ActionMapPrint:
		return nil
	default:
		return errors.New(ErrInvalidActionHotspotHost)
	}
	return nil
}

// HotspotHostData is a representation of a logged-in HotSpot user.
type HotspotHostData struct {
	ID           string
	MacAddress   string
	Address      string
	Server       string
	Uptime       *types.AiTimeDuration
	BytesIn      *types.AiInt64
	BytesOut     *types.AiInt64
	PacketsIn    *types.AiInt64
	PacketsOut   *types.AiInt64
	IdleTimeout  *types.AiTimeDuration
	IdleTime     *types.AiTimeDuration
	Bypassed     bool
	HostDeadTime *types.AiTimeDuration
	FoundBy      string
	DHCP         bool
	Authorized   bool
	Comment      string
}

func (h *HotspotHostData) UnmarshalMap(m map[string]string) error {
	h.ID = m[".id"]
	h.MacAddress = m["mac-address"]
	h.Address = m["address"]
	h.Server = m["server"]
	h.Uptime = util.FindKeyToDuration(m, "uptime")
	h.BytesIn = util.FindKeyToInt64(m, "bytes-in")
	h.BytesOut = util.FindKeyToInt64(m, "bytes-out")
	h.PacketsIn = util.FindKeyToInt64(m, "packets-in")
	h.PacketsOut = util.FindKeyToInt64(m, "packets-out")
	h.IdleTimeout = util.FindKeyToDuration(m, "idle-timeout")
	h.IdleTime = util.FindKeyToDuration(m, "idle-time")
	h.Bypassed = util.FindKeyToBool(m, "bypassed")
	h.HostDeadTime = util.FindKeyToDuration(m, "host-dead-time")
	h.FoundBy = m["found-by"]
	h.DHCP = util.FindKeyToBool(m, "dhcp")
	h.Authorized = util.FindKeyToBool(m, "authorized")
	h.Comment = m["comment"]
	return nil
}

func (h *HotspotHostData) MarshalMap(action types.ActionMap) (map[string]string, error) {
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
		result["server"] = h.Server
		result["uptime"] = h.Uptime.String()
		result["bytes-in"] = h.BytesIn.String()
		result["bytes-out"] = h.BytesOut.String()
		result["packets-in"] = h.PacketsIn.String()
		result["packets-out"] = h.PacketsOut.String()
		result["idle-timeout"] = h.IdleTimeout.String()
		result["idle-time"] = h.IdleTime.String()
		result["bypassed"] = strconv.FormatBool(h.Bypassed)
		result["host-dead-time"] = h.HostDeadTime.String()
		result["found-by"] = h.FoundBy
		result["dhcp"] = strconv.FormatBool(h.DHCP)
		result["authorized"] = strconv.FormatBool(h.Authorized)
		result["comment"] = h.Comment
		return result, nil
	default:
		return map[string]string{}, errors.New(ErrInvalidActionHotspotHost)
	}
}
