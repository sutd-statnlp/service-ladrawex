package apiutil_test

import (
	"github.com/sutd-statnlp/service-ladrawex/web/apiutil"
	"testing"
	"github.com/stretchr/testify/assert"
)

// CreateFakeRequestCommonProperty creates a fake requesting common property for testing.
func CreateFakeRequestCommonProperty() *apiutil.Common {
	return &apiutil.Common{
		Border:   apiutil.Border{},
		Size:     apiutil.Size{},
		Position: apiutil.Position{},
		Text:     apiutil.TextProperty{},
	}
}

// CreateFakeRequestRectangle creates a fake requesting rectangle for testing.
func CreateFakeRequestRectangle() *apiutil.Rectangle {
	return &apiutil.Rectangle{
		Common: *CreateFakeRequestCommonProperty(),
	}
}

// CreateFakeRequestCircle creates a fake requesting circle for testing.
func CreateFakeRequestCircle() *apiutil.Circle {
	return &apiutil.Circle{
		Common: *CreateFakeRequestCommonProperty(),
	}
}

// CreateFakeRequestLine creates a fake requesting line for testing.
func CreateFakeRequestLine() *apiutil.Line {
	return &apiutil.Line{
		StartPosition: apiutil.Position{},
		EndPosition:   apiutil.Position{},
	}
}

// CreateFakeRequestText creates a fake requesting text for testing.
func CreateFakeRequestText() *apiutil.Text {
	return &apiutil.Text{
		Common: *CreateFakeRequestCommonProperty(),
	}
}

func TestRectangleToComponent(t *testing.T) {
	rect := apiutil.RectangleToComponent(CreateFakeRequestRectangle())
	assert.NotNil(t, rect)
	assert.NotNil(t, rect.Common)
}

func TestCommonToProperty(t *testing.T) {
	common := apiutil.CommonToProperty(CreateFakeRequestCommonProperty())
	assert.NotNil(t, common)
	assert.NotNil(t, common.Border)
	assert.NotNil(t, common.BackgroundColor)
	assert.NotNil(t, common.Size)
	assert.NotNil(t, common.Position)
	assert.NotNil(t, common.Text)
}

func TestCircleToComponent(t *testing.T) {
	circle := apiutil.CircleToComponent(CreateFakeRequestCircle())
	assert.NotNil(t, circle)
	assert.NotNil(t, circle.Common)
}

func TestLineToComponent(t *testing.T) {
	line := apiutil.LineToComponent(CreateFakeRequestLine())
	assert.NotNil(t, line)
	assert.NotNil(t, line.StartPosition)
	assert.NotNil(t, line.EndPosition)
}

func TestTextToComponent(t *testing.T) {
	text := apiutil.TextToComponent(CreateFakeRequestText())
	assert.NotNil(t, text)
	assert.NotNil(t, text.Common)
}
