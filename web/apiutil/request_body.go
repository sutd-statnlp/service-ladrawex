package apiutil

type Border struct {
	Color string
	Thick float32
}

type Size struct {
	Width float32
	Height float32
}

type Position struct {
	X float32
	Y float32
}

type TextProperty struct {
	Content string
}

type Common struct {
	Border Border
	BackgroundColor string
	Size Size
	Position Position
	Text TextProperty
}

type Rectangle struct {
	Common Common
}

type Circle struct {
	Common Common
}

type Line struct {
	Color string
	Width float32
	StartPosition Position
	EndPosition Position
}

type Text struct {
	Common Common
}

// RequestBody is the body of post request for drawing.
type RequestBody struct {
	Rectangles []Rectangle
	Circles []Circle
	Lines []Line
	Texts []Text
}


