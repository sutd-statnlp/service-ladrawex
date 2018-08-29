package api_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/sutd-statnlp/service-ladrawex/core/component"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/core"
	"github.com/sutd-statnlp/service-ladrawex/web/api"
)

// CreateFakeDrawexRest creates a fake drawex rest for testing.
func CreateFakeDrawexRest() api.DrawexRest {
	return api.NewDrawexRest()
}

// CreateFakeDrawexRestWithField creates a fake drawex rest with field for testing.
func CreateFakeDrawexRestWithField(drawex core.Drawex) api.DrawexRest {
	return api.NewDrawexRestWithField(drawex)
}

// CreateFakeDrawex creates fake drawex for testing.
func CreateFakeDrawex() core.Drawex {
	return core.NewDrawex()
}

// CreateFakeRouter creates a fake router for testing.
func CreateFakeRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	api.Group(router)
	return router
}

// CreateFakeRequestBody creates fake request body.
func CreateFakeRequestBody() *api.RequestBody {
	return &api.RequestBody{
		Rectangles: []*component.Rectangle{},
	}
}

func TestNewDrawexRest(t *testing.T) {
	assert.NotNil(t, api.NewDrawexRest())
}

func TestDefaultDrawexRest(t *testing.T) {
	assert.NotNil(t, api.DefaultDrawexRest())
}

func TestGroup(t *testing.T) {
	assert.True(t, api.Group(gin.New()))
}

func TestCreateDrawexRestWithField(t *testing.T) {
	drawex := core.NewDrawex()
	assert.NotNil(t, api.NewDrawexRestWithField(drawex))
}

type APITestSuite struct {
	suite.Suite
	fakeDrawex      core.Drawex
	fakeRequestBody *api.RequestBody
}

func (suite *APITestSuite) SetupSuite() {
	suite.fakeRequestBody = CreateFakeRequestBody()
	suite.fakeDrawex = CreateFakeDrawex()
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}

func (suite *APITestSuite) TestLatexFromRectangles() {
	latex := api.LatexFromRectangles(suite.fakeDrawex, suite.fakeRequestBody.Rectangles)
	suite.NotNil(latex)
}

func (suite *APITestSuite) TestLatexFromCircles() {
	latex := api.LatexFromCircles(suite.fakeDrawex, suite.fakeRequestBody.Circles)
	suite.NotNil(latex)
}

func (suite *APITestSuite) TestLatexFromLines() {
	latex := api.LatexFromLines(suite.fakeDrawex, suite.fakeRequestBody.Lines)
	suite.NotNil(latex)
}

func (suite *APITestSuite) TestLatexFromTexts() {
	latex := api.LatexFromTexts(suite.fakeDrawex, suite.fakeRequestBody.Texts)
	suite.NotNil(latex)
}

func (suite *APITestSuite) TestLatexFromDiamons() {
	latex := api.LatexFromDiamons(suite.fakeDrawex, suite.fakeRequestBody.Diamonds)
	suite.NotNil(latex)
}

func (suite *APITestSuite) TestLatexFromPolygons() {
	latex := api.LatexFromPolygons(suite.fakeDrawex, suite.fakeRequestBody.Polygons)
	suite.NotNil(latex)
}

func (suite *APITestSuite) TestLatexFromConnectors() {
	latex := api.LatexFromConnectors(suite.fakeDrawex, suite.fakeRequestBody.Connectors)
	suite.NotNil(latex)
}

func (suite *APITestSuite) TestLatexFromTriangles() {
	latex := api.LatexFromTriangles(suite.fakeDrawex, suite.fakeRequestBody.Triangles)
	suite.NotNil(latex)
}

func (suite *APITestSuite) TestLatexFromStars() {
	latex := api.LatexFromStars(suite.fakeDrawex, suite.fakeRequestBody.Stars)
	suite.NotNil(latex)
}

func (suite *APITestSuite) TestDocumentFromRequestBody() {
	var fakeRequestBody api.RequestBody
	err := json.Unmarshal([]byte(RequestJSON), &fakeRequestBody)
	suite.Nil(err)
	suite.NotNil(fakeRequestBody)

	doc := api.DocumentFromRequestBody(suite.fakeDrawex, &fakeRequestBody)
	suite.NotNil(doc)
	suite.True(len(*doc) > 0)
}
