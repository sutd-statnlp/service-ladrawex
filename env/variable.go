package env

import (
	"os"
)

// LoadVariables gets environment variables: enableProdMode and serverHost.
func LoadVariables() (bool, string) {
	var (
		enableProdMode = false
		serverHost     = ":9000"
	)
	if os.Getenv("GBP_PROMODE") == "true" {
		enableProdMode = true
	}

	host := os.Getenv("GBP_HOST")
	if len(host) > 0 {
		serverHost = host
	}
	return enableProdMode, serverHost
}
