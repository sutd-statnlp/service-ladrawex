package tikz_test

import (
	"testing"

	"github.com/sutd-statnlp/service-ladrawex/core/component"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core/tikz"
)

func CreateFakeRectangle() *component.Rectangle {
	return component.NewRectangle()
}

func TestDefault(t *testing.T) {
	drawer := tikz.Default()
	assert.NotNil(t, drawer)
}

func TestNew(t *testing.T) {
	drawer := tikz.New()
	assert.NotNil(t, drawer)
}
