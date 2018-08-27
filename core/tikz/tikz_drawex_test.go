package tikz_test

import (
	"testing"

	"github.com/sutd-statnlp/service-ladrawex/core/component"

	"github.com/stretchr/testify/suite"
	"github.com/sutd-statnlp/service-ladrawex/core/tikz"
)

type DrawexImplTestSuite struct {
	suite.Suite
	drawex         *tikz.DrawexImpl
	rectangleLatex string
	fakeRectangle  *component.Rectangle
	circleLatex    string
	fakeCircle     *component.Circle
	lineLatex      string
	fakeLine       *component.Line
	textLatex      string
	fakeText       *component.Text
}

func (suite *DrawexImplTestSuite) SetupTest() {
	suite.drawex = tikz.New()

	suite.rectangleLatex = CreateDefaultLatex("rectangle", "")
	suite.fakeRectangle = CreateFakeRectangle()

	suite.circleLatex = CreateDefaultLatex("circle", "")
	suite.fakeCircle = CreateFakeCircle()

	suite.lineLatex = DefaultLineLatex
	suite.fakeLine = CreateFakeLine()

	suite.textLatex = CreateDefaultLatex("rectangle", "text")
	suite.fakeText = CreateFakeText()
}

func TestDrawexImplTestSuite(t *testing.T) {
	suite.Run(t, new(DrawexImplTestSuite))
}

func (suite *DrawexImplTestSuite) TestDrawRectangle() {
	latex := suite.drawex.DrawRectangle(suite.fakeRectangle)
	suite.NotNil(latex)
	suite.NotEmpty(latex)
	suite.True(len(latex) > 0)
	suite.Equal(suite.rectangleLatex, latex)
	suite.Empty(suite.drawex.DrawRectangle(nil))
}

func (suite *DrawexImplTestSuite) TestDrawNode() {
	latex := suite.drawex.DrawNode(tikz.RectangleShape, suite.fakeRectangle.Common)
	suite.NotEmpty(latex)
	suite.True(len(latex) > 0)
	suite.Equal(suite.rectangleLatex, latex)
	suite.Empty(suite.drawex.DrawNode(tikz.RectangleShape, nil))
}

func (suite *DrawexImplTestSuite) TestDrawCircle() {
	latex := suite.drawex.DrawCircle(suite.fakeCircle)
	suite.NotEmpty(latex)
	suite.True(len(latex) > 0)
	suite.Equal(suite.circleLatex, latex)
	suite.Empty(suite.drawex.DrawCircle(nil))
}

func (suite *DrawexImplTestSuite) TestDocument() {
	tex := "AA"
	latexs := []*string{&tex, &tex}
	doc := *suite.drawex.Document(latexs)
	suite.NotNil(doc)
	suite.NotNil(len(doc) > 0)
	suite.Equal(`\documentclass{article}\usepackage{tikz}\begin{document}\begin{tikzpicture}AAAA\end{tikzpicture}\end{document}`,
		doc,
	)
}

func (suite *DrawexImplTestSuite) TestDrawLine() {
	latex := suite.drawex.DrawLine(suite.fakeLine)
	suite.NotNil(latex)
	suite.True(len(latex) > 0)
	suite.Equal(suite.lineLatex, latex)
	suite.Empty(suite.drawex.DrawLine(nil))
}

func (suite *DrawexImplTestSuite) TestDrawText() {
	latex := suite.drawex.DrawText(suite.fakeText)
	suite.NotNil(latex)
	suite.True(len(latex) > 0)
	suite.Equal(suite.textLatex, latex)
	suite.Empty(suite.drawex.DrawText(nil))
}
