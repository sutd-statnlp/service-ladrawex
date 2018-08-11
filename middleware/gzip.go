package middleware

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// Gzip creates gzip middleware.
func Gzip() gin.HandlerFunc {
	return gzip.Gzip(GzipLevel())
}

// GzipLevel returns gzip level.
func GzipLevel() int {
	return gzip.DefaultCompression
}
