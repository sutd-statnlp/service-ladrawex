package tikz_test

import (
	"testing"

	"github.com/sutd-statnlp/service-ladrawex/core/component"
	"github.com/sutd-statnlp/service-ladrawex/core/stringutil"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/property"
	"github.com/sutd-statnlp/service-ladrawex/core/tikz"
)

const (
	// DefaultLineLatex is the line latex for testing.
	DefaultLineLatex = `\draw [fill={rgb,255:red,0;green,0;blue,0}, line width=0mm] (0,0) -- (0,0);`
)

// CreateDefaultLatex creates the default latex for testing from diffrent shapes.
func CreateDefaultLatex(shapName string, text string, r, g, b uint8) string {
	return stringutil.Prepare(`\node [shape=?, draw={rgb,255:red,0;green,0;blue,0}, line width=0mm, fill={rgb,255:red,?;green,?;blue,?}, minimum width=0mm, minimum height=0mm] at (0,0) {?};`, shapName, r, g, b, text)
}

// CreateFakeRectangle creates fake rectangle for testing.
func CreateFakeRectangle() *component.Rectangle {
	return component.NewRectangle()
}

// CreateFakeCircle creates fake circle for testing.
func CreateFakeCircle() *component.Circle {
	return component.NewCircle()
}

// CreateFakeLine creates fake line for testing.
func CreateFakeLine() *component.Line {
	return component.NewLine()
}

// CreateFakeText creates fake text for testing.
func CreateFakeText() *component.Text {
	return component.NewText()
}

// CreateFakeDiamond creates fake diamond for testing.
func CreateFakeDiamond() *component.Diamond {
	return component.NewDiamond()
}
func TestDefault(t *testing.T) {
	drawer := tikz.Default()
	assert.NotNil(t, drawer)
}

func TestNew(t *testing.T) {
	drawer := tikz.New()
	assert.NotNil(t, drawer)
}

func TestColorToQuery(t *testing.T) {
	color := new(property.Color)
	query := tikz.ColorToQuery("???", color)
	assert.Equal(t, "000", query)

	color = nil
	query = tikz.ColorToQuery("???", color)
	assert.Equal(t, "", query)
}
