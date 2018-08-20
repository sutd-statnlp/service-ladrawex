package env

import (
	"os"

	"github.com/sutd-statnlp/service-ladrawex/constant"
	"github.com/sutd-statnlp/service-ladrawex/util/stringutil"
)

// String gets environment string value from given key.
// If the key does not exist, it will return default value.
func String(key, defaultValue string) string {
	key = stringutil.Concat(constant.EnvPrefix, "_", key)
	value := os.Getenv(key)
	if len(value) > 0 {
		return value
	}
	return defaultValue
}
