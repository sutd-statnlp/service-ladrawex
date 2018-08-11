package middleware_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/middleware"
	"github.com/sutd-statnlp/service-ladrawex/util"
)

func TestCors(test *testing.T) {
	cors := middleware.Cors()
	assert.NotNil(test, cors)
}

func TestCorsConfig(test *testing.T) {
	config := middleware.CorsConfig()
	assert.True(test, config.AllowAllOrigins)

	allowMethods := config.AllowMethods
	assert.True(test, util.StringSliceContains(allowMethods, "GET", "POST", "PUT", "DELETE"))

	assert.True(test, config.AllowCredentials)

	allowHeaders := config.AllowHeaders
	assert.True(test, util.StringSliceContains(allowHeaders, "Origin"))

	assert.True(test, config.MaxAge > time.Hour)

}
