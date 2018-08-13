package api

import (
	"github.com/gin-gonic/gin"
)

// Group returns the api group.
func Group() *gin.RouterGroup {
	apiGroup := gin.RouterGroup{}
	return &apiGroup
}
