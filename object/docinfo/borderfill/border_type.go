package borderfill

// BorderType represents the type of border.
type BorderType uint8

const (
	// None represents no border.
	BorderTypeNone BorderType = iota
	// Solid represents a solid line.
	Solid
	// Dash represents a long dashed line.
	Dash
	// Dot represents a dotted line.
	Dot
	// DashDot represents a dash-dot line.
	DashDot
	// DashDotDot represents a dash-dot-dot line.
	DashDotDot
	// LongDash represents a longer dash line.
	LongDash
	// CircleDot represents a line with circle dots.
	CircleDot
	// Double represents a double line.
	Double
	// ThinThick represents a thin-thick double line.
	ThinThick
	// ThickThin represents a thick-thin double line.
	ThickThin
	// ThinThickThin represents a triple line with thin-thick-thin pattern.
	ThinThickThin
	// Wave represents a wavy line.
	Wave
	// DoubleWave represents a double wavy line.
	DoubleWave
	// Thick3D represents a thick 3D line.
	Thick3D
	// Thick3DReverseLighting represents a thick 3D line with reverse lighting.
	Thick3DReverseLighting
	// Solid3D represents a solid 3D line.
	Solid3D
	// Solid3DReverseLighting represents a solid 3D line with reverse lighting.
	Solid3DReverseLighting
)

// valueOf returns the corresponding BorderType for a given value.
func BorderTypeValueOf(value uint8) BorderType {
	if value >= uint8(BorderTypeNone) && value <= uint8(Solid3DReverseLighting) {
		return BorderType(value)
	}
	return Solid
}

// Value returns the uint8 value of the BorderType.
func (bt BorderType) Value() uint8 {
	return uint8(bt)
}
