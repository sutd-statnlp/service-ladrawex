package component

import (
	"github.com/sutd-statnlp/service-ladrawex/core/property"
)

// NewRectangle creates new rectangle.
func NewRectangle() *Rectangle {
	return &Rectangle{
		Border:          property.NewBorder(),
		Color:           new(property.Color),
		BackgroundColor: new(property.Color),
		Size:            new(property.Size),
		Position:        new(property.Position),
	}
}
