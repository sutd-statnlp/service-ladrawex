package middleware_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/util/sliceutil"
	"github.com/sutd-statnlp/service-ladrawex/web/middleware"
)

func TestCors(test *testing.T) {
	cors := middleware.Cors()
	assert.NotNil(test, cors)
}

func TestCorsConfig(test *testing.T) {
	config := middleware.CorsConfig()
	assert.True(test, config.AllowAllOrigins)

	allowMethods := config.AllowMethods
	assert.True(test, sliceutil.ContainString(allowMethods, "GET", "POST", "PUT", "DELETE"))

	assert.True(test, config.AllowCredentials)

	allowHeaders := config.AllowHeaders
	assert.True(test, sliceutil.ContainString(allowHeaders, "Origin"))

	assert.True(test, config.MaxAge > time.Hour)

}
