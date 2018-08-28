package tikz

import (
	"bytes"

	"github.com/sutd-statnlp/service-ladrawex/core/component"
	"github.com/sutd-statnlp/service-ladrawex/core/property"
	"github.com/sutd-statnlp/service-ladrawex/core/stringutil"
	"github.com/sutd-statnlp/service-ladrawex/log"
)

const (
	// DocumentStart is the starting area of latex document.
	DocumentStart = `\documentclass{article}\usepackage{tikz}\usetikzlibrary{shapes,snakes}\begin{document}\begin{tikzpicture}`

	// DocumentEnd is the end area of latex document.
	DocumentEnd = `\end{tikzpicture}\end{document}`
)

const (
	// DrawOption is the draw style option.
	DrawOption = "draw="

	// FillOption is the fill style option.
	FillOption = "fill="
)

const (
	// RGBQuery is the query for drawing rgb color.
	RGBQuery = `{rgb,255:red,?;green,?;blue,?}`

	// NodeDrawQuery is the query for drawing tikz node.
	NodeDrawQuery = `\node [shape=?, ?, line width=?mm, ?, minimum width=?mm, minimum height=?mm] at (?,?) {?};`

	// LineDrawQuery is the query for drawing line component.
	LineDrawQuery = `\draw [?, line width=?mm] (?,?) -- (?,?);`

	// ConnectorDrawQuery is the query for drawing connector component.
	ConnectorDrawQuery = `\draw [->, >=stealth, ?, line width=?mm] (?,?) -- (?,?);`
)

const (
	// RectangleShape is the name of rectangle shape.
	RectangleShape = "rectangle"

	// CircleShape is the name of circle shape.
	CircleShape = "circle"

	// DiamondShape is the name of diamond shape.
	DiamondShape = "diamond"
)

var (
	// DrawRGBQuery is the query for drawing rgb.
	DrawRGBQuery string

	// FillRGBQuery is the query for filling rgb.
	FillRGBQuery string
)

func init() {
	DrawRGBQuery = stringutil.Concat(DrawOption, RGBQuery)
	FillRGBQuery = stringutil.Concat(FillOption, RGBQuery)
}

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

	borderColor := ColorToQuery(DrawRGBQuery, common.Border.Color)
	backgroundColor := ColorToQuery(FillRGBQuery, common.BackgroundColor)
	latex := stringutil.Prepare(NodeDrawQuery,
		shapeName,
		borderColor,
		common.Border.Thick,
		backgroundColor,
		common.Size.Width,
		common.Size.Height,
		common.Position.X,
		common.Position.Y,
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
	color := ColorToQuery(FillRGBQuery, line.Color)
	latex := stringutil.Prepare(LineDrawQuery,
		color,
		line.Width,
		line.StartPosition.X,
		line.StartPosition.Y,
		line.EndPosition.X,
		line.EndPosition.Y,
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

// DrawDiamond draws a diamond and returns latex.
func (drawex *DrawexImpl) DrawDiamond(diamon *component.Diamond) string {
	log.Debug("TikzDrawexImpl request to draw diamond: ", stringutil.JSON(diamon))
	if diamon == nil {
		log.Error("TikzDrawexImpl request to draw null diamon")
		return ""
	}
	return drawex.DrawNode(DiamondShape, diamon.Common)
}

// DrawConnector draws a connection and returns latex.
func (drawex *DrawexImpl) DrawConnector(connector *component.Connector) string {
	log.Debug("TikzDrawexImpl request to draw a connector: ", stringutil.JSON(connector))
	if connector == nil {
		log.Error("TikzDrawexImpl request to draw null connector")
		return ""
	}
	color := ColorToQuery(FillRGBQuery, connector.Line.Color)
	latex := stringutil.Prepare(ConnectorDrawQuery,
		color,
		connector.Line.Width,
		connector.Line.StartPosition.X,
		connector.Line.StartPosition.Y,
		connector.Line.EndPosition.X,
		connector.Line.EndPosition.Y,
	)
	return latex
}
