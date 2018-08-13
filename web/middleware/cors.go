package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

// Cors creates cors middleware.
func Cors() gin.HandlerFunc {
	return cors.New(CorsConfig())
}

// CorsConfig creates cors config.
func CorsConfig() cors.Config {
	return cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin"},
		MaxAge:           12 * time.Hour,
	}
}
