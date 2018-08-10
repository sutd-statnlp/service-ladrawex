package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/middleware"
)

// TestGetStatic .
func TestGetStatic(test *testing.T) {
	expect := middleware.Static()
	assert.NotNil(test, expect)
}
