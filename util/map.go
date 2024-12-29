package util

import (
	"strconv"
	"strings"

	"github.com/aidapedia/airouteros/types"
)

// SetTimeDuration sets the time duration value to the map.
func SetTimeDuration(result map[string]string, k string, v *types.AiTimeDuration) map[string]string {
	s := v.String()
	if s != types.None {
		result[k] = s
	}
	return result
}

// FindKeyToBool returns the boolean value of the key in the map.
func FindKeyToBool(m map[string]string, key string) bool {
	result, ok := m[key]
	if !ok {
		return false
	}
	return result == "true"
}

// FindKeyToString returns the string value of the key in the map.
func FindKeyToString(m map[string]string, key string) string {
	result, ok := m[key]
	if !ok {
		return ""
	}
	return result
}

// FindKeyToDuration returns the duration value of the key in the map.
func FindKeyToDuration(m map[string]string, key string) *types.AiTimeDuration {
	result, ok := m[key]
	if !ok {
		return nil
	}
	return types.ParseDuration(result)
}

// FindKeyToInt returns the integer value of the key in the map.
func FindKeyToInt(m map[string]string, key string) *types.AiInt {
	var v types.AiInt
	result, ok := m[key]
	if !ok {
		return &v
	}
	if result == types.Unlimited {
		return nil
	}
	i, err := strconv.Atoi(result)
	if err != nil {
		return &v
	}
	v = types.AiInt(i)
	return &v
}

// FindKeyToInt64 returns the integer value of the key in the map.
func FindKeyToInt64(m map[string]string, key string) *types.AiInt64 {
	var v types.AiInt64
	result, ok := m[key]
	if !ok {
		return &v
	}
	if result == types.Unlimited {
		return nil
	}
	i, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		return &v
	}
	v = types.AiInt64(i)
	return &v
}

// FindKeyToStringSlice returns the string slice value of the key in the map.
func FindKeyToStringSlice(m map[string]string, key string) []string {
	result, ok := m[key]
	if !ok {
		return []string{}
	}
	return strings.Split(result, ",")
}
