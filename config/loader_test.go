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

func TestNew(test *testing.T) {
	assert.NotNil(test, config.New())
}
