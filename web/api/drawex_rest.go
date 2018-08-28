package api

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/sutd-statnlp/service-ladrawex/core"
	"github.com/sutd-statnlp/service-ladrawex/core/stringutil"
	"github.com/sutd-statnlp/service-ladrawex/log"
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
func (rest *DrawexRestImpl) DocumentFromRequestBody(requestBody *RequestBody) *string {
	log.Debug("DrawexRest request to get latex doc from request body: ", stringutil.JSON(requestBody))
	var latexs []*string
	for _, item := range requestBody.Rectangles {
		latex := rest.drawex.DrawRectangle(item)
		latexs = append(latexs, &latex)
	}
	for _, item := range requestBody.Circles {
		latex := rest.drawex.DrawCircle(item)
		latexs = append(latexs, &latex)
	}
	for _, item := range requestBody.Lines {
		latex := rest.drawex.DrawLine(item)
		latexs = append(latexs, &latex)
	}
	for _, item := range requestBody.Texts {
		latex := rest.drawex.DrawText(item)
		latexs = append(latexs, &latex)
	}
	for _, item := range requestBody.Diamonds {
		latex := rest.drawex.DrawDiamond(item)
		latexs = append(latexs, &latex)
	}
	for _, item := range requestBody.Polygons {
		latex := rest.drawex.DrawPolygon(item)
		latexs = append(latexs, &latex)
	}
	for _, item := range requestBody.Connectors {
		latex := rest.drawex.DrawConnector(item)
		latexs = append(latexs, &latex)
	}
	for _, item := range requestBody.Triangles {
		latex := rest.drawex.DrawTriangle(item)
		latexs = append(latexs, &latex)
	}
	for _, item := range requestBody.Stars {
		latex := rest.drawex.DrawStar(item)
		latexs = append(latexs, &latex)
	}
	doc := rest.drawex.Document(latexs)
	log.Debug("Latex Doc: ", *doc)
	return doc
}
