package util

// isEmpty checks if a boolean array is empty or nil.
func isEmptyBool(array []bool) bool {
	return len(array) == 0
}

// isEmpty checks if a byte array is empty or nil.
func isEmptyByte(array []byte) bool {
	return len(array) == 0
}

// isEmpty checks if a char array is empty or nil.
func isEmptyRune(array []rune) bool {
	return len(array) == 0
}

// isEmpty checks if a float64 array is empty or nil.
func isEmptyFloat64(array []float64) bool {
	return len(array) == 0
}

// isEmpty checks if a float32 array is empty or nil.
func isEmptyFloat32(array []float32) bool {
	return len(array) == 0
}

// isEmpty checks if an int array is empty or nil.
func isEmptyInt(array []int) bool {
	return len(array) == 0
}

// isEmpty checks if an int64 array is empty or nil.
func isEmptyInt64(array []int64) bool {
	return len(array) == 0
}

// isEmpty checks if an int16 array is empty or nil.
func isEmptyInt16(array []int16) bool {
	return len(array) == 0
}

// isEmpty checks if an interface array is empty or nil.
func isEmptyInterface(array []interface{}) bool {
	return len(array) == 0
}

// isNotEmpty checks if a boolean array is not empty and not nil.
func isNotEmptyBool(array []bool) bool {
	return !isEmptyBool(array)
}

// isNotEmpty checks if a byte array is not empty and not nil.
func isNotEmptyByte(array []byte) bool {
	return !isEmptyByte(array)
}

// isNotEmpty checks if a char array is not empty and not nil.
func isNotEmptyRune(array []rune) bool {
	return !isEmptyRune(array)
}

// isNotEmpty checks if a float64 array is not empty and not nil.
func isNotEmptyFloat64(array []float64) bool {
	return !isEmptyFloat64(array)
}

// isNotEmpty checks if a float32 array is not empty and not nil.
func isNotEmptyFloat32(array []float32) bool {
	return !isEmptyFloat32(array)
}

// isNotEmpty checks if an int array is not empty and not nil.
func isNotEmptyInt(array []int) bool {
	return !isEmptyInt(array)
}

// isNotEmpty checks if an int64 array is not empty and not nil.
func isNotEmptyInt64(array []int64) bool {
	return !isEmptyInt64(array)
}

// isNotEmpty checks if an int16 array is not empty and not nil.
func isNotEmptyInt16(array []int16) bool {
	return !isEmptyInt16(array)
}

// isNotEmpty checks if an interface array is not empty and not nil.
func isNotEmptyInterface(array []interface{}) bool {
	return !isEmptyInterface(array)
}
