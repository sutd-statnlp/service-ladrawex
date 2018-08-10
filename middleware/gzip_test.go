package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/middleware"
)

// TestGetGzip .
func TestGetGzip(test *testing.T) {
	expect := middleware.Gzip()
	assert.NotNil(test, expect)
}
