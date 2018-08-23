package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/log"
	"encoding/json"
	"github.com/sutd-statnlp/service-ladrawex/core/stringutil"
	"github.com/sutd-statnlp/service-ladrawex/core"
	"github.com/sutd-statnlp/service-ladrawex/web/apiutil"
)

// DrawexRest is the restful api of drawex.
type DrawexRest interface {
	PostDraw(context *gin.Context)
	DocumentFromRequestBody(requestBody *apiutil.RequestBody) *string
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
	var requestbody apiutil.RequestBody
	err = json.Unmarshal(data, &requestbody)
	if err != nil {
		log.Error(err)
		context.JSON(400, err.Error())
	}
	context.JSON(200, rest.DocumentFromRequestBody(&requestbody))
}

// DocumentFromRequestBody returns latex document from request body.
func (rest *DrawexRestImpl) DocumentFromRequestBody(requestBody *apiutil.RequestBody) *string  {
	log.Debug("DrawexRest request to get latex doc from request body: ", stringutil.JSON(requestBody))
	var latexs []*string
	for _, item := range requestBody.Rectangles {
		rect := apiutil.RectangleToComponent(&item)
		rectLatex := rest.drawex.DrawRectangle(rect)
		latexs = append(latexs, &rectLatex)
	}
	for _, item := range requestBody.Circles {
		circle := apiutil.CircleToComponent(&item)
		circleLatex := rest.drawex.DrawCircle(circle)
		latexs = append(latexs, &circleLatex)
	}
	for _, item := range requestBody.Lines {
		line := apiutil.LineToComponent(&item)
		lineLatex := rest.drawex.DrawLine(line)
		latexs = append(latexs, &lineLatex)
	}
	for _, item := range requestBody.Texts {
		text := apiutil.TextToComponent(&item)
		textLatex := rest.drawex.DrawText(text)
		latexs = append(latexs, &textLatex)
	}
	doc := rest.drawex.Document(latexs)
	log.Debug("Latex Doc: ", *doc)
	return doc
}