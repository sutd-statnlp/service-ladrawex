package component_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/component"
)

func TestTriangle(t *testing.T) {
	item := new(component.Triangle)
	assert.NotNil(t, item)
	assert.Nil(t, item.Common)
	assert.Nil(t, item.SecondPosition)
	assert.Nil(t, item.ThirdPosition)
}
