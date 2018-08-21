package property_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/property"
)

func TestCreateNewBorder(t *testing.T) {
	border := property.NewBorder()
	assert.NotNil(t, border)
	assert.NotNil(t, border.Color)
}
