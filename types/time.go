package types

import (
	"fmt"
	"strconv"
	"strings"
)

type TimeUnit string

const (
	Second TimeUnit = "s"
	Minute TimeUnit = "m"
	Hour   TimeUnit = "h"
	Day    TimeUnit = "d"
	Week   TimeUnit = "w"
)

var (
	duration = map[TimeUnit]int64{
		Week:   604800,
		Day:    86400,
		Hour:   3600,
		Minute: 60,
		Second: 1,
	}
	unit = []TimeUnit{
		Week, Day, Hour, Minute, Second,
	}
)

type AiTimeDuration int64

func (a *AiTimeDuration) String() string {
	if a == nil {
		return "none"
	}
	var res string
	val := int64(*a)
	for _, u := range unit {
		d := duration[u]
		r := val / d
		if r > 0 {
			res = fmt.Sprintf("%s%d%s", res, r, u)
			val = val - (r * d)
		}
	}
	return res
}

func ParseDuration(s string) *AiTimeDuration {
	if s == "none" {
		return nil
	}
	var res AiTimeDuration
	for _, u := range unit {
		d := duration[u]
		splitS := strings.Split(s, string(u))
		if len(splitS) == 2 {
			v, err := strconv.ParseInt(splitS[0], 10, 64)
			if err != nil {
				return nil
			}
			res = res + AiTimeDuration(v*d)
			s = s[:len(splitS[0])]
		}
	}
	return &res
}
