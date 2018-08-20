package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/config"
)

func TestLoad(test *testing.T) {
	assert.True(test, config.Load())
}

func TestWatch(test *testing.T) {
	assert.True(test, config.Watch())
}

func TestMerge(test *testing.T) {
	assert.True(test, config.Merge())
}

func TestDefault(test *testing.T) {
	appConfig := config.Default()
	assert.NotNil(test, appConfig)
	assert.NotNil(test, appConfig.Web)
	assert.NotNil(test, appConfig.Web.ServerAddress)
	assert.NotNil(test, appConfig.Web.Middleware)
	assert.NotNil(test, appConfig.Web.Middleware.Cors)
	assert.NotNil(test, appConfig.Web.Middleware.Gzip)
	assert.NotNil(test, appConfig.Web.Middleware.Static)
}

func TestNew(test *testing.T) {
	appConfig := config.New()
	assert.NotNil(test, appConfig)
	assert.NotNil(test, appConfig.Web)
	assert.NotNil(test, appConfig.Web.ServerAddress)
	assert.NotNil(test, appConfig.Web.Middleware)
	assert.NotNil(test, appConfig.Web.Middleware.Cors)
	assert.NotNil(test, appConfig.Web.Middleware.Gzip)
	assert.NotNil(test, appConfig.Web.Middleware.Static)
}
