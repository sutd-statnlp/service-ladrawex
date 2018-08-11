package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/util"
)

func TestStringSliceContains(test *testing.T) {
	slice := []string{"GET", "POST", "PUT", "DELETE"}
	assert.True(test, util.StringSliceContains(slice, "GET"))
	assert.True(test, util.StringSliceContains(slice, "POST"))
	assert.True(test, util.StringSliceContains(slice, "GET", "POST", "PUT", "DELETE"))

	assert.False(test, util.StringSliceContains(slice, "HELLO"))
	assert.False(test, util.StringSliceContains(slice, ""))
	assert.False(test, util.StringSliceContains(slice))
	assert.False(test, util.StringSliceContains(slice, "GET", "POST", "PUT", "HELLO"))
}
