package util

import (
	"github.com/gin-gonic/gin"
)

// CreateFakeRouter returns the fake router for testing.
func CreateFakeRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}
