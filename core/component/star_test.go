package component_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/component"
)

func TestStar(t *testing.T) {
	item := new(component.Star)
	assert.NotNil(t, item)
	assert.Nil(t, item.Common)
}
