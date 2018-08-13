package middleware

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/constant"
)

// Static creates static middleware.
func Static() gin.HandlerFunc {
	return static.Serve("/", static.LocalFile(constant.WebViewPath, true))
}
