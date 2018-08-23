package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/log"

	"github.com/gin-contrib/cors"
)

// Cors creates cors middleware.
func Cors() gin.HandlerFunc {
	log.Debug("Request to create cors middleware")
	return cors.New(CorsConfig())
}

// CorsConfig creates cors config.
func CorsConfig() cors.Config {
	return cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "content-type", "accept"},
		MaxAge:           12 * time.Hour,
	}
}
