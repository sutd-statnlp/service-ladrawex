package middleware

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/log"
)

// Gzip creates gzip middleware.
func Gzip() gin.HandlerFunc {
	log.Debug("Request to create gzip middleware")
	return gzip.Gzip(GzipLevel())
}

// GzipLevel returns gzip level.
func GzipLevel() int {
	return gzip.DefaultCompression
}
