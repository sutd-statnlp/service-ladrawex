package property_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/property"
)

func TestCommon(t *testing.T) {
	common := new(property.Common)
	assert.NotNil(t, common)
	assert.Nil(t, common.Border)
	assert.Nil(t, common.BackgroundColor)
	assert.Nil(t, common.Size)
	assert.Nil(t, common.Position)
}
