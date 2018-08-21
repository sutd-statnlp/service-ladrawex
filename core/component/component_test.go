package component_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/component"
)

func TestCreateNewRectangle(t *testing.T) {
	rect := component.NewRectangle()
	assert.NotNil(t, rect)
	assert.NotNil(t, rect.Common)
}

func TestNewCommonProperty(t *testing.T) {
	common := component.NewCommonProperty()
	assert.NotNil(t, common)
	assert.NotNil(t, common.Border)
	assert.NotNil(t, common.BackgroundColor)
	assert.NotNil(t, common.Size)
	assert.NotNil(t, common.Position)
}
