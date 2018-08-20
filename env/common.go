package env

import (
	"os"

	"github.com/sutd-statnlp/service-ladrawex/constant"
	"github.com/sutd-statnlp/service-ladrawex/log"
)

var (
	// EnvMode is the mode of running environment.
	EnvMode string

	// ServerAddress is the server address for running web server.
	ServerAddress string

	// GoPath is the go path.
	GoPath string
)

// init loads environment variables.
func init() {
	log.Debug("Init loads environment variables")
	EnvMode = String(constant.EnvMode, constant.DevEnv)
	ServerAddress = String(constant.EnvServerAddress, constant.DefaultServerAddress)
	GoPath = os.Getenv("GOPATH")
}
