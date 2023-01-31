package utils

import (
	"strings"
)

// ParseMapFromString is used for converting name/value pairs
// from command line arguments. For example adding tags to Cloud
// assets
func ParseMapFromString(in string) (out map[string]string) {
	out = make(map[string]string, 0)

	for _, kv := range strings.Split(in, ";") {
		kv1 := strings.Split(kv, "=")
		if len(kv1) == 2 {
			out[kv1[0]] = kv1[1]
		}
	}
	return out
}
