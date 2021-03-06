package tikz_test

import (
	"testing"

	"github.com/sutd-statnlp/service-ladrawex/core/component"
	"github.com/sutd-statnlp/service-ladrawex/core/stringutil"

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
	diamondLatex   string
	fakeDiamond    *component.Diamond
	connectorLatex string
	fakeConnector  *component.Connector
	polygonLatex   string
	fakePolygon    *component.Polygon
	triangleLatex  string
	fakeTriangle   *component.Triangle
	starLatex      string
	fakeStar       *component.Star
}

func (suite *DrawexImplTestSuite) SetupTest() {
	suite.drawex = tikz.New()

	suite.rectangleLatex = CreateDefaultLatex("rectangle", "", 0, 0, 0)
	suite.fakeRectangle = CreateFakeRectangle()

	suite.circleLatex = CreateDefaultLatex("circle", "", 0, 0, 0)
	suite.fakeCircle = CreateFakeCircle()

	suite.lineLatex = DefaultLineLatex
	suite.fakeLine = CreateFakeLine()

	suite.textLatex = CreateDefaultLatex("rectangle", "text", 255, 255, 255)
	suite.fakeText = CreateFakeText()

	suite.diamondLatex = CreateDefaultLatex("diamond", "", 0, 0, 0)
	suite.fakeDiamond = CreateFakeDiamond()

	suite.connectorLatex = DefaultConnectorLatex
	suite.fakeConnector = CreateFakeConnector()

	suite.polygonLatex = CreateDefaultLatex("regular polygon,regular polygon sides=5", "", 0, 0, 0)
	suite.fakePolygon = CreateFakePolygon()

	suite.triangleLatex = DefaultTriangleLatex
	suite.fakeTriangle = CreateFakeTriangle()

	suite.starLatex = CreateDefaultLatex("star,star points=0", "", 0, 0, 0)
	suite.fakeStar = CreateFakeStar()
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
	suite.Equal(stringutil.Concat(tikz.DocumentStart, "AAAA", tikz.DocumentEnd),
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

func (suite *DrawexImplTestSuite) TestDrawDiamond() {
	latex := suite.drawex.DrawDiamond(suite.fakeDiamond)
	suite.NotNil(latex)
	suite.True(len(latex) > 0)
	suite.Equal(suite.diamondLatex, latex)
	suite.Empty(suite.drawex.DrawText(nil))
}

func (suite *DrawexImplTestSuite) TestDrawConnector() {
	latex := suite.drawex.DrawConnector(suite.fakeConnector)
	suite.NotNil(latex)
	suite.True(len(latex) > 0)
	suite.Equal(suite.connectorLatex, latex)
	suite.Empty(suite.drawex.DrawConnector(nil))

	suite.fakeConnector.Bidirectional = true
	latex = suite.drawex.DrawConnector(suite.fakeConnector)
	suite.NotEqual(suite.connectorLatex, latex)
}

func (suite *DrawexImplTestSuite) TestDrawPolygon() {
	latex := suite.drawex.DrawPolygon(suite.fakePolygon)
	suite.NotNil(latex)
	suite.True(len(latex) > 0)
	suite.Equal(suite.polygonLatex, latex)
	suite.Empty(suite.drawex.DrawPolygon(nil))
}

func (suite *DrawexImplTestSuite) TestDrawTriangle() {
	latex := suite.drawex.DrawTriangle(suite.fakeTriangle)
	suite.NotNil(latex)
	suite.True(len(latex) > 0)
	suite.Equal(suite.triangleLatex, latex)
	suite.Empty(suite.drawex.DrawTriangle(nil))
}

func (suite *DrawexImplTestSuite) TestDrawStar() {
	latex := suite.drawex.DrawStar(suite.fakeStar)
	suite.True(len(latex) > 0)
	suite.Equal(suite.starLatex, latex)
	suite.Empty(suite.drawex.DrawStar(nil))
}
