package stringutil

import (
	"bytes"
)

// Concat concatinates all strings.
func Concat(args ...string) string {
	var buffer bytes.Buffer
	for _, item := range args {
		buffer.WriteString(item)
	}
	return buffer.String()
}
