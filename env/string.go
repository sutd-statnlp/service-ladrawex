package env

import (
	"os"
)

// String gets environment string value from given key.
// If the key does not exist, it will return default value.
func String(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) > 0 {
		return value
	}
	return defaultValue
}
