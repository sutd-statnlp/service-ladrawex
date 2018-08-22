package core

import "github.com/sutd-statnlp/service-ladrawex/core/component"

// Drawex is the latex drawer, which provides drawing features and generating latex.
type Drawex interface {
	DrawRectangle(rect *component.Rectangle) string
	DrawCircle(rect *component.Circle) string
}
