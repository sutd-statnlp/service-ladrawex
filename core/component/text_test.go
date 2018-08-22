package component_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/component"
)

func TestText(t *testing.T) {
	text := new(component.Text)
	assert.NotNil(t, text)
	assert.Nil(t, text.Common)
}
