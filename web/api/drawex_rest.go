package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/log"
	"encoding/json"
	"github.com/sutd-statnlp/service-ladrawex/core/stringutil"
	"github.com/sutd-statnlp/service-ladrawex/core"
)

// DrawexRest is the restful api of drawex.
type DrawexRest interface {
	PostDraw(context *gin.Context)
	DocumentFromRequestBody(requestBody *RequestBody) *string
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
	context.JSON(200, rest.DocumentFromRequestBody(&requestbody))
}

// DocumentFromRequestBody returns latex document from request body.
func (rest *DrawexRestImpl) DocumentFromRequestBody(requestBody *RequestBody) *string  {
	log.Debug("DrawexRest request to get latex doc from request body: ", stringutil.JSON(requestBody))
	var latexs []*string
	for _, item := range requestBody.Rectangles {
		rectLatex := rest.drawex.DrawRectangle(item)
		latexs = append(latexs, &rectLatex)
	}
	for _, item := range requestBody.Circles {
		circleLatex := rest.drawex.DrawCircle(item)
		latexs = append(latexs, &circleLatex)
	}
	for _, item := range requestBody.Lines {
		lineLatex := rest.drawex.DrawLine(item)
		latexs = append(latexs, &lineLatex)
	}
	for _, item := range requestBody.Texts {
		textLatex := rest.drawex.DrawText(item)
		latexs = append(latexs, &textLatex)
	}
	doc := rest.drawex.Document(latexs)
	log.Debug("Latex Doc: ", *doc)
	return doc
}