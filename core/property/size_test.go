package property_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/property"
)

func TestNewSize(t *testing.T) {
	size := new(property.Size)
	assert.NotNil(t, size)
	assert.NotNil(t, size.Width)
	assert.Equal(t, float32(0), size.Width)
	assert.NotNil(t, size.Height)
	assert.Equal(t, float32(0), size.Height)
}
