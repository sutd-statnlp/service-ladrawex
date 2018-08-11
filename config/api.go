package config

import (
	"github.com/gin-gonic/gin"
)

// API attaches API routes into router.
func API(router *gin.Engine) []*gin.RouterGroup {
	apiGroup := router.Group("/api")
	{

	}
	return []*gin.RouterGroup{apiGroup}
}
