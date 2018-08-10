package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/middleware"
)

// TestGetCors .
func TestGetCors(test *testing.T) {
	expect := middleware.Cors()
	assert.NotNil(test, expect)
}
