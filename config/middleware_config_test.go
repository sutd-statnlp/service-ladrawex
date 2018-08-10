package config_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/config"
)

// TestSetMiddlewares .
func TestSetMiddlewares(test *testing.T) {
	expect := config.SetMiddleWares(gin.New())
	assert.NotNil(test, expect)
}
