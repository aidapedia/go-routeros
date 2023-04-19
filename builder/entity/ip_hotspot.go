package entity

type Hotspot map[string]string

// GetID returns the id of the hotspot active.
func (h Hotspot) GetID() string {
	result, ok := h[".id"]
	if !ok {
		return ""
	}
	return result
}

// // GetMacAddress returns the mac address of the hotspot active.
// func (h HotspotActive) GetMacAddress() string {
// 	result, ok := h["mac-address"]
// 	if !ok {
// 		return ""
// 	}
// 	return result
// }

// // GetAddress returns the address of the hotspot active.
// func (h HotspotActive) GetAddress() string {
// 	result, ok := h["address"]
// 	if !ok {
// 		return ""
// 	}
// 	return result
// }

// // GetUser returns the user of the hotspot active.
// func (h HotspotActive) GetUser() string {
// 	result, ok := h["user"]
// 	if !ok {
// 		return ""
// 	}
// 	return result
// }

// // GetServer returns the server of the hotspot active.
// func (h HotspotActive) GetServer() string {
// 	result, ok := h["server"]
// 	if !ok {
// 		return ""
// 	}
// 	return result
// }

// // GetUptime returns the uptime of the hotspot active.
// func (h HotspotActive) GetUptime() time.Duration {
// 	result, ok := h["uptime"]
// 	if !ok {
// 		return time.Duration(0)
// 	}
// 	resultDur, err := time.ParseDuration(result)
// 	if err != nil {
// 		return time.Duration(0)
// 	}
// 	return resultDur
// }

// // GetBytesIn returns the bytes in of the hotspot active.
// func (h HotspotActive) GetBytesIn() int64 {
// 	result, ok := h["bytes-in"]
// 	if !ok {
// 		return 0
// 	}
// 	resInt64, err := strconv.ParseInt(result, 10, 64)
// 	if err != nil {
// 		return 0
// 	}
// 	return resInt64
// }

// // GetBytesOut returns the bytes out of the hotspot active.
// func (h HotspotActive) GetBytesOut() int64 {
// 	result, ok := h["bytes-out"]
// 	if !ok {
// 		return 0
// 	}
// 	resInt64, err := strconv.ParseInt(result, 10, 64)
// 	if err != nil {
// 		return 0
// 	}
// 	return resInt64
// }

// // GetPacketsIn returns the packets in of the hotspot active.
// func (h HotspotActive) GetPacketsIn() int64 {
// 	result, ok := h["packets-in"]
// 	if !ok {
// 		return 0
// 	}
// 	resInt64, err := strconv.ParseInt(result, 10, 64)
// 	if err != nil {
// 		return 0
// 	}
// 	return resInt64
// }

// // GetPacketsOut returns the packets out of the hotspot active.
// func (h HotspotActive) GetPacketsOut() int64 {
// 	result, ok := h["packets-out"]
// 	if !ok {
// 		return 0
// 	}
// 	resInt64, err := strconv.ParseInt(result, 10, 64)
// 	if err != nil {
// 		return 0
// 	}
// 	return resInt64
// }

// // GetComment returns the comment of the hotspot active.
// func (h HotspotActive) GetComment() string {
// 	result, ok := h["comment"]
// 	if !ok {
// 		return ""
// 	}
// 	return result
// }

// // GetIdleTime returns the idle time of the hotspot active.
// func (h HotspotActive) GetIdleTime() time.Duration {
// 	result, ok := h["idle-time"]
// 	if !ok {
// 		return time.Duration(0)
// 	}
// 	resultDur, err := time.ParseDuration(result)
// 	if err != nil {
// 		return time.Duration(0)
// 	}
// 	return resultDur
// }
