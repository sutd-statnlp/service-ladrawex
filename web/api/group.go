package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/log"
)

// Group returns the api group.
func Group() *gin.RouterGroup {
	log.Debug("Request to create API group")
	apiGroup := gin.RouterGroup{}
	return &apiGroup
}
