package component_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/component"
)

func TestRectangle(t *testing.T) {
	rect := new(component.Rectangle)
	assert.NotNil(t, rect)
	assert.Nil(t, rect.Common)
}
