package middleware

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/constant"
	"github.com/sutd-statnlp/service-ladrawex/log"
)

// Static creates static middleware.
func Static() gin.HandlerFunc {
	log.Debug("Request to create static middleware")
	return static.Serve("/", static.LocalFile(constant.WebViewPath, true))
}
