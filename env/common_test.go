package env_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/constant"
	"github.com/sutd-statnlp/service-ladrawex/env"
)

func TestEnvMode(test *testing.T) {
	assert.Equal(test, env.EnvMode, constant.DevMode)
}

func TestBasePath(test *testing.T) {
	assert.NotNil(test, env.BasePath)
	assert.True(test, len(env.BasePath) > 0)
}
