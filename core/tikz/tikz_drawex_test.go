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
	suite.Equal(latex, `\draw (0,0) rectangle (0,0);`)

	suite.Empty(suite.drawex.DrawRectangle(nil))
}
