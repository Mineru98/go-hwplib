package borderfill

// CenterLineSort represents the type for center line sorts.
type CenterLineSort byte

const (
	// None represents no line.
	CenterLineSortNone CenterLineSort = iota
	// Horizontal represents a horizontal line.
	Horizontal
	// Vertical represents a vertical line.
	// Cross represents a cross line.
	Vertical
	Cross
)

// Value returns the byte value associated with the CenterLineSort.
func (cls CenterLineSort) Value() byte {
	return byte(cls)
}

// ValueOf returns the CenterLineSort corresponding to the given byte value.
// If the value does not match any enum, it returns None.
func CenterLineSortValueOf(value byte) CenterLineSort {
	switch value {
	case 1:
		return Horizontal
	case 2:
		return Vertical
	case 3:
		return Cross
	default:
		return CenterLineSortNone
	}
}
