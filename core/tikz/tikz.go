package tikz

import (
	"github.com/sutd-statnlp/service-ladrawex/core/property"
	"github.com/sutd-statnlp/service-ladrawex/core/stringutil"
)

var (
	drawex *DrawexImpl
)

func init() {
	drawex = New()
}

// Default returns default tikz drawex.
func Default() *DrawexImpl {
	return drawex
}

// New creates new drawex.
func New() *DrawexImpl {
	return &DrawexImpl{}
}

// ColorToQuery converts color property into query.
func ColorToQuery(query string, color *property.Color) string {
	var result string
	if color != nil {
		result = stringutil.Prepare(query,
			color.R,
			color.G,
			color.B,
		)
	}
	return result
}
