package borderfill

// SlashDiagonalShape represents the shape of a diagonal slash.
type SlashDiagonalShape byte

const (
	// SlashDiagonalShapeNone represents no diagonal slash.
	SlashDiagonalShapeNone SlashDiagonalShape = iota
	// Slash represents a slash from RightTop to LeftBottom ("/").
	SlashDiagonalShapeSlash = 2
	// RightTopToBottomEdge represents a diagonal from RightTop to Bottom Edge.
	RightTopToBottomEdge = 3
	// RightTopToLeftEdge represents a diagonal from RightTop to Left Edge.
	RightTopToLeftEdge = 6
	// RightTopToBottomLeftEdge represents a diagonal from RightTop to both Bottom and Left Edges.
	RightTopToBottomLeftEdge = 7
)

// Value returns the byte value stored in the enum.
func (sds SlashDiagonalShape) Value() byte {
	return byte(sds)
}

// ValueOf returns the SlashDiagonalShape corresponding to the given byte value.
func SlashDiagonalShapeValueOf(value byte) SlashDiagonalShape {
	switch value {
	case 2:
		return SlashDiagonalShapeSlash
	case 3:
		return RightTopToBottomEdge
	case 6:
		return RightTopToLeftEdge
	case 7:
		return RightTopToBottomLeftEdge
	default:
		return SlashDiagonalShapeNone
	}
}
