package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/config"
	"github.com/sutd-statnlp/service-ladrawex/log"
	"github.com/sutd-statnlp/service-ladrawex/util/fileutil"
	"github.com/sutd-statnlp/service-ladrawex/web/api"
	"github.com/sutd-statnlp/service-ladrawex/web/middleware"
)

// GinWebServer is the implementation of WebServer interface by using Gin.
type GinWebServer struct {
	appConfig *config.AppConfig
	engine    *gin.Engine
}

// Config configures web server before starting.
func (webServer GinWebServer) Config() bool {
	gin.SetMode(gin.ReleaseMode)

	middlewareConfig := webServer.appConfig.Web.Middleware
	if middlewareConfig.Cors.Enable {
		webServer.engine.Use(middleware.Cors())
	}
	if middlewareConfig.Gzip.Enable {
		webServer.engine.Use(middleware.Gzip())
	}
	if middlewareConfig.Static.Enable {
		staticPath := fileutil.FullPath(webServer.appConfig.Web.StaticPath)
		webServer.engine.Use(middleware.Static(staticPath))
	}

	webServer.engine.RouterGroup = *api.Group()
	return webServer.engine != nil
}

// Run starts web server with Gin Engine.
func (webServer GinWebServer) Run() {
	if !webServer.Config() {
		log.Panic("Web server have configuration error!")
	}
	serverAddress := webServer.appConfig.Web.ServerAddress
	log.Info("Web server is configured and running at ", serverAddress)
	err := webServer.engine.Run(serverAddress)
	if err != nil {
		log.Panic(err)
	}
}
