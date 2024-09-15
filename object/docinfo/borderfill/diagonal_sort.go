package borderfill

// DiagonalSort represents the diagonal line types.
type DiagonalSort byte

const (
	DiagonalSortSlash DiagonalSort = iota
	BackSlash
	CrookedSlash
)

// Value returns the byte value of the DiagonalSort.
func (ds DiagonalSort) Value() byte {
	return byte(ds)
}

// ValueOf returns the DiagonalSort corresponding to the given byte value.
// If the value doesn't match any DiagonalSort, it defaults to Slash.
func DiagonalSortValueOf(value byte) DiagonalSort {
	switch DiagonalSort(value) {
	case DiagonalSortSlash:
		return DiagonalSortSlash
	case BackSlash:
		return BackSlash
	case CrookedSlash:
		return CrookedSlash
	default:
		return DiagonalSortSlash
	}
}
