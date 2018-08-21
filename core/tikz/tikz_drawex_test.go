package tikz_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/sutd-statnlp/service-ladrawex/core/tikz"
)

type DrawexImplTestSuite struct {
	suite.Suite
	drawex *tikz.DrawexImpl
}

func (suite *DrawexImplTestSuite) SetupTest() {
	suite.drawex = tikz.New()
}

func TestDrawexImplTestSuite(t *testing.T) {
	suite.Run(t, new(DrawexImplTestSuite))
}

func (suite *DrawexImplTestSuite) TestDrawRectangle() {
	latex := suite.drawex.DrawRectangle(CreateFakeRectangle())
	suite.NotNil(latex)
	suite.NotEmpty(latex)
	suite.True(len(latex) > 0)
	suite.Equal(latex, `\node [draw={rgb,255:red,0;green,0;blue,0}, line width=0mm, shape=rectangle, fill={rgb,255:red,0;green,0;blue,0}, minimum width=0mm, minimum height=0mm] at (0,0) {};`)
	suite.Empty(suite.drawex.DrawRectangle(nil))
}
