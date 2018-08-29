package api_test

import (
	"testing"

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

// CreareFakeDrawex creates fake drawex for testing.
func CreareFakeDrawex() core.Drawex {
	return core.NewDrawex()
}

// CreateFakeRouter creates a fake router for testing.
func CreateFakeRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	api.Group(router)
	return router
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
