package env

import (
	"os"
	"strconv"

	"github.com/sutd-statnlp/service-ladrawex/log"
)

// Bool returns environment boolean value from given key.
// If the key does not exist, it returns default value.
func Bool(key string, defaultValue bool) bool {
	log.Debug("Request to get env bool value with: key=", key, " and defaultValue=", defaultValue)
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		log.Error(err)
		return defaultValue
	}
	return boolValue
}
