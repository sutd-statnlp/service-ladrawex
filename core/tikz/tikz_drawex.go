package tikz

import (
	"bytes"

	"github.com/sutd-statnlp/service-ladrawex/core/component"
	"github.com/sutd-statnlp/service-ladrawex/core/property"
	"github.com/sutd-statnlp/service-ladrawex/core/stringutil"
	"github.com/sutd-statnlp/service-ladrawex/log"
)

const (
	// NodeDrawQuery is the query for drawing tikz node.
	NodeDrawQuery = `\node [shape=?, draw={rgb,255:red,?;green,?;blue,?}, line width=?mm, fill={rgb,255:red,?;green,?;blue,?}, minimum width=?mm, minimum height=?mm] at (?,?) {?};`

	// LineDrawQuery is the query for drawing line component.
	LineDrawQuery = `\draw [draw={rgb,255:red,?;green,?;blue,?}, line width=?mm] (?,?) -- (?,?);`
)

const (
	// DocumentStart is the starting area of latex document.
	DocumentStart = `\documentclass{article}\usepackage{tikz}\begin{document}\begin{tikzpicture}`

	// DocumentEnd is the end area of latex document.
	DocumentEnd = `\end{tikzpicture}\end{document}`
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
		common.Text.Content,
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

// Document returns latex document.
func (drawex *DrawexImpl) Document(latexs []*string) *string {
	log.Debug("TikzDrawexImpl request to get latex document")
	var buffer bytes.Buffer
	buffer.WriteString(DocumentStart)
	for _, item := range latexs {
		buffer.WriteString(*item)
	}
	buffer.WriteString(DocumentEnd)
	doc := buffer.String()
	return &doc
}

// DrawLine draws line and returns latex.
func (drawex *DrawexImpl) DrawLine(line *component.Line) string {
	log.Debug("TikzDrawexImpl request to draw a line: ", stringutil.JSON(line))
	if line == nil {
		log.Error("TikzDrawexImpl request to draw null line")
		return ""
	}
	latex := stringutil.Prepare(LineDrawQuery,
		line.Color.R,
		line.Color.G,
		line.Color.B,
		line.Width,
		line.StartPosition.XAxis,
		line.StartPosition.YAxis,
		line.EndPosition.XAxis,
		line.EndPosition.YAxis,
	)
	return latex
}

// DrawText draws text and returns latex.
func (drawex *DrawexImpl) DrawText(text *component.Text) string {
	log.Debug("TikzDrawexImpl request to draw text: ", stringutil.JSON(text))
	if text == nil {
		log.Error("TikzDrawexImpl request to draw null text")
		return ""
	}
	return drawex.DrawNode(RectangleShape, text.Common)
}
