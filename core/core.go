package core

import "github.com/sutd-statnlp/service-ladrawex/core/tikz"

var (
	drawex Drawex
)

func init() {
	drawex = NewDrawex()
}

// DefaultDrawex returns default latex drawer.
func DefaultDrawex() Drawex {
	return drawex
}

// NewDrawex creates new latex drawer.
func NewDrawex() Drawex {
	return tikz.New()
}

// SetDrawex sets latex drawer by given drawex.
func SetDrawex(inputDrawex Drawex) bool {
	if inputDrawex == nil {
		return false
	}
	drawex = inputDrawex
	return true
}
