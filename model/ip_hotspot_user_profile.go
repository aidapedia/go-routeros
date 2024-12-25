package model

import (
	"strconv"

	"github.com/aidapedia/airouteros/types"
	"github.com/aidapedia/airouteros/util"
)

// HotspotUserProfile represents a request to hotspot user profile.
type HotspotUserProfile struct {
	ID                string
	Name              string
	RateLimit         *types.NetworkLimit
	SharedUser        types.AiInt
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
		SharedUser:        types.AiInt(util.FindKeyToInt(m, "shared-users")),
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

func (h HotspotUserProfile) ToMap() map[string]string {
	return map[string]string{
		".id":                h.ID,
		"name":               h.Name,
		"rate-limit":         h.RateLimit.String(),
		"shared-users":       h.SharedUser.String(),
		"session-timeout":    h.SessionTimeout.String(),
		"idle-timeout":       h.IdleTimeout.String(),
		"keepalive-timeout":  h.KeepAliveTimeout.String(),
		"status-autorefresh": h.StatusAutorefresh.String(),
		"address-list":       h.AddressList,
		"transparent-proxy":  strconv.FormatBool(h.TransparentProxy),
		"add-mac-cookie":     strconv.FormatBool(h.AddMACCookie),
		"mac-cookie-timeout": h.MACCookieTimeout.String(),
		"default":            strconv.FormatBool(h.Default),
	}
}
