package component

import (
	"github.com/sutd-statnlp/service-ladrawex/core/property"
)

const (
	// DefaultTextContent is the default content for text component.
	DefaultTextContent = "text"
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
		Text:            new(property.Text),
	}
}

// NewCircle creates new circle.
func NewCircle() *Circle {
	return &Circle{
		Common: NewCommonProperty(),
	}
}

// NewLine creates new line.
func NewLine() *Line {
	return &Line{
		Color:         new(property.Color),
		StartPosition: new(property.Position),
		EndPosition:   new(property.Position),
	}
}

// NewText creates new text.
func NewText() *Text {
	text := &Text{
		Common: NewCommonProperty(),
	}
	text.Common.Text.Content = DefaultTextContent
	text.Common.BackgroundColor.SetRGB(255, 255, 255)
	return text
}

// NewTriangle creates new triangle component.
func NewTriangle() *Triangle {
	item := &Triangle{
		Common:         NewCommonProperty(),
		SecondPosition: new(property.Position),
		ThirdPosition:  new(property.Position),
	}
	item.SecondPosition.X = 4
	item.ThirdPosition.X = 4
	item.ThirdPosition.Y = 4
	return item
}

// NewPolygon creates new polygon component.
func NewPolygon() *Polygon {
	return &Polygon{
		Sides:  5,
		Common: NewCommonProperty(),
	}
}

// NewDiamond creates new diamond component.
func NewDiamond() *Diamond {
	return &Diamond{
		Common: NewCommonProperty(),
	}
}

// NewConnector creates new connector component.
func NewConnector() *Connector {
	return &Connector{
		Line: NewLine(),
	}
}

// NewStar creates new start component.
func NewStar() *Star {
	return &Star{
		Common: NewCommonProperty(),
	}
}
