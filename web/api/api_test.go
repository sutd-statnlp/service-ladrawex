package api_test

import (
	"github.com/sutd-statnlp/service-ladrawex/web/api"
	"github.com/gin-gonic/gin"
	"testing"
	"github.com/stretchr/testify/assert"
)


// CreateFakeDrawexRest creates a fake drawex rest for testing.
func CreateFakeDrawexRest() api.DrawexRest {
	return api.NewDrawexRest()
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

func TestDefaultDrawexRest(t *testing.T)  {
	assert.NotNil(t, api.DefaultDrawexRest())
}

func TestGroup(t *testing.T) {
	assert.True(t, api.Group(gin.New()))
}
