package middleware_test

import (
	"testing"

	"github.com/gin-contrib/gzip"
	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/web/middleware"
)

func TestGetGzip(test *testing.T) {
	gzip := middleware.Gzip()
	assert.NotNil(test, gzip)
}

func TestGzipLevel(test *testing.T) {
	gzipLevel := middleware.GzipLevel()
	assert.Equal(test, gzip.DefaultCompression, gzipLevel)
}
