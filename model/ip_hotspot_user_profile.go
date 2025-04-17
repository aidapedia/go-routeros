package model

import (
	"strconv"

	"github.com/aidapedia/go-routeros/types"
	"github.com/aidapedia/go-routeros/util"
)

// HotspotUserProfile represents a request to hotspot user profile.
type HotspotUserProfile struct {
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

func ParseHotspotUserProfile(m map[string]string) HotspotUserProfile {
	return HotspotUserProfile{
		ID:                m[".id"],
		Name:              m["name"],
		RateLimit:         types.ParseNetworkLimit(m["rate-limit"]),
		SharedUser:        util.FindKeyToInt(m, "shared-users"),
		SessionTimeout:    util.FindKeyToDuration(m, "session-timeout"),
		IdleTimeout:       util.FindKeyToDuration(m, "idle-timeout"),
		KeepAliveTimeout:  util.FindKeyToDuration(m, "keepalive-timeout"),
		StatusAutorefresh: util.FindKeyToDuration(m, "status-autorefresh"),
		AddressList:       m["address-list"],
		TransparentProxy:  util.FindKeyToBool(m, "transparent-proxy"),
		AddMACCookie:      util.FindKeyToBool(m, "add-mac-cookie"),
		MACCookieTimeout:  util.FindKeyToDuration(m, "mac-cookie-timeout"),
		Default:           util.FindKeyToBool(m, "default"),
	}
}

func (h HotspotUserProfile) ToMap(action types.ActionMap) map[string]string {
	switch action {
	case types.ActionMapAdd:
		return h.mutableToMap()
	case types.ActionMapSet:
		result := h.mutableToMap()
		result[".id"] = h.ID
		return result
	case types.ActionMapRemove:
		result := map[string]string{}
		result[".id"] = h.ID
		return result
	case types.ActionMapPrint:
		fallthrough
	default:
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
		return result
	}
}

func (h HotspotUserProfile) mutableToMap() map[string]string {
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
	return result
}
