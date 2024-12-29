package types

import (
	"fmt"
	"strings"
)

type NetworkSpeed struct {
	Amount int
	Unit   NetworkSpeedUnit
}

func ParseNetworkLimit(limit string) *NetworkLimit {
	var download, upload int
	var downloadUnit, uploadUnit string

	speed := strings.Split(limit, "/")
	if limit == "" {
		return nil
	}
	fmt.Sscanf(speed[0], "%d%s", &download, &downloadUnit)
	fmt.Sscanf(speed[1], "%d%s", &upload, &uploadUnit)
	res := &NetworkLimit{
		Download: NetworkSpeed{
			Amount: download,
			Unit:   NetworkSpeedUnit(downloadUnit),
		}, Upload: NetworkSpeed{
			Amount: upload,
			Unit:   NetworkSpeedUnit(uploadUnit),
		},
	}
	return res
}

func (n NetworkSpeed) String() string {
	return fmt.Sprintf("%d%s", n.Amount, n.Unit)
}

type NetworkLimit struct {
	Upload   NetworkSpeed
	Download NetworkSpeed
}

func (n *NetworkLimit) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("%s/%s", n.Download.String(), n.Upload.String())
}

// Custom Type
