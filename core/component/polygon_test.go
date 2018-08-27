package component_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/component"
)

func TestPolygon(t *testing.T) {
	item := new(component.Polygon)
	assert.NotNil(t, item)
	assert.Nil(t, item.Common)
	assert.Equal(t, uint8(0), item.Sides)
}
