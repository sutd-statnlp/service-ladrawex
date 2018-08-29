package api

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/core"
	"github.com/sutd-statnlp/service-ladrawex/log"
)

// DrawexRest is the restful api of drawex.
type DrawexRest interface {
	PostDraw(context *gin.Context)
}

// DrawexRestImpl is the implementation of DrawexRest interface.
type DrawexRestImpl struct {
	drawex core.Drawex
}

// PostDraw handles post drawing request.
func (rest *DrawexRestImpl) PostDraw(context *gin.Context) {
	log.Debug("DrawexRest request to post drawing")
	data, err := context.GetRawData()
	if err != nil {
		log.Error(err)
		context.JSON(400, err.Error())
	}
	var requestbody RequestBody
	err = json.Unmarshal(data, &requestbody)
	if err != nil {
		log.Error(err)
		context.JSON(400, err.Error())
	}
	doc := DocumentFromRequestBody(rest.drawex, &requestbody)
	log.Debug("Latex Doc: ", *doc)
	context.JSON(200, doc)
}
