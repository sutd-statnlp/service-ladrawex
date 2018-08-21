package stringutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/stringutil"
)

func TestConcat(t *testing.T) {
	result := stringutil.Concat("AA", "BB")
	assert.NotNil(t, result)
	assert.True(t, len(result) > 0)
	assert.Equal(t, "AABB", result)
}

func TestPrepare(t *testing.T) {
	result := stringutil.Prepare("? ? ?", 12, uint8(42), float32(22.2))
	assert.NotNil(t, result)
	assert.True(t, len(result) > 0)
	assert.Equal(t, "12 42 22.2", result)
}

func TestJSON(test *testing.T) {
	input := "Akaka"
	result := stringutil.JSON(&input)
	assert.NotEmpty(test, result)

	result = stringutil.JSON(nil)
	assert.Equal(test, "null", result)
}
