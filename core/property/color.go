package property

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/sutd-statnlp/service-ladrawex/log"
)

// Color is the property of component.
type Color struct {
	R uint8
	G uint8
	B uint8
}

// SetRGB sets rgb values.
func (color *Color) SetRGB(r, g, b uint8) {
	color.R = r
	color.G = g
	color.B = b
}

// RGB returns rgb values.
func (color *Color) RGB() (uint8, uint8, uint8) {
	return color.R, color.G, color.B
}

// Hex returns hexdecimal value.
func (color *Color) Hex() string {
	hexColor := colorful.LinearRgb(
		float64(color.R)/255,
		float64(color.G)/255,
		float64(color.B)/255,
	)
	return hexColor.Hex()
}

// FromHex sets rgb values from hex string.
func (color *Color) FromHex(hex string) bool {
	hexColor, err := colorful.Hex(hex)
	if err != nil {
		log.Error(err)
		return false
	}
	color.R = uint8(hexColor.R * 255)
	color.G = uint8(hexColor.G * 255)
	color.B = uint8(hexColor.B * 255)
	return true
}
