package component_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/component"
)

func TestCreateNewRectangle(t *testing.T) {
	rect := component.NewRectangle()
	assert.NotNil(t, rect)
	assert.NotNil(t, rect.Border)
	assert.NotNil(t, rect.Color)
	assert.NotNil(t, rect.BackgroundColor)
	assert.NotNil(t, rect.Size)
	assert.NotNil(t, rect.Position)
}
