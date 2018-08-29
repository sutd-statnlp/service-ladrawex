package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/core"
)

var (
	drawexRest DrawexRest
)

func init() {
	drawexRest = NewDrawexRest()
}

// NewDrawexRest creates new drawex rest api.
func NewDrawexRest() DrawexRest {
	return &DrawexRestImpl{
		drawex: core.DefaultDrawex(),
	}
}

// DefaultDrawexRest returns the default drawex rest api.
func DefaultDrawexRest() DrawexRest {
	return drawexRest
}

// Group returns the api group.
func Group(router *gin.Engine) bool {
	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/draw", drawexRest.PostDraw)
	}
	return apiGroup != nil
}

// NewDrawexRestWithField creates new drawex rest with field.
func NewDrawexRestWithField(drawex core.Drawex) DrawexRest {
	return &DrawexRestImpl{
		drawex: drawex,
	}
}
