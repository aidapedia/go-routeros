package model

import (
	"strconv"

	"github.com/aidapedia/go-routeros/types"
	"github.com/aidapedia/go-routeros/util"
)

type HotspotHost struct {
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

func ParseHotspotHost(m map[string]string) HotspotHost {
	return HotspotHost{
		ID:           m[".id"],
		MacAddress:   m["mac-address"],
		Address:      m["address"],
		Server:       m["server"],
		Uptime:       util.FindKeyToDuration(m, "uptime"),
		BytesIn:      util.FindKeyToInt64(m, "bytes-in"),
		BytesOut:     util.FindKeyToInt64(m, "bytes-out"),
		PacketsIn:    util.FindKeyToInt64(m, "packets-in"),
		PacketsOut:   util.FindKeyToInt64(m, "packets-out"),
		IdleTimeout:  util.FindKeyToDuration(m, "idle-timeout"),
		IdleTime:     util.FindKeyToDuration(m, "idle-time"),
		Bypassed:     util.FindKeyToBool(m, "bypassed"),
		HostDeadTime: util.FindKeyToDuration(m, "host-dead-time"),
		FoundBy:      m["found-by"],
		DHCP:         util.FindKeyToBool(m, "dhcp"),
		Authorized:   util.FindKeyToBool(m, "authorized"),
		Comment:      m["comment"],
	}
}

func (h HotspotHost) ToMap(action types.ActionMap) map[string]string {
	switch action {
	case types.ActionMapAdd:
		return map[string]string{}
	case types.ActionMapSet:
		return map[string]string{}
	case types.ActionMapRemove:
		result := map[string]string{}
		result[".id"] = h.ID
		return result
	case types.ActionMapPrint:
		fallthrough
	default:
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
		return result
	}
}
