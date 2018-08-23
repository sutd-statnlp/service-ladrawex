package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/config"
)

// WebServer is the http server interface.
type WebServer interface {
	Config() bool
	Run()
}

var (
	webServer WebServer
)

func init() {
	webServer = New()
}

// Default returns default server.
func Default() WebServer {
	return webServer
}

// New creates new web server.
func New() WebServer {
	gin.SetMode(gin.ReleaseMode)
	return &GinWebServer{
		appConfig: config.Default(),
		engine:    gin.Default(),
	}
}
