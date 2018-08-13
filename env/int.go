package env

import (
	"os"
	"strconv"

	"github.com/sutd-statnlp/service-ladrawex/log"
)

// Int returns environment integer value.
// If the key does not exist, it will return integer value.
func Int(key string, defaultValue uint32) uint32 {
	log.Debug("Request to get env init value with: key=", key, " and defaultValue=", defaultValue)
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	intValue, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		log.Error(err)
		return defaultValue
	}
	return uint32(intValue)
}
