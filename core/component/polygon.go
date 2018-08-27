package component

import (
	"github.com/sutd-statnlp/service-ladrawex/core/property"
)

// Polygon is the polygon drawing component.
type Polygon struct {
	Sides  uint8
	Common *property.Common
}
