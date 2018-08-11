package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/middleware"
)

// Middleware attaches middlewares to router.
func Middleware(router *gin.Engine) gin.IRoutes {
	return router.Use(
		middleware.Cors(),
		middleware.Gzip(),
		middleware.Static(),
	)
}
