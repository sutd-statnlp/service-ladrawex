package constant

const (

	// EnvServerAddress is the environment key for specifying server address.
	EnvServerAddress = "SL_SERVER_ADDRESS"

	// EnvMode is the environment key for running environment mode.
	EnvMode = "SL_ENV_MODE"

	// EnvBasePath is the environment key for base path.
	EnvBasePath = "SL_BASE_PATH"
)

const (
	// DevEnv is the development environment.
	DevEnv = "dev"

	// ProdEnv is the production environment.
	ProdEnv = "prod"

	// DefaultServerAddress is the default web server address.
	DefaultServerAddress = ":9000"

	// DefaultProjectPath is the path of project.
	DefaultProjectPath = "github.com/sutd-statnlp/service-ladrawex"
)

const (
	// ConfigName is the default config file name.
	ConfigName = "ladrawex"

	// DevConfigName is the config file name for development.
	DevConfigName = "ladrawex-dev"

	// ProdConfigName is the config file name for production.
	ProdConfigName = "ladrawex-prod"

	// ConfigPath is the path for finding config files.
	ConfigPath = "/"

	// ConfigType is the type of config file.
	ConfigType = "yaml"
)
