package property_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/property"
)

func TestText(t *testing.T) {
	text := new(property.Text)
	assert.NotNil(t, text)
	assert.NotNil(t, text.Content)
}
