package bindata

// BinDataState represents the state of binary data.
type BinDataState byte

const (
	// NotAccess indicates that the data has not been accessed yet.
	NotAccess BinDataState = iota
	// SuccessAccess indicates that the data was successfully accessed.
	SuccessAccess
	// FailAccess indicates that access to the data failed.
	FailAccess
	// FailAccessButIgnore indicates that access to the data failed but was ignored.
	FailAccessButIgnore
)

// Value returns the byte value stored in the enum.
func (bds BinDataState) Value() byte {
	return byte(bds)
}

// ValueOf returns the corresponding BinDataState for a given byte value.
func StateValueOf(value byte) BinDataState {
	switch value {
	case byte(SuccessAccess):
		return SuccessAccess
	case byte(FailAccess):
		return FailAccess
	case byte(FailAccessButIgnore):
		return FailAccessButIgnore
	default:
		return NotAccess
	}
}
