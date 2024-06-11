package strUtils

import (
	"strings"
)

func TrimSuffix(str, suffix string) string {
	if strings.HasSuffix(str, suffix) {
		return str[:len(str)-len(suffix)]
	}

	return str
}
