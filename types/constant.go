package types

type contextKey string

const (
	ContextReconnect contextKey = "reconnect"
)

type ActionMap string

const (
	ActionMapPrint  ActionMap = "print"
	ActionMapSet    ActionMap = "set"
	ActionMapAdd    ActionMap = "add"
	ActionMapRemove ActionMap = "remove"
)

const (
	None      = "none"
	Unlimited = "unlimited"
	Unset     = "unset"
)

type NetworkSpeedUnit string

const (
	NetworkSpeedUnitKbps NetworkSpeedUnit = "k"
	NetworkSpeedUnitMbps NetworkSpeedUnit = "M"
	NetworkSpeedUnitGbps NetworkSpeedUnit = "G"
)

type TimeUnit string

const (
	Second TimeUnit = "s"
	Minute TimeUnit = "m"
	Hour   TimeUnit = "h"
	Day    TimeUnit = "d"
	Week   TimeUnit = "w"
)
