package env

import (
	"os"
	"strconv"
)

// Int returns environment integer value.
// If the key does not exist, it will return integer value.
func Int(key string, defaultValue uint32) uint32 {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	intValue, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return defaultValue
	}
	return uint32(intValue)
}
