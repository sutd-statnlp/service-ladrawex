package constant

const (

	// EnvPrefix is the prefix of environment key.
	EnvPrefix = "SL"

	// EnvWebServerAddress is the environment key for specifying server address.
	EnvWebServerAddress = "WEB_SERVER_ADDRESS"

	// EnvMode is the environment key for running environment mode.
	EnvMode = "MODE"

	// EnvBasePath is the environment key for base path.
	EnvBasePath = "BASE_PATH"

	// EnvLogLevel is the environment key for log level.
	EnvLogLevel = "LOG_LEVEL"
)

const (
	// DevMode is the development mode.
	DevMode = "dev"

	// ProdMode is the production mode.
	ProdMode = "prod"

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
