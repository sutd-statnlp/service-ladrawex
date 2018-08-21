package tikz

import (
	"github.com/sutd-statnlp/service-ladrawex/core/component"
	"github.com/sutd-statnlp/service-ladrawex/core/stringutil"
	"github.com/sutd-statnlp/service-ladrawex/log"
)

// DrawexImpl is the implementation of drawex interface.
type DrawexImpl struct {
}

// DrawRectangle draws rectangle and returns latex.
func (drawex *DrawexImpl) DrawRectangle(rect *component.Rectangle) string {
	log.Debug("Tikz DrawexImpl request to draw rectangle: ", stringutil.JSON(rect))
	if rect == nil {
		log.Error("Tikz DrawexImpl request to draw null rectangle")
		return ""
	}
	latex := stringutil.Prepare(
		`\draw (?,?) rectangle (?,?);`,
		rect.Position.XAxis,
		rect.Position.YAxis,
		rect.Size.Width,
		rect.Size.Height,
	)
	return latex
}
