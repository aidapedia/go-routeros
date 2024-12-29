package model

import (
	"github.com/aidapedia/airouteros/types"
	"github.com/aidapedia/airouteros/util"
)

// HotspotActive is a representation of a logged-in HotSpot user.
type HotspotActive struct {
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

func ParseHotspotActive(m map[string]string) HotspotActive {
	return HotspotActive{
		ID:         m[".id"],
		MacAddress: m["mac-address"],
		Address:    m["address"],
		User:       m["user"],
		Server:     m["server"],
		Uptime:     util.FindKeyToDuration(m, "uptime"),
		BytesIn:    util.FindKeyToInt64(m, "bytes-in"),
		BytesOut:   util.FindKeyToInt64(m, "bytes-out"),
		PacketsIn:  util.FindKeyToInt64(m, "packets-in"),
		PacketsOut: util.FindKeyToInt64(m, "packets-out"),
		Comment:    m["comment"],
		IdleTime:   util.FindKeyToDuration(m, "idle-time"),
	}
}

func (h HotspotActive) ToMap(action types.ActionMap) map[string]string {
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
		result["user"] = h.User
		result["server"] = h.Server
		result["uptime"] = h.Uptime.String()
		result["bytes-in"] = h.BytesIn.String()
		result["bytes-out"] = h.BytesOut.String()
		result["packets-in"] = h.PacketsIn.String()
		result["packets-out"] = h.PacketsOut.String()
		result["comment"] = h.Comment
		result["idle-time"] = h.IdleTime.String()
		return result
	}
}
