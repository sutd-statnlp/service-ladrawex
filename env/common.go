package env

import (
	"os"

	"github.com/sutd-statnlp/service-ladrawex/constant"
	"github.com/sutd-statnlp/service-ladrawex/util/stringutil"
)

var (
	// EnvMode is the mode of running environment.
	EnvMode string

	// BasePath is the project path.
	BasePath string
)

// init loads environment variables.
func init() {
	EnvMode = String(constant.EnvMode, constant.DevMode)

	defaultBasePath := stringutil.Concat(os.Getenv("GOPATH"), "/src/", constant.DefaultProjectPath)
	BasePath = String(constant.EnvBasePath, defaultBasePath)
}
