package tikz_test

import (
	"testing"

	"github.com/sutd-statnlp/service-ladrawex/core/component"
	"github.com/sutd-statnlp/service-ladrawex/core/stringutil"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/tikz"
)

// CreateDefaultPreparedQuery creates the default prepared query for testing from diffrent shapes.
func CreateDefaultPreparedQuery(shapName string) string {
	return stringutil.Prepare(`\node [shape=?, draw={rgb,255:red,0;green,0;blue,0}, line width=0mm, fill={rgb,255:red,0;green,0;blue,0}, minimum width=0mm, minimum height=0mm] at (0,0) {};`, shapName)
}

// CreateFakeRectangle creates fake rectangle for testing.
func CreateFakeRectangle() *component.Rectangle {
	return component.NewRectangle()
}

// CreateFakeCircle creates fake circle for testing.
func CreateFakeCircle() *component.Circle {
	return component.NewCircle()
}

func TestDefault(t *testing.T) {
	drawer := tikz.Default()
	assert.NotNil(t, drawer)
}

func TestNew(t *testing.T) {
	drawer := tikz.New()
	assert.NotNil(t, drawer)
}
