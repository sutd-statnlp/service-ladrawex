package testutil

import (
	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/web/server"
)

// CreateFakeServer returns the fake router for testing.
func CreateFakeServer() server.WebServer {
	gin.SetMode(gin.TestMode)
	return server.New()
}
