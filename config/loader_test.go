package config_test

import (
	"testing"

	"github.com/spf13/viper"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/sutd-statnlp/service-ladrawex/config"
)

type LoaderTestSuite struct {
	suite.Suite
	loader *viper.Viper
}

func (suite *LoaderTestSuite) SetupTest() {
	suite.loader = config.NewLoader()
}

func (suite *LoaderTestSuite) TestMerge() {
	suite.True(config.Merge(suite.loader))
}

func (suite *LoaderTestSuite) TestLoad() {
	suite.True(config.Load(suite.loader))
}

func (suite *LoaderTestSuite) TestWatch() {
	suite.True(config.Watch(suite.loader))
}

func TestLoaderTestSuite(test *testing.T) {
	suite.Run(test, new(LoaderTestSuite))
}

func TestNewLoader(test *testing.T) {
	loader := config.NewLoader()
	assert.NotNil(test, loader)
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
	assert.NotNil(test, appConfig.Log)
	assert.NotNil(test, appConfig.Log.FilePath)
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
	assert.NotNil(test, appConfig.Log)
	assert.NotNil(test, appConfig.Log.FilePath)
}
