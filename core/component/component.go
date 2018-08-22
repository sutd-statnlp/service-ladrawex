package component

import (
	"github.com/sutd-statnlp/service-ladrawex/core/property"
)

// NewRectangle creates new rectangle.
func NewRectangle() *Rectangle {
	return &Rectangle{
		Common: NewCommonProperty(),
	}
}

// NewCommonProperty creates new common property.
func NewCommonProperty() *property.Common {
	return &property.Common{
		Border:          property.NewBorder(),
		BackgroundColor: new(property.Color),
		Size:            new(property.Size),
		Position:        new(property.Position),
	}
}

// NewCircle creates new circle.
func NewCircle() *Circle {
	return &Circle{
		Common: NewCommonProperty(),
	}
}
