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
	query := `\node [draw={rgb,255:red,?;green,?;blue,?}, line width=?mm, shape=rectangle, fill={rgb,255:red,?;green,?;blue,?}, minimum width=?mm, minimum height=?mm] at (?,?) {};`
	latex := stringutil.Prepare(query,
		rect.Common.Border.Color.R,
		rect.Common.Border.Color.G,
		rect.Common.Border.Color.B,
		rect.Common.Border.Thick,
		rect.Common.BackgroundColor.R,
		rect.Common.BackgroundColor.G,
		rect.Common.BackgroundColor.B,
		rect.Common.Size.Width,
		rect.Common.Size.Height,
		rect.Common.Position.XAxis,
		rect.Common.Position.YAxis,
	)
	return latex
}
