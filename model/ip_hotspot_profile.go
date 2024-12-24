package model

import (
	"strings"
	"time"
)

type HotspotProfile map[string]string

// GetID returns the id of the hotspot profile.
func (h HotspotProfile) GetID() string {
	result, ok := h[".id"]
	if !ok {
		return ""
	}
	return result
}

// GetHostspotAddress returns the hotspot address of the hotspot profile.
func (h HotspotProfile) GetHostspotAddress() string {
	result, ok := h["hotspot-address"]
	if !ok {
		return ""
	}
	return result
}

// GetLoginBy returns the login by of the hotspot profile.
func (h HotspotProfile) GetLoginBy() []string {
	result, ok := h["login-by"]
	if !ok {
		return []string{}
	}
	return strings.Split(result, ",")
}

// GetHTTPCookieLifetime returns the http cookie lifetime of the hotspot profile.
func (h HotspotProfile) GetHTTPCookieLifetime() time.Duration {
	result, ok := h["http-cookie-lifetime"]
	if !ok || result == "none" {
		return time.Duration(0)
	}
	duration, err := time.ParseDuration(result)
	if err != nil {
		return time.Duration(0)
	}
	return duration
}

// GetSplitUserDomain returns the split user domain of the hotspot profile.
func (h HotspotProfile) GetSplitUserDomain() bool {
	result, ok := h["split-user-domain"]
	if !ok {
		return false
	}
	return result == "true"
}

// GetDefault returns the default of the hotspot profile.
func (h HotspotProfile) GetDefault() bool {
	result, ok := h["default"]
	if !ok {
		return false
	}
	return result == "true"
}

// GetUseRadius returns the use radius of the hotspot profile.
func (h HotspotProfile) GetUseRadius() bool {
	result, ok := h["use-radius"]
	if !ok {
		return false
	}
	return result == "true"
}

// GetHTMLDirectoryOverride returns the html directory override of the hotspot profile.
func (h HotspotProfile) GetHTMLDirectoryOverride() string {
	result, ok := h["html-directory-override"]
	if !ok {
		return ""
	}
	return result
}

// GetSMTPServer returns the smtp server of the hotspot profile.
func (h HotspotProfile) GetSMTPServer() string {
	result, ok := h["smtp-server"]
	if !ok {
		return ""
	}
	return result
}

// GetName returns the name of the hotspot profile.
func (h HotspotProfile) GetName() string {
	result, ok := h["name"]
	if !ok {
		return ""
	}
	return result
}

// GetDNSName returns the dns name of the hotspot profile.
func (h HotspotProfile) GetDNSName() string {
	result, ok := h["dns-name"]
	if !ok {
		return ""
	}
	return result
}

// GetHTMLDirectory returns the html directory of the hotspot profile.
func (h HotspotProfile) GetHTMLDirectory() string {
	result, ok := h["html-directory"]
	if !ok {
		return ""
	}
	return result
}

// GetInstallHotspotQueue returns the install hotspot queue of the hotspot profile.
func (h HotspotProfile) GetInstallHotspotQueue() string {
	result, ok := h["install-hotspot-queue"]
	if !ok {
		return ""
	}
	return result
}

// GetHTTPProxy returns the http proxy of the hotspot profile.
func (h HotspotProfile) GetHTTPProxy() string {
	result, ok := h["http-proxy"]
	if !ok {
		return ""
	}
	return result
}

// GetRateLimit returns the rate limit of the hotspot profile.
func (h HotspotProfile) GetRateLimit() string {
	result, ok := h["rate-limit"]
	if !ok {
		return ""
	}
	return result
}
