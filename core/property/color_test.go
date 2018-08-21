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
func TestColorTestSuite(t *testing.T) {
	suite.Run(t, new(ColorTestSuite))
}
func (suite *ColorTestSuite) TestSetRGB() {
	suite.color.SetRGB(uint8(225), uint8(225), uint8(225))
	suite.Equal(uint8(225), suite.color.R)
	suite.Equal(uint8(225), suite.color.G)
	suite.Equal(uint8(225), suite.color.B)

	suite.color.SetRGB(0, 0, uint8(225))
	suite.Equal(uint8(0), suite.color.R)
	suite.Equal(uint8(0), suite.color.G)
	suite.Equal(uint8(225), suite.color.B)

	suite.color.SetRGB(192, 192, 192)
	suite.Equal(uint8(192), suite.color.R)
	suite.Equal(uint8(192), suite.color.G)
	suite.Equal(uint8(192), suite.color.B)
}

func (suite *ColorTestSuite) TestRGB() {
	r, g, b := suite.color.RGB()
	suite.Equal(uint8(0), r)
	suite.Equal(uint8(0), g)
	suite.Equal(uint8(0), b)

	suite.color.R = 255
	suite.color.G = 255
	suite.color.B = 255
	r, g, b = suite.color.RGB()
	suite.Equal(uint8(255), r)
	suite.Equal(uint8(255), g)
	suite.Equal(uint8(255), b)

	suite.color.R = 192
	suite.color.G = 192
	suite.color.B = 192
	r, g, b = suite.color.RGB()
	suite.Equal(uint8(192), r)
	suite.Equal(uint8(192), g)
	suite.Equal(uint8(192), b)
}

func (suite *ColorTestSuite) TestHex() {
	hex := suite.color.Hex()
	suite.NotNil(hex)
	suite.True(len(hex) > 0)
	suite.Equal("#000000", hex)

	suite.color.R = 255
	suite.color.G = 255
	suite.color.B = 255
	hex = suite.color.Hex()
	hex = strings.ToUpper(hex)
	suite.Equal("#FFFFFF", hex)

	suite.color.R = 0
	suite.color.G = 0
	suite.color.B = 255
	hex = suite.color.Hex()
	hex = strings.ToUpper(hex)
	suite.Equal("#0000FF", hex)

	suite.color.R = 128
	suite.color.G = 128
	suite.color.B = 128
	hex = suite.color.Hex()
	hex = strings.ToUpper(hex)
	suite.Equal("#BCBCBC", hex)
}

func (suite *ColorTestSuite) TestFromHex() {
	suite.True(suite.color.FromHex("#000000"))
	suite.Equal(uint8(0), suite.color.R)
	suite.Equal(uint8(0), suite.color.G)
	suite.Equal(uint8(0), suite.color.B)

	suite.True(suite.color.FromHex("#FFFFFF"))
	suite.Equal(uint8(255), suite.color.R)
	suite.Equal(uint8(255), suite.color.G)
	suite.Equal(uint8(255), suite.color.B)

	suite.True(suite.color.FromHex("#0000FF"))
	suite.Equal(uint8(0), suite.color.R)
	suite.Equal(uint8(0), suite.color.G)
	suite.Equal(uint8(255), suite.color.B)

	suite.True(suite.color.FromHex("#C0C0C0"))
	suite.Equal(uint8(192), suite.color.R)
	suite.Equal(uint8(192), suite.color.G)
	suite.Equal(uint8(192), suite.color.B)

	suite.False(suite.color.FromHex("AAAA"))
}
