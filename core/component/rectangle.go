package component

import "github.com/sutd-statnlp/service-ladrawex/core/property"

// Rectangle is the rectangle drawing component.
type Rectangle struct {
	Border          *property.Border
	Color           *property.Color
	BackgroundColor *property.Color
	Size            *property.Size
	Position        *property.Position
}
