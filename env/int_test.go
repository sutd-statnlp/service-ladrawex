package env_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/env"
)

func TestInt(test *testing.T) {
	assert.NotNil(test, env.Int("AA", 2))
	assert.Equal(test, uint32(2), env.Int("AA", 2))
	assert.Equal(test, uint32(2), env.Int("", 2))

	os.Setenv("BB", "4")
	assert.Equal(test, uint32(4), env.Int("BB", 2))

	os.Setenv("BB", "b4")
	assert.Equal(test, uint32(2), env.Int("BB", 2))
}
