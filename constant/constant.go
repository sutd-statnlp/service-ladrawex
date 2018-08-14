package constant

const (

	// EnvServerAddress is the environment key for specifying server address.
	EnvServerAddress = "SL_SERVER_ADDRESS"

	// EnvMode is the environment key for running environment mode.
	EnvMode = "SL_ENV_MODE"

	// DevEnv is the development environment.
	DevEnv = "dev"

	// ProdEnv is the production environment.
	ProdEnv = "prod"

	// LogFilePath is the the path of logging file.
	LogFilePath = "./default.log"

	// DefaultServerAddress is the default web server address.
	DefaultServerAddress = ":9000"

	// WebViewPath is the path of web ui.
	WebViewPath = "./web/view"

	// DevConfigName is the config file name for development.
	DevConfigName = "ladrawex"

	// ProdConfigName is the config file name for production.
	ProdConfigName = "ladrawex-prod"

	// ConfigPath is the path for finding config files.
	ConfigPath = "../"

	// ConfigType is the type of config file.
	ConfigType = "yaml"
)
