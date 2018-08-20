package env_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/env"
)

func TestString(test *testing.T) {
	assert.NotNil(test, env.String("AA", "default value"))
	assert.Equal(test, "value", env.String("AA", "value"))
	assert.Equal(test, "value", env.String("", "value"))

	os.Setenv("SL_BB", "bb")
	assert.Equal(test, "bb", env.String("BB", "value"))
}
