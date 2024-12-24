package model

import (
	"strconv"
	"time"
)

type Hotspot map[string]string

// GetID returns the id of the hotspot active.
func (h Hotspot) GetID() string {
	result, ok := h[".id"]
	if !ok {
		return ""
	}
	return result
}

// GetInterface returns the interface of the hotspot active.
func (h Hotspot) GetInterface() string {
	result, ok := h["interface"]
	if !ok {
		return ""
	}
	return result
}

// GetInvalid returns the invalid of the hotspot active.
func (h Hotspot) GetInvalid() bool {
	result, ok := h["invalid"]
	if !ok {
		return false
	}
	return result == "true"
}

// GetDisabled returns the disabled of the hotspot active.
func (h Hotspot) GetDisabled() bool {
	result, ok := h["disabled"]
	if !ok {
		return false
	}
	return result == "true"
}

// GetName returns the name of the hotspot active.
func (h Hotspot) GetName() string {
	result, ok := h["name"]
	if !ok {
		return ""
	}
	return result
}

// GetLoginTimeout returns the login-timeout of the hotspot active.
func (h Hotspot) GetLoginTimeout() time.Duration {
	result, ok := h["login-timeout"]
	if !ok || result == "none" {
		return time.Duration(0)
	}
	i, _ := strconv.Atoi(result)
	return time.Duration(i) * time.Second
}

// GetKeepAliveTimeout returns the keepalive-timeout of the hotspot active.
func (h Hotspot) GetKeepAliveTimeout() time.Duration {
	result, ok := h["keepalive-timeout"]
	if !ok || result == "none" {
		return time.Duration(0)
	}
	i, _ := strconv.Atoi(result)
	return time.Duration(i) * time.Second
}

// GetIPofDNSName returns the ip of dns-name of the hotspot active.
func (h Hotspot) GetIPofDNSName() string {
	result, ok := h["ip-of-dns-name"]
	if !ok {
		return ""
	}
	return result
}

// GetAddressesPerMAC returns the addresses-per-mac of the hotspot active.
func (h Hotspot) GetAddressesPerMAC() int {
	result, ok := h["addresses-per-mac"]
	if !ok {
		return 0
	}
	if result == "unlimited" {
		return -1
	}
	i, _ := strconv.Atoi(result)
	return i
}

// GetHTTPS returns the https of the hotspot active.
func (h Hotspot) GetHTTPS() bool {
	result, ok := h["https"]
	if !ok {
		return false
	}
	return result == "true"
}

// GetProfile returns the profile of the hotspot active.
func (h Hotspot) GetProfile() string {
	result, ok := h["profile"]
	if !ok {
		return ""
	}
	return result
}

// GetIdleTimeout returns the idle-timeout of the hotspot active.
func (h Hotspot) GetIdleTimeout() time.Duration {
	result, ok := h["idle-timeout"]
	if !ok || result == "none" {
		return time.Duration(0)
	}
	i, _ := strconv.Atoi(result)
	return time.Duration(i) * time.Second
}

// GetProxyStatus returns the proxy-status of the hotspot active.
func (h Hotspot) GetProxyStatus() string {
	result, ok := h["proxy-status"]
	if !ok {
		return ""
	}
	return result
}
