package api

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/core"
	"github.com/sutd-statnlp/service-ladrawex/core/component"
	"github.com/sutd-statnlp/service-ladrawex/core/stringutil"
)

const (
	// LatexChanCapacity is the capacity of latex buffered channel.
	LatexChanCapacity = 9
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
	var wg sync.WaitGroup
	wg.Add(LatexChanCapacity)
	latexChan := make(chan *string, LatexChanCapacity)
	go func() {
		defer wg.Done()
		latexChan <- LatexFromRectangles(drawex, requestBody.Rectangles)
	}()
	go func() {
		defer wg.Done()
		latexChan <- LatexFromCircles(drawex, requestBody.Circles)
	}()
	go func() {
		defer wg.Done()
		latexChan <- LatexFromLines(drawex, requestBody.Lines)
	}()
	go func() {
		defer wg.Done()
		latexChan <- LatexFromTexts(drawex, requestBody.Texts)
	}()
	go func() {
		defer wg.Done()
		latexChan <- LatexFromDiamons(drawex, requestBody.Diamonds)
	}()
	go func() {
		defer wg.Done()
		latexChan <- LatexFromPolygons(drawex, requestBody.Polygons)
	}()
	go func() {
		defer wg.Done()
		latexChan <- LatexFromConnectors(drawex, requestBody.Connectors)
	}()
	go func() {
		defer wg.Done()
		latexChan <- LatexFromTriangles(drawex, requestBody.Triangles)
	}()
	go func() {
		defer wg.Done()
		latexChan <- LatexFromStars(drawex, requestBody.Stars)
	}()
	go func() {
		wg.Wait()
		close(latexChan)
	}()
	var latexs []*string
	for item := range latexChan {
		latexs = append(latexs, item)
	}
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
