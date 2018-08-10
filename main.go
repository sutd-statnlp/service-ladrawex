package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/config"
	"github.com/sutd-statnlp/service-ladrawex/env"
)

func main() {
	enableProdMode, serverPort := env.LoadVariables()
	if enableProdMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	config.SetMiddleWares(router)
	config.SetAPIPaths(router)

	router.Run(serverPort)
}
