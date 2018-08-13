package env

import (
	"github.com/sutd-statnlp/service-ladrawex/constant"
	"github.com/sutd-statnlp/service-ladrawex/log"
)

var (
	// EnableProdMode enables production mode.
	EnableProdMode bool

	// ServerAddress is the server address for running web server.
	ServerAddress string
)

// init loads environment variables.
func init() {
	log.Debug("Init loads environment variables")
	EnableProdMode = Bool(constant.EnvEnableProd, false)
	ServerAddress = String(constant.EnvServerAddress, constant.ServerAddress)
}
