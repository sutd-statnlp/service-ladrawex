package config

// AppConfig is the app configuration.
type AppConfig struct {
	Mode   string
	Web    *WebConfig
	Log    *LogConfig
	Loader *LoaderConfig
}
