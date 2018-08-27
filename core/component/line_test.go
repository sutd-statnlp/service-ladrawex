package component_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/component"
)

func TestLine(t *testing.T) {
	line := new(component.Line)
	assert.NotNil(t, line)
	assert.Nil(t, line.Color)
	assert.Nil(t, line.StartPosition)
	assert.Nil(t, line.EndPosition)
	assert.NotNil(t, line.Width)
	assert.Equal(t, float32(0), line.Width)
	assert.False(t, line.Dashed)
}
