package module

import (
	"errors"
	"strconv"

	"github.com/aidapedia/go-routeros/types"
	"github.com/aidapedia/go-routeros/util"
)

const (
	HotspotModule           Module = "/ip/hotspot"
	ErrInvalidActionHotspot        = "invalid action for Hotspot"
)

func init() {
	register(HotspotModule, &Hotspot{})
}

type Hotspot struct{}

// GetData returns the data of the module.
func (h *Hotspot) GetData() ModuleDataInterface {
	return &HotspotData{}
}

// GetQueryPath returns the path of the module.
func (h *Hotspot) GetQueryPath() string {
	return string(HotspotModule)
}

// CheckAction checks if the action is valid for the module.
func (h *Hotspot) CheckAction(action types.ActionMap) error {
	switch action {
	case types.ActionMapRemove:
	case types.ActionMapSet:
	case types.ActionMapAdd:
	case types.ActionMapPrint:
		return nil
	default:
		return errors.New(ErrInvalidActionHotspot)
	}
	return nil
}

// HotspotData is a representation of a logged-in HotSpot user.
type HotspotData struct {
	ID               string
	Interface        string
	Invalid          bool
	Disabled         bool
	Name             string
	LoginTimeout     *types.AiTimeDuration
	KeepAliveTimeout *types.AiTimeDuration
	IPofDNSName      string
	AddressesPerMAC  int
	HTTPS            bool
	Profile          string
	IdleTimeout      *types.AiTimeDuration
	ProxyStatus      string
}

func (h *HotspotData) UnmarshalMap(m map[string]string) error {
	h.ID = m[".id"]
	h.Interface = m["interface"]
	h.Invalid = util.FindKeyToBool(m, "invalid")
	h.Disabled = util.FindKeyToBool(m, "disabled")
	h.Name = m["name"]
	h.LoginTimeout = util.FindKeyToDuration(m, "login-timeout")
	h.KeepAliveTimeout = util.FindKeyToDuration(m, "keepalive-timeout")
	h.IPofDNSName = m["ip-of-dns-name"]
	h.AddressesPerMAC, _ = strconv.Atoi(m["addresses-per-mac"])
	h.HTTPS = util.FindKeyToBool(m, "https")
	h.Profile = m["profile"]
	h.IdleTimeout = util.FindKeyToDuration(m, "idle-timeout")
	h.ProxyStatus = m["proxy-status"]
	return nil
}

func (h *HotspotData) MarshalMap(action types.ActionMap) (map[string]string, error) {
	switch action {
	case types.ActionMapPrint:
		result := map[string]string{}
		result[".id"] = h.ID
		result["interface"] = h.Interface
		result["invalid"] = strconv.FormatBool(h.Invalid)
		result["disabled"] = strconv.FormatBool(h.Disabled)
		result["name"] = h.Name
		result["login-timeout"] = h.LoginTimeout.String()
		result["keepalive-timeout"] = h.KeepAliveTimeout.String()
		result["ip-of-dns-name"] = h.IPofDNSName
		result["addresses-per-mac"] = strconv.Itoa(h.AddressesPerMAC)
		result["https"] = strconv.FormatBool(h.HTTPS)
		result["profile"] = h.Profile
		result["idle-timeout"] = h.IdleTimeout.String()
		result["proxy-status"] = h.ProxyStatus
		return result, nil
	default:
		return map[string]string{}, errors.New(ErrInvalidActionHotspot)
	}
}
