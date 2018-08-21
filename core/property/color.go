package property

import "github.com/lucasb-eyer/go-colorful"

// Color is the property of component.
type Color struct {
	R uint8
	G uint8
	B uint8
}

// Hex returns hexdecimal value.
func (color *Color) Hex() string {
	colorful := colorful.Color{
		R: float64(color.R / 255),
		G: float64(color.G / 255),
		B: float64(color.B / 255),
	}
	return colorful.Hex()
}
