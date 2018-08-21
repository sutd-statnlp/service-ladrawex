package property_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/property"
)

func TestNewBorder(t *testing.T) {
	border := new(property.Border)
	assert.NotNil(t, border)
	assert.NotNil(t, border.Thick)
	assert.Equal(t, float32(0), border.Thick)
	assert.Nil(t, border.Color)
}
