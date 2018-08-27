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
	DefaultLineLatex = `\draw [, line width=0mm] (0,0) -- (0,0);`
)

// CreateDefaultLatex creates the default latex for testing from diffrent shapes.
func CreateDefaultLatex(shapName string, text string) string {
	return stringutil.Prepare(`\node [shape=?, , line width=0mm, , minimum width=0mm, minimum height=0mm] at (0,0) {?};`, shapName, text)
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
	assert.Equal(t, "", query)

	color.Enable = true
	query = tikz.ColorToQuery("???", color)
	assert.Equal(t, "000", query)
}
