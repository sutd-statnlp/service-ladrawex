package env

import (
	"os"

	"github.com/sutd-statnlp/service-ladrawex/log"
)

// String gets environment string value from given key.
// If the key does not exist, it will return default value.
func String(key, defaultValue string) string {
	log.Debug("Request to get env string value with: key=", key, " and defaultValue=", defaultValue)
	value := os.Getenv(key)
	if len(value) > 0 {
		return value
	}
	return defaultValue
}
