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
	text.Common.BackgroundColor.SetRGB(255,255,255)
	return text
}
