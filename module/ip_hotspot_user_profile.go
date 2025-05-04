package module

import (
	"errors"
	"strconv"

	"github.com/aidapedia/go-routeros/types"
	"github.com/aidapedia/go-routeros/util"
)

const (
	HotspotUserProfileModule           Module = "/ip/hotspot/user/profile"
	ErrInvalidActionHotspotUserProfile        = "invalid action for HotspotUserProfile"
)

func init() {
	register(HotspotUserProfileModule, &HotspotUserProfile{})
}

type HotspotUserProfile struct{}

// GetData returns the data of the module.
func (h *HotspotUserProfile) GetData() ModuleDataInterface {
	return &HotspotUserProfileData{}
}

// GetQueryPath returns the path of the module.
func (h *HotspotUserProfile) GetQueryPath() string {
	return string(HotspotUserProfileModule)
}

// CheckAction checks if the action is valid for the module.
func (h *HotspotUserProfile) CheckAction(action types.ActionMap) error {
	switch action {
	case types.ActionMapRemove:
	case types.ActionMapPrint:
		return nil
	default:
		return errors.New(ErrInvalidActionHotspotUserProfile)
	}
	return nil
}

// HotspotUserProfileData is a representation of a logged-in HotSpot user.
type HotspotUserProfileData struct {
	ID                string
	Name              string
	RateLimit         *types.NetworkLimit
	SharedUser        *types.AiInt
	SessionTimeout    *types.AiTimeDuration
	IdleTimeout       *types.AiTimeDuration
	KeepAliveTimeout  *types.AiTimeDuration
	StatusAutorefresh *types.AiTimeDuration
	AddressList       string
	TransparentProxy  bool
	AddMACCookie      bool
	MACCookieTimeout  *types.AiTimeDuration
	Default           bool
}

func (h *HotspotUserProfileData) UnmarshalMap(m map[string]string) error {
	h.ID = m[".id"]
	h.Name = m["name"]
	h.RateLimit = types.ParseNetworkLimit(m["rate-limit"])
	h.SharedUser = util.FindKeyToInt(m, "shared-users")
	h.SessionTimeout = util.FindKeyToDuration(m, "session-timeout")
	h.IdleTimeout = util.FindKeyToDuration(m, "idle-timeout")
	h.KeepAliveTimeout = util.FindKeyToDuration(m, "keepalive-timeout")
	h.StatusAutorefresh = util.FindKeyToDuration(m, "status-autorefresh")
	h.AddressList = m["address-list"]
	h.TransparentProxy = util.FindKeyToBool(m, "transparent-proxy")
	h.AddMACCookie = util.FindKeyToBool(m, "add-mac-cookie")
	h.MACCookieTimeout = util.FindKeyToDuration(m, "mac-cookie-timeout")
	h.Default = util.FindKeyToBool(m, "default")
	return nil
}

func (h *HotspotUserProfileData) MarshalMap(action types.ActionMap) (map[string]string, error) {
	switch action {
	case types.ActionMapAdd:
		result := map[string]string{}
		result["name"] = h.Name
		result["rate-limit"] = h.RateLimit.String()
		result["shared-users"] = h.SharedUser.String()
		util.SetTimeDuration(result, "session-timeout", h.SessionTimeout)
		util.SetTimeDuration(result, "idle-timeout", h.IdleTimeout)
		util.SetTimeDuration(result, "keepalive-timeout", h.KeepAliveTimeout)
		util.SetTimeDuration(result, "status-autorefresh", h.StatusAutorefresh)
		result["address-list"] = h.AddressList
		result["transparent-proxy"] = strconv.FormatBool(h.TransparentProxy)
		result["add-mac-cookie"] = strconv.FormatBool(h.AddMACCookie)
		util.SetTimeDuration(result, "mac-cookie-timeout", h.MACCookieTimeout)
		return result, nil
	case types.ActionMapSet:
		result := map[string]string{}
		result[".id"] = h.ID
		result["name"] = h.Name
		result["rate-limit"] = h.RateLimit.String()
		result["shared-users"] = h.SharedUser.String()
		util.SetTimeDuration(result, "session-timeout", h.SessionTimeout)
		util.SetTimeDuration(result, "idle-timeout", h.IdleTimeout)
		util.SetTimeDuration(result, "keepalive-timeout", h.KeepAliveTimeout)
		util.SetTimeDuration(result, "status-autorefresh", h.StatusAutorefresh)
		result["address-list"] = h.AddressList
		result["transparent-proxy"] = strconv.FormatBool(h.TransparentProxy)
		result["add-mac-cookie"] = strconv.FormatBool(h.AddMACCookie)
		util.SetTimeDuration(result, "mac-cookie-timeout", h.MACCookieTimeout)
		return result, nil
	case types.ActionMapRemove:
		result := map[string]string{}
		result[".id"] = h.ID
		return result, nil
	case types.ActionMapPrint:
		result := map[string]string{}
		result[".id"] = h.ID
		result["name"] = h.Name
		result["rate-limit"] = h.RateLimit.String()
		result["shared-users"] = h.SharedUser.String()
		result["session-timeout"] = h.SessionTimeout.String()
		result["idle-timeout"] = h.IdleTimeout.String()
		result["keepalive-timeout"] = h.KeepAliveTimeout.String()
		result["status-autorefresh"] = h.StatusAutorefresh.String()
		result["address-list"] = h.AddressList
		result["transparent-proxy"] = strconv.FormatBool(h.TransparentProxy)
		result["add-mac-cookie"] = strconv.FormatBool(h.AddMACCookie)
		result["mac-cookie-timeout"] = h.MACCookieTimeout.String()
		result["default"] = strconv.FormatBool(h.Default)
		return result, nil
	default:
		return map[string]string{}, errors.New(ErrInvalidActionHotspotUserProfile)
	}
}
