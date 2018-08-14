package config

// AppConfig is the app configuration.
type AppConfig interface {
	ServerAddress() *string
}

// AppConfigImpl is the app configuration implementation.
type AppConfigImpl struct {
	serverAddress *string
}

// ServerAddress returns server address.
func (config AppConfigImpl) ServerAddress() *string {
	return config.serverAddress
}
