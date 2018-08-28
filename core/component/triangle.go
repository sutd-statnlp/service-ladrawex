package component

import (
	"github.com/sutd-statnlp/service-ladrawex/core/property"
)

// Triangle is the triangle drawing component.
type Triangle struct {
	SecondPosition *property.Position
	ThirdPosition  *property.Position
	Common         *property.Common
}
