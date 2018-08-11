package env

import (
	"os"
	"strconv"
)

// Bool returns environment boolean value from given key.
// If the key does not exist, it returns default value.
func Bool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if len(key) == 0 {
		return defaultValue
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return boolValue
}
