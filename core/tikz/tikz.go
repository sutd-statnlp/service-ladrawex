package tikz

var (
	drawex *DrawexImpl
)

func init() {
	drawex = New()
}

// Default returns default tikz drawex.
func Default() *DrawexImpl {
	return drawex
}

// New creates new drawex.
func New() *DrawexImpl {
	return &DrawexImpl{}
}
