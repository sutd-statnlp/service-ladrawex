package middleware

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// Static creates static middleware.
func Static() gin.HandlerFunc {
	return static.Serve("/", static.LocalFile("./view", true))
}
