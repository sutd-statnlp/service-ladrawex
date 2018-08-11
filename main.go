package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/config"
	"github.com/sutd-statnlp/service-ladrawex/env"
)

func main() {

	if env.EnableProdMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	config.Middleware(router)
	config.API(router)

	router.Run(env.ServerAddress)
}
