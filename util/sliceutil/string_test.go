package sliceutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/util/sliceutil"
)

func TestContainString(test *testing.T) {
	slice := []string{"GET", "POST", "PUT", "DELETE"}
	assert.True(test, sliceutil.ContainString(slice, "GET"))
	assert.True(test, sliceutil.ContainString(slice, "POST"))
	assert.True(test, sliceutil.ContainString(slice, "GET", "POST", "PUT", "DELETE"))

	assert.False(test, sliceutil.ContainString(slice, "HELLO"))
	assert.False(test, sliceutil.ContainString(slice, ""))
	assert.False(test, sliceutil.ContainString(slice))
	assert.False(test, sliceutil.ContainString(slice, "GET", "POST", "PUT", "HELLO"))
}
