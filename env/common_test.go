package env_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/constant"
	"github.com/sutd-statnlp/service-ladrawex/env"
)

func TestEnableProdMode(test *testing.T) {
	assert.False(test, env.EnableProdMode)
}

func TestServerAddress(test *testing.T) {
	assert.Equal(test, env.ServerAddress, constant.ServerAddress)
}
