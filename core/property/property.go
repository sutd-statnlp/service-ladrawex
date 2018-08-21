package property

// NewBorder creates new border.
func NewBorder() *Border {
	return &Border{
		Color: new(Color),
	}
}
