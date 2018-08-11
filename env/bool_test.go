package env_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/env"
)

func TestBool(test *testing.T) {
	assert.Equal(test, true, env.Bool("AA", true))
	assert.Equal(test, false, env.Bool("AA", false))

	os.Setenv("BB", "true")
	assert.Equal(test, true, env.Bool("BB", false))

	os.Setenv("BB", "btrue")
	assert.Equal(test, false, env.Bool("BB", false))
}
