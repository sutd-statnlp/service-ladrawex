package env_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/constant"
	"github.com/sutd-statnlp/service-ladrawex/env"
)

func TestEnvMode(test *testing.T) {
	assert.Equal(test, env.EnvMode, constant.DevEnv)
}

func TestServerAddress(test *testing.T) {
	assert.Equal(test, env.ServerAddress, constant.DefaultServerAddress)
}

func TestBasePath(test *testing.T) {
	assert.NotNil(test, env.BasePath)
	assert.True(test, len(env.BasePath) > 0)
}
