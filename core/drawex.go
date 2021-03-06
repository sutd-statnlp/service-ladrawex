package core

import "github.com/sutd-statnlp/service-ladrawex/core/component"

// Drawex is the latex drawer, which provides drawing features and generating latex.
type Drawex interface {
	DrawRectangle(rect *component.Rectangle) string
	DrawCircle(rect *component.Circle) string
	Document(latexs []*string) *string
	DrawLine(line *component.Line) string
	DrawText(text *component.Text) string
	DrawDiamond(diamond *component.Diamond) string
	DrawPolygon(polygon *component.Polygon) string
	DrawConnector(connector *component.Connector) string
	DrawTriangle(triangle *component.Triangle) string
	DrawStar(star *component.Star) string
}
