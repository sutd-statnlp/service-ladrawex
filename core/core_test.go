package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core"
	"github.com/sutd-statnlp/service-ladrawex/core/tikz"
)

func TestDefaultDrawex(t *testing.T) {
	drawex := core.DefaultDrawex()
	assert.NotNil(t, drawex)
}

func TestNewDrawex(t *testing.T) {
	drawex := core.NewDrawex()
	assert.NotNil(t, drawex)
}

func TestSetDrawex(t *testing.T) {
	assert.True(t, core.SetDrawex(tikz.New()))
	assert.False(t, core.SetDrawex(nil))
}
