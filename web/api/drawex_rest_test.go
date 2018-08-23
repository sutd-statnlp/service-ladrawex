package api_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/sutd-statnlp/service-ladrawex/web/api"
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/gin-gonic/gin"
	"strings"
	"encoding/json"
	"github.com/sutd-statnlp/service-ladrawex/web/apiutil"
)

const (
	RequestJSON = `{"rectangles":[{"common":{"border":{"color":"#000000","thick":1},"backgroundColor":"#FFFFFF","size":{"width":40,"height":40},"position":{"x":100,"y":100},"text":{}}}],"circles":[{"common":{"border":{"color":"#000000","thick":1},"backgroundColor":"#FFFFFF","size":{"width":40,"height":40},"position":{"x":10,"y":40},"text":{}}}],"lines":[{"color":"#000000","width":2,"startPosition":{"x":500,"y":500},"endPosition":{"x":2000,"y":500}}],"texts":[{"common":{"border":{"color":null,"thick":1},"backgroundColor":"#000000","size":{"width":29.990234375,"height":22.599999999999998},"position":{"x":20,"y":20},"text":{"content":"text"}}}]}`
)

type DrawexRestTestSuite struct {
	suite.Suite
	rest            api.DrawexRest
	router          *gin.Engine
	fakeBody        *strings.Reader
}

func (suite *DrawexRestTestSuite) SetupSuite() {
	suite.router = CreateFakeRouter()
	suite.fakeBody = strings.NewReader(RequestJSON)
}

func (suite *DrawexRestTestSuite) SetupTest() {
	suite.rest = CreateFakeDrawexRest()
}

func TestDrawexRestTestSuite(t *testing.T) {
	suite.Run(t, new(DrawexRestTestSuite))
}

func (suite *DrawexRestTestSuite) TestPostDraw() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/draw", suite.fakeBody)
	suite.router.ServeHTTP(w, req)
	suite.Equal(200, w.Code)
	suite.NotNil(w.Body)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/draw", strings.NewReader(""))
	suite.router.ServeHTTP(w, req)
	suite.Equal(400, w.Code)
	suite.NotNil(w.Body)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/draw", strings.NewReader("[]"))
	suite.router.ServeHTTP(w, req)
	suite.Equal(400, w.Code)
	suite.NotNil(w.Body)
}

func (suite *DrawexRestTestSuite) TestDocumentFromRequestBody() {
	var fakeRequestBody apiutil.RequestBody
	err := json.Unmarshal([]byte(RequestJSON), &fakeRequestBody)
	suite.Nil(err)
	suite.NotNil(fakeRequestBody)


	doc := suite.rest.DocumentFromRequestBody(&fakeRequestBody)
	suite.NotNil(doc)
	suite.True(len(*doc) > 0)
}
