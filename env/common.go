package env

import (
	"github.com/sutd-statnlp/service-ladrawex/constant"
	"github.com/sutd-statnlp/service-ladrawex/log"
)

var (
	// EnvMode is the mode of running environment.
	EnvMode string

	// ServerAddress is the server address for running web server.
	ServerAddress string
)

// init loads environment variables.
func init() {
	log.Debug("Init loads environment variables")
	EnvMode = String(constant.EnvMode, constant.DevEnv)
	ServerAddress = String(constant.EnvServerAddress, constant.DefaultServerAddress)
}
