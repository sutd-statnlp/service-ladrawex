package component_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sutd-statnlp/service-ladrawex/core/component"
)

func TestCircle(t *testing.T) {
	circle := new(component.Circle)
	assert.NotNil(t, circle)
	assert.Nil(t, circle.Common)
}
