package property_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/property"
)

func TestNewPosition(t *testing.T) {
	position := new(property.Position)
	assert.NotNil(t, position)
	assert.NotNil(t, position.X)
	assert.Equal(t, float32(0), position.X)
	assert.NotNil(t, position.Y)
	assert.Equal(t, float32(0), position.Y)
}
