package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/log"
	"github.com/sutd-statnlp/service-ladrawex/web/api"
	"github.com/sutd-statnlp/service-ladrawex/web/middleware"
)

// GinWebServer is the implementation of WebServer interface by using Gin.
type GinWebServer struct {
	ServerAddress string
	engine        *gin.Engine
}

// Config configures web server before starting.
func (webServer GinWebServer) Config() bool {
	gin.SetMode(gin.ReleaseMode)

	webServer.engine = gin.Default()
	webServer.engine.Use(
		middleware.Cors(),
		middleware.Gzip(),
		middleware.Static(),
	)
	webServer.engine.RouterGroup = *api.Group()
	return webServer.engine != nil
}

// Run starts web server with Gin Engine.
func (webServer GinWebServer) Run() {
	if !webServer.Config() {
		log.Panic("Web server have configuration error!")
	}
	log.Info("Web server is configured and running at ", webServer.ServerAddress)
	err := webServer.engine.Run(webServer.ServerAddress)
	if err != nil {
		log.Panic(err)
	}
}
