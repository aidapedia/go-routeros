package module

import (
	"errors"
	"strconv"
	"strings"

	"github.com/aidapedia/go-routeros/types"
	"github.com/aidapedia/go-routeros/util"
)

const (
	HotspotProfileModule           Module = "/ip/hotspot/profile"
	ErrInvalidActionHotspotProfile        = "invalid action for HotspotProfile"
)

func init() {
	register(HotspotProfileModule, &HotspotProfile{})
}

type HotspotProfile struct{}

// GetData returns the data of the module.
func (h *HotspotProfile) GetData() ModuleDataInterface {
	return &HotspotProfileData{}
}

// GetQueryPath returns the path of the module.
func (h *HotspotProfile) GetQueryPath() string {
	return string(HotspotProfileModule)
}

// CheckAction checks if the action is valid for the module.
func (h *HotspotProfile) CheckAction(action types.ActionMap) error {
	switch action {
	case types.ActionMapPrint:
		return nil
	default:
		return errors.New(ErrInvalidActionHotspotProfile)
	}
}

// HotspotProfileData is a representation of a logged-in HotSpot user.
type HotspotProfileData struct {
	ID                    string
	HotspotAddress        string
	LoginBy               []string
	HTTPCookieLifetime    *types.AiTimeDuration
	SplitUserDomain       bool
	Default               bool
	UseRadius             bool
	HTMLDirectoryOverride string
	SMTPServer            string
	Name                  string
	DNSName               string
	HTMLDirectory         string
	InstallHotspotQueue   string
	HTTPProxy             string
	RateLimit             *types.NetworkLimit
}

func (h *HotspotProfileData) UnmarshalMap(m map[string]string) error {
	h.ID = m[".id"]
	h.HotspotAddress = m["hotspot-address"]
	h.LoginBy = strings.Split(m["login-by"], ",")
	h.HTTPCookieLifetime = util.FindKeyToDuration(m, "http-cookie-lifetime")
	h.SplitUserDomain = util.FindKeyToBool(m, "split-user-domain")
	h.Default = util.FindKeyToBool(m, "default")
	h.UseRadius = util.FindKeyToBool(m, "use-radius")
	h.HTMLDirectoryOverride = m["html-directory-override"]
	h.SMTPServer = m["smtp-server"]
	h.Name = m["name"]
	h.DNSName = m["dns-name"]
	h.HTMLDirectory = m["html-directory"]
	h.InstallHotspotQueue = m["install-hotspot-queue"]
	h.HTTPProxy = m["http-proxy"]
	h.RateLimit = types.ParseNetworkLimit(m["rate-limit"])
	return nil
}

func (h *HotspotProfileData) MarshalMap(action types.ActionMap) (map[string]string, error) {
	switch action {

	case types.ActionMapPrint:
		result := map[string]string{}
		result[".id"] = h.ID
		result["hotspot-address"] = h.HotspotAddress
		result["login-by"] = strings.Join(h.LoginBy, ",")
		result["http-cookie-lifetime"] = h.HTTPCookieLifetime.String()
		result["split-user-domain"] = strconv.FormatBool(h.SplitUserDomain)
		result["default"] = strconv.FormatBool(h.Default)
		result["use-radius"] = strconv.FormatBool(h.UseRadius)
		result["html-directory-override"] = h.HTMLDirectoryOverride
		result["smtp-server"] = h.SMTPServer
		result["name"] = h.Name
		result["dns-name"] = h.DNSName
		result["html-directory"] = h.HTMLDirectory
		result["install-hotspot-queue"] = h.InstallHotspotQueue
		result["http-proxy"] = h.HTTPProxy
		result["rate-limit"] = h.RateLimit.String()
		return result, nil
	default:
		return map[string]string{}, errors.New(ErrInvalidActionHotspotProfile)
	}
}
