package main

import (
	"github.com/sutd-statnlp/service-ladrawex/env"
	"github.com/sutd-statnlp/service-ladrawex/web/server"
)

func main() {
	var webServer server.WebServer

	webServer = server.GinWebServer{
		ServerAddress: env.ServerAddress(),
	}

	webServer.Run()
}
