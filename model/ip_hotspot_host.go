package model

import (
	"strconv"
	"time"
)

type HotspotHost map[string]string

// GetID returns the id of the hotspot host.
func (h HotspotHost) GetID() string {
	result, ok := h[".id"]
	if !ok {
		return ""
	}
	return result
}

// GetMacAddress returns the mac address of the hotspot host.
func (h HotspotHost) GetMacAddress() string {
	result, ok := h["mac-address"]
	if !ok {
		return ""
	}
	return result
}

// GetAddress returns the address of the hotspot host.
func (h HotspotHost) GetAddress() string {
	result, ok := h["address"]
	if !ok {
		return ""
	}
	return result
}

// GetToAddress returns the to address of the hotspot host.
func (h HotspotHost) GetToAddress() string {
	result, ok := h["to-address"]
	if !ok {
		return ""
	}
	return result
}

// GetServer returns the server of the hotspot host.
func (h HotspotHost) GetServer() string {
	result, ok := h["server"]
	if !ok {
		return ""
	}
	return result
}

// GetUptime returns the uptime of the hotspot host.
func (h HotspotHost) GetUptime() time.Duration {
	result, ok := h["uptime"]
	if !ok {
		return time.Duration(0)
	}
	resultDur, err := time.ParseDuration(result)
	if err != nil {
		return time.Duration(0)
	}
	return resultDur
}

// GetBytesIn returns the bytes in of the hotspot host.
func (h HotspotHost) GetBytesIn() int64 {
	result, ok := h["bytes-in"]
	if !ok {
		return 0
	}
	resInt64, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		return 0
	}
	return resInt64
}

// GetBytesOut returns the bytes out of the hotspot host.
func (h HotspotHost) GetBytesOut() int64 {
	result, ok := h["bytes-out"]
	if !ok {
		return 0
	}
	resInt64, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		return 0
	}
	return resInt64
}

// GetPacketsIn returns the packets in of the hotspot host.
func (h HotspotHost) GetPacketsIn() int64 {
	result, ok := h["packets-in"]
	if !ok {
		return 0
	}
	resInt64, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		return 0
	}
	return resInt64
}

// GetPacketsOut returns the packets out of the hotspot host.
func (h HotspotHost) GetPacketsOut() int64 {
	result, ok := h["packets-out"]
	if !ok {
		return 0
	}
	resInt64, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		return 0
	}
	return resInt64
}

// GetIdleTimeout returns the idle time out of the hotspot host.
func (h HotspotHost) GetIdleTimeout() time.Duration {
	result, ok := h["idle-timeout"]
	if !ok {
		return time.Duration(0)
	}
	resultDur, err := time.ParseDuration(result)
	if err != nil {
		return time.Duration(0)
	}
	return resultDur
}

// GetIdleTime returns the idle time of the hotspot host.
func (h HotspotHost) GetIdleTime() time.Duration {
	result, ok := h["idle-time"]
	if !ok {
		return time.Duration(0)
	}
	resultDur, err := time.ParseDuration(result)
	if err != nil {
		return time.Duration(0)
	}
	return resultDur
}

// GetBypassed returns the bypassed of the hotspot host.
func (h HotspotHost) GetBypassed() bool {
	result, ok := h["bypassed"]
	if !ok {
		return false
	}
	return result == "true"
}

// GetHostDeadTime returns the host dead time of the hotspot host.
func (h HotspotHost) GetHostDeadTime() time.Duration {
	result, ok := h["host-dead-time"]
	if !ok {
		return time.Duration(0)
	}
	resultDur, err := time.ParseDuration(result)
	if err != nil {
		return time.Duration(0)
	}
	return resultDur
}

// GetFoundBy returns the found by of the hotspot host.
func (h HotspotHost) GetFoundBy() string {
	result, ok := h["found-by"]
	if !ok {
		return ""
	}
	return result
}

// GetDHCP returns the dhcp of the hotspot host.
func (h HotspotHost) GetDHCP() bool {
	result, ok := h["dhcp"]
	if !ok {
		return false
	}
	return result == "true"
}

// GetAuthorized returns the authorized of the hotspot host.
func (h HotspotHost) GetAuthorized() bool {
	result, ok := h["authorized"]
	if !ok {
		return false
	}
	return result == "true"
}

// GetComment returns the comment of the hotspot host.
func (h HotspotHost) GetComment() string {
	result, ok := h["comment"]
	if !ok {
		return ""
	}
	return result
}
