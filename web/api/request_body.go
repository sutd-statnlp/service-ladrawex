package api

import "github.com/sutd-statnlp/service-ladrawex/core/component"


// RequestBody is the body of post request for drawing.
type RequestBody struct {
	Rectangles []*component.Rectangle
	Circles []*component.Circle
	Lines []*component.Line
	Texts []*component.Text
}


