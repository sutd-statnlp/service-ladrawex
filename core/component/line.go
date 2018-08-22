package component

import "github.com/sutd-statnlp/service-ladrawex/core/property"

// Line is the line drawing component.
type Line struct {
	Color         *property.Color
	Width         float32
	StartPosition *property.Position
	EndPosition   *property.Position
}
