package main

import (
	"github.com/sutd-statnlp/service-ladrawex/config"
	"github.com/sutd-statnlp/service-ladrawex/web/server"
)

func main() {
	appConfig := config.Default()
	var webServer server.WebServer

	webServer = server.GinWebServer{
		ServerAddress: appConfig.Web.ServerAddress,
	}

	webServer.Run()
}
