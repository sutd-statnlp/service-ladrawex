package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/core"
	"github.com/sutd-statnlp/service-ladrawex/core/component"
	"github.com/sutd-statnlp/service-ladrawex/core/stringutil"
)

var (
	drawexRest DrawexRest
)

func init() {
	drawexRest = NewDrawexRest()
}

// NewDrawexRest creates new drawex rest api.
func NewDrawexRest() DrawexRest {
	return &DrawexRestImpl{
		drawex: core.DefaultDrawex(),
	}
}

// DefaultDrawexRest returns the default drawex rest api.
func DefaultDrawexRest() DrawexRest {
	return drawexRest
}

// Group returns the api group.
func Group(router *gin.Engine) bool {
	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/draw", drawexRest.PostDraw)
	}
	return apiGroup != nil
}

// NewDrawexRestWithField creates new drawex rest with field.
func NewDrawexRestWithField(drawex core.Drawex) DrawexRest {
	return &DrawexRestImpl{
		drawex: drawex,
	}
}

// DocumentFromRequestBody returns latex document from request body.
func DocumentFromRequestBody(drawex core.Drawex, requestBody *RequestBody) *string {
	var latexs []*string
	latexs = append(latexs, LatexFromRectangles(drawex, requestBody.Rectangles))
	latexs = append(latexs, LatexFromCircles(drawex, requestBody.Circles))
	latexs = append(latexs, LatexFromLines(drawex, requestBody.Lines))
	latexs = append(latexs, LatexFromTexts(drawex, requestBody.Texts))
	latexs = append(latexs, LatexFromDiamons(drawex, requestBody.Diamonds))
	latexs = append(latexs, LatexFromPolygons(drawex, requestBody.Polygons))
	latexs = append(latexs, LatexFromConnectors(drawex, requestBody.Connectors))
	latexs = append(latexs, LatexFromTriangles(drawex, requestBody.Triangles))
	latexs = append(latexs, LatexFromStars(drawex, requestBody.Stars))
	return drawex.Document(latexs)
}

// LatexFromRectangles returns latex from rectangles.
func LatexFromRectangles(drawex core.Drawex, rectangles []*component.Rectangle) *string {
	var latex string
	for _, item := range rectangles {
		latex = stringutil.Concat(latex, drawex.DrawRectangle(item))
	}
	return &latex
}

// LatexFromCircles returns latex from circles.
func LatexFromCircles(drawex core.Drawex, circles []*component.Circle) *string {
	var latex string
	for _, item := range circles {
		latex = stringutil.Concat(latex, drawex.DrawCircle(item))
	}
	return &latex
}

// LatexFromLines returns latex from lines.
func LatexFromLines(drawex core.Drawex, lines []*component.Line) *string {
	var latex string
	for _, item := range lines {
		latex = stringutil.Concat(latex, drawex.DrawLine(item))
	}
	return &latex
}

// LatexFromTexts returns latex from texts.
func LatexFromTexts(drawex core.Drawex, texts []*component.Text) *string {
	var latex string
	for _, item := range texts {
		latex = stringutil.Concat(latex, drawex.DrawText(item))
	}
	return &latex
}

// LatexFromDiamons returns latex from diamonds.
func LatexFromDiamons(drawex core.Drawex, diamonds []*component.Diamond) *string {
	var latex string
	for _, item := range diamonds {
		latex = stringutil.Concat(latex, drawex.DrawDiamond(item))
	}
	return &latex
}

// LatexFromPolygons returns latex from polygons.
func LatexFromPolygons(drawex core.Drawex, polygons []*component.Polygon) *string {
	var latex string
	for _, item := range polygons {
		latex = stringutil.Concat(latex, drawex.DrawPolygon(item))
	}
	return &latex
}

// LatexFromConnectors returns latex from connectors.
func LatexFromConnectors(drawex core.Drawex, connectors []*component.Connector) *string {
	var latex string
	for _, item := range connectors {
		latex = stringutil.Concat(latex, drawex.DrawConnector(item))
	}
	return &latex
}

// LatexFromTriangles returns latex from triangles.
func LatexFromTriangles(drawex core.Drawex, triangles []*component.Triangle) *string {
	var latex string
	for _, item := range triangles {
		latex = stringutil.Concat(latex, drawex.DrawTriangle(item))
	}
	return &latex
}

// LatexFromStars returns latex from stars.
func LatexFromStars(drawex core.Drawex, stars []*component.Star) *string {
	var latex string
	for _, item := range stars {
		latex = stringutil.Concat(latex, drawex.DrawStar(item))
	}
	return &latex
}
