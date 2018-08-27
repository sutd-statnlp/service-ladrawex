package component_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/component"
)

func TestDiamond(t *testing.T) {
	item := new(component.Diamond)
	assert.NotNil(t, item)
	assert.Nil(t, item.Common)
}
