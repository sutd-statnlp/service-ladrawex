package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/web/middleware"
)

func TestGetStatic(test *testing.T) {
	expect := middleware.Static("/path")
	assert.NotNil(test, expect)
}
