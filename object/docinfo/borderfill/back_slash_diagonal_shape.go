package borderfill

// BackSlashDiagonalShape represents the diagonal shape of a backslash
type BackSlashDiagonalShape byte

const (
	// None represents no diagonal shape
	BackSlashDiagonalShapeNone BackSlashDiagonalShape = iota
	// BackSlash represents a backslash diagonal (leftTop --> rightBottom) "\"
	BackSlashDiagonalShapeBackSlash = 2
	// LeftTopToBottomEdge represents diagonal from LeftTop to Bottom Edge
	LeftTopToBottomEdge = 3
	// LeftTopToRightEdge represents diagonal from LeftTop to Right Edge
	LeftTopToRightEdge = 6
	// LeftTopToBottomRightEdge represents diagonal from LeftTop to Bottom & Right Edge
	LeftTopToBottomRightEdge = 7
)

// Value returns the byte value of the BackSlashDiagonalShape
func (b BackSlashDiagonalShape) Value() byte {
	return byte(b)
}

// ValueOf returns the BackSlashDiagonalShape corresponding to the given byte value
func BackSlashDiagonalShapeValueOf(value byte) BackSlashDiagonalShape {
	switch value {
	case 2:
		return BackSlashDiagonalShapeBackSlash
	case 3:
		return LeftTopToBottomEdge
	case 6:
		return LeftTopToRightEdge
	case 7:
		return LeftTopToBottomRightEdge
	default:
		return BackSlashDiagonalShapeNone
	}
}
