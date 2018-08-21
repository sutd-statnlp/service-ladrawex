package property_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sutd-statnlp/service-ladrawex/core/property"
)

func TestNewColor(t *testing.T) {
	color := new(property.Color)
	assert.NotNil(t, color)
	assert.NotNil(t, color.R)
	assert.NotNil(t, color.G)
	assert.NotNil(t, color.B)
	assert.Equal(t, uint8(0), color.R)
	assert.Equal(t, uint8(0), color.G)
	assert.Equal(t, uint8(0), color.B)
}
