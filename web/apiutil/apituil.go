package apiutil

import (
	"github.com/sutd-statnlp/service-ladrawex/core/component"
	"github.com/sutd-statnlp/service-ladrawex/core/property"
)

// RectangleToComponent convert requesting rectangle into real component.
func RectangleToComponent(requestRect *Rectangle) *component.Rectangle {
	return &component.Rectangle{
		Common: CommonToProperty(&requestRect.Common),
	}
}

// CommonToProperty convert requesting common property into real common property.
func CommonToProperty(requestCommon *Common) *property.Common {
	borderColor := new(property.Color)
	borderColor.FromHex(requestCommon.Border.Color)

	backgroundColor := new(property.Color)
	backgroundColor.FromHex(requestCommon.BackgroundColor)

	return &property.Common{
		Border: &property.Border{
			Color: borderColor,
			Thick: requestCommon.Border.Thick,
		},
		Size: &property.Size{
			Width:  requestCommon.Size.Width,
			Height: requestCommon.Size.Height,
		},
		Position: &property.Position{
			XAxis: requestCommon.Position.X,
			YAxis: requestCommon.Position.Y,
		},
		BackgroundColor: backgroundColor,
		Text: &property.Text{
			Content: requestCommon.Text.Content,
		},
	}
}

// CircleToComponent converts the requesting circle into real component.
func CircleToComponent(requestCircle *Circle) *component.Circle {
	return &component.Circle{
		Common: CommonToProperty(&requestCircle.Common),
	}
}

// LineToComponent converts the requesting line into real component.
func LineToComponent(requestLine *Line) *component.Line {
	color := new(property.Color)
	color.FromHex(requestLine.Color)
	return &component.Line{
		Color: color,
		Width: requestLine.Width,
		StartPosition: &property.Position{
			XAxis: requestLine.StartPosition.X,
			YAxis: requestLine.StartPosition.Y,
		},
		EndPosition: &property.Position{
			XAxis: requestLine.EndPosition.X,
			YAxis: requestLine.EndPosition.Y,
		},
	}
}

// TextToComponent converts the requesting text into real component.
func TextToComponent(requestText *Text) *component.Text {
	return &component.Text{
		Common: CommonToProperty(&requestText.Common),
	}
}
