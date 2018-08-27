package component_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/component"
)

func TestConnector(t *testing.T) {
	item := new(component.Connector)
	assert.NotNil(t, item)
	assert.Nil(t, item.Line)
	assert.False(t, item.Bidirectional)
}
