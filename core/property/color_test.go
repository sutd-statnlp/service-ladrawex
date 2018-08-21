package property_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/stretchr/testify/assert"

	"github.com/sutd-statnlp/service-ladrawex/core/property"
)

func TestNewColor(t *testing.T) {
	color := new(property.Color)
	assert.NotNil(t, color)
	assert.NotNil(t, color.R)
	assert.NotNil(t, color.G)
	assert.NotNil(t, color.B)
	assert.Equal(t, uint8(0), color.R)
	assert.Equal(t, uint8(0), color.G)
	assert.Equal(t, uint8(0), color.B)
}

type ColorTestSuite struct {
	suite.Suite
	color *property.Color
}

func (suite *ColorTestSuite) SetupTest() {
	suite.color = new(property.Color)
}

func (suite *ColorTestSuite) TestHex() {
	hex := suite.color.Hex()
	suite.NotNil(hex)
	suite.True(len(hex) > 0)
	suite.True(strings.Contains(hex, "#"))
}

func TestColorTestSuite(t *testing.T) {
	suite.Run(t, new(ColorTestSuite))
}
