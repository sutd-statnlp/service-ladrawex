package tikz

import (
	"github.com/sutd-statnlp/service-ladrawex/core/component"
	"github.com/sutd-statnlp/service-ladrawex/core/property"
	"github.com/sutd-statnlp/service-ladrawex/core/stringutil"
	"github.com/sutd-statnlp/service-ladrawex/log"
)

const (
	// NodeDrawQuery is the query for drawing tikz node.
	NodeDrawQuery = `\node [shape=?, draw={rgb,255:red,?;green,?;blue,?}, line width=?mm, fill={rgb,255:red,?;green,?;blue,?}, minimum width=?mm, minimum height=?mm] at (?,?) {};`
)

const (
	// RectangleShape is the name of rectangle shape.
	RectangleShape = "rectangle"

	// CircleShape is the name of circle shape.
	CircleShape = "circle"
)

// DrawexImpl is the implementation of drawex interface.
type DrawexImpl struct {
}

// DrawRectangle draws rectangle and returns latex.
func (drawex *DrawexImpl) DrawRectangle(rect *component.Rectangle) string {
	log.Debug("TikzDrawexImpl request to draw rectangle: ", stringutil.JSON(rect))
	if rect == nil {
		log.Error("TikzDrawexImpl request to draw null rectangle")
		return ""
	}
	return drawex.DrawNode(RectangleShape, rect.Common)
}

// DrawNode draws common node from shape name and common property.
func (drawex *DrawexImpl) DrawNode(shapeName string, common *property.Common) string {
	log.Debug("TikzDrawexImpl request to draw a node with shape name: ", shapeName, " and common: ", stringutil.JSON(common))
	if common == nil {
		return ""
	}
	latex := stringutil.Prepare(NodeDrawQuery,
		shapeName,
		common.Border.Color.R,
		common.Border.Color.G,
		common.Border.Color.B,
		common.Border.Thick,
		common.BackgroundColor.R,
		common.BackgroundColor.G,
		common.BackgroundColor.B,
		common.Size.Width,
		common.Size.Height,
		common.Position.XAxis,
		common.Position.YAxis,
	)
	return latex
}

// DrawCircle draws circle and returns latex.
func (drawex *DrawexImpl) DrawCircle(circle *component.Circle) string {
	log.Debug("TikzDrawexImpl request to draw circle: ", stringutil.JSON(circle))
	if circle == nil {
		log.Error("TikzDrawexImpl request to draw null circle")
		return ""
	}
	return drawex.DrawNode(CircleShape, circle.Common)
}
