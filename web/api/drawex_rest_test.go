package api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"github.com/sutd-statnlp/service-ladrawex/web/api"
)

const (
	RequestJSON = `{"rectangles":[{"common":{"border":{"color":{"r":0,"g":0,"b":0},"thick":0.5},"backgroundColor":{"r":0,"g":0,"b":0},"size":{"width":10,"height":10},"position":{"x":3.3333333333333335,"y":10},"text":{}}}],"circles":[{"common":{"border":{"color":{"r":0,"g":0,"b":0},"thick":0.5},"backgroundColor":{"r":0,"g":0,"b":0},"size":{"width":10,"height":10},"position":{"x":0.3333333333333333,"y":12},"text":{}}},{"common":{"border":{"color":{"r":0,"g":0,"b":0},"thick":0.5},"backgroundColor":{"r":0,"g":0,"b":0},"size":{"width":10,"height":10},"position":{"x":4.733333333333333,"y":12},"text":{}}}],"lines":[{"color":{"r":0,"g":0,"b":0},"width":1,"startPosition":{"x":0.8333333333333334,"y":32.5},"endPosition":{"x":3.3333333333333335,"y":32.5}}],"texts":[{"common":{"border":{"color":{"r":0,"g":0,"b":0},"thick":0.5},"backgroundColor":{"r":0,"g":0,"b":0},"size":{"width":7.49755859375,"height":5.6499999999999995},"position":{"x":0.6666666666666666,"y":12.666666666666666},"text":{"content":"text"}}}]}`
)

type DrawexRestTestSuite struct {
	suite.Suite
	rest     api.DrawexRest
	router   *gin.Engine
	fakeBody *strings.Reader
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
