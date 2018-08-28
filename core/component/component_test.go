package component_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/component"
)

func TestCreateNewRectangle(t *testing.T) {
	rect := component.NewRectangle()
	assert.NotNil(t, rect)
	assert.NotNil(t, rect.Common)
}

func TestNewCommonProperty(t *testing.T) {
	common := component.NewCommonProperty()
	assert.NotNil(t, common)
	assert.NotNil(t, common.Border)
	assert.NotNil(t, common.BackgroundColor)
	assert.NotNil(t, common.Size)
	assert.NotNil(t, common.Position)
}

func TestNewCircle(t *testing.T) {
	circle := component.NewCircle()
	assert.NotNil(t, circle)
	assert.NotNil(t, circle.Common)
}

func TestNewLine(t *testing.T) {
	line := component.NewLine()
	assert.NotNil(t, line)
	assert.NotNil(t, line.Color)
	assert.NotNil(t, line.StartPosition)
	assert.NotNil(t, line.EndPosition)
}

func TestNewText(t *testing.T) {
	text := component.NewText()
	assert.NotNil(t, text)
	assert.NotNil(t, text.Common)
	assert.Equal(t, uint8(255), text.Common.BackgroundColor.R)
	assert.Equal(t, uint8(255), text.Common.BackgroundColor.G)
	assert.Equal(t, uint8(255), text.Common.BackgroundColor.B)
}

func TestNewTriangle(t *testing.T) {
	item := component.NewTriangle()
	assert.NotNil(t, item)
	assert.NotNil(t, item.Common)
	assert.NotNil(t, item.SecondPosition)
	assert.Equal(t, float32(4), item.SecondPosition.X)
	assert.NotNil(t, item.ThirdPosition)
	assert.Equal(t, float32(4), item.ThirdPosition.X)
	assert.Equal(t, float32(4), item.ThirdPosition.Y)
}

func TestNewPolygon(t *testing.T) {
	item := component.NewPolygon()
	assert.NotNil(t, item)
	assert.NotNil(t, item.Common)
	assert.Equal(t, uint8(5), item.Sides)
}

func TestNewDiamond(t *testing.T) {
	item := component.NewDiamond()
	assert.NotNil(t, item)
	assert.NotNil(t, item.Common)
}

func TestNewConnector(t *testing.T) {
	item := component.NewConnector()
	assert.NotNil(t, item)
	assert.NotNil(t, item.Line)
}

func TestNewStar(t *testing.T) {
	item := component.NewStar()
	assert.NotNil(t, item)
	assert.NotNil(t, item.Common)
}
