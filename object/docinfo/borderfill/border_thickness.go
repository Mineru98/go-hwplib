package borderfill

// BorderThickness defines the thickness of the border
type BorderThickness byte

const (
	MM0_1  BorderThickness = iota // 0.1 mm
	MM0_12                        // 0.12 mm
	MM0_15                        // 0.15 mm
	MM0_2                         // 0.2 mm
	MM0_25                        // 0.25 mm
	MM0_3                         // 0.3 mm
	MM0_4                         // 0.4 mm
	MM0_5                         // 0.5 mm
	MM0_6                         // 0.6 mm
	MM0_7                         // 0.7 mm
	MM1_0                         // 1.0 mm
	MM1_5                         // 1.5 mm
	MM2_0                         // 2.0 mm
	MM3_0                         // 3.0 mm
	MM4_0                         // 4.0 mm
	MM5_0                         // 5.0 mm
)

// Value returns the byte value stored in the file
func (bt BorderThickness) Value() byte {
	return byte(bt)
}

// ValueOf returns the corresponding enum value for the given byte value
func ValueOf(value byte) BorderThickness {
	switch value {
	case 0:
		return MM0_1
	case 1:
		return MM0_12
	case 2:
		return MM0_15
	case 3:
		return MM0_2
	case 4:
		return MM0_25
	case 5:
		return MM0_3
	case 6:
		return MM0_4
	case 7:
		return MM0_5
	case 8:
		return MM0_6
	case 9:
		return MM0_7
	case 10:
		return MM1_0
	case 11:
		return MM1_5
	case 12:
		return MM2_0
	case 13:
		return MM3_0
	case 14:
		return MM4_0
	case 15:
		return MM5_0
	default:
		return MM0_1
	}
}
