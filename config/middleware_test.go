package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/config"
	"github.com/sutd-statnlp/service-ladrawex/util"
)

func TestSetMiddlewares(test *testing.T) {
	routes := config.Middleware(util.CreateFakeRouter())
	assert.NotNil(test, routes)
}
