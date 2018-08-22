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
	rectangleQuery string
	fakeRectangle  *component.Rectangle
	circleQuery    string
	fakeCircle     *component.Circle
}

func (suite *DrawexImplTestSuite) SetupTest() {
	suite.drawex = tikz.New()
	suite.rectangleQuery = CreateDefaultPreparedQuery("rectangle")
	suite.fakeRectangle = CreateFakeRectangle()
	suite.circleQuery = CreateDefaultPreparedQuery("circle")
	suite.fakeCircle = CreateFakeCircle()
}

func TestDrawexImplTestSuite(t *testing.T) {
	suite.Run(t, new(DrawexImplTestSuite))
}

func (suite *DrawexImplTestSuite) TestDrawRectangle() {
	latex := suite.drawex.DrawRectangle(suite.fakeRectangle)
	suite.NotNil(latex)
	suite.NotEmpty(latex)
	suite.True(len(latex) > 0)
	suite.Equal(latex, suite.rectangleQuery)
	suite.Empty(suite.drawex.DrawRectangle(nil))
}

func (suite *DrawexImplTestSuite) TestDrawNode() {
	latex := suite.drawex.DrawNode(tikz.RectangleShape, suite.fakeRectangle.Common)
	suite.NotEmpty(latex)
	suite.True(len(latex) > 0)
	suite.Equal(latex, suite.rectangleQuery)
	suite.Empty(suite.drawex.DrawNode(tikz.RectangleShape, nil))
}

func (suite *DrawexImplTestSuite) TestDrawCircle() {
	latex := suite.drawex.DrawCircle(suite.fakeCircle)
	suite.NotEmpty(latex)
	suite.True(len(latex) > 0)
	suite.Equal(latex, latex, suite.circleQuery)
	suite.Empty(suite.drawex.DrawCircle(nil))
}
