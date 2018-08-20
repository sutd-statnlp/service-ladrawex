package config

// WebConfig is the web configuration.
type WebConfig struct {
	ServerAddress string
	StaticPath    string
	Middleware    *MiddlewareConfig
}

// MiddlewareConfig is the middleware configuration.
type MiddlewareConfig struct {
	Cors   *CorsConfig
	Gzip   *GzipConfig
	Static *StaticConfig
}

// CorsConfig is the cors configuration.
type CorsConfig struct {
	Enable bool
}

// GzipConfig is the gzip configuration.
type GzipConfig struct {
	Enable bool
}

// StaticConfig is the static configuration.
type StaticConfig struct {
	Enable bool
}
