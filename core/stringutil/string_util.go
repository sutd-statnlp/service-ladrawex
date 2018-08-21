package stringutil

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/spf13/cast"
)

// Concat concatinates all strings.
func Concat(args ...string) string {
	var buffer bytes.Buffer
	for _, item := range args {
		buffer.WriteString(item)
	}
	return buffer.String()
}

// Prepare prepares a string with different arguments.
func Prepare(inputString string, args ...interface{}) string {
	for _, item := range args {
		inputString = strings.Replace(inputString, "?", cast.ToString(item), 1)
	}
	return inputString
}

// JSON converts any type into json string.
func JSON(input interface{}) string {
	output, _ := json.Marshal(input)
	return string(output)
}
