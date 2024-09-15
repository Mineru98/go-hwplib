package bindata

// BinDataCompress는 바이너리 데이터의 압축 방법을 나타내는 enum입니다.
type BinDataCompress byte

const (
	// ByStorageDefault는 스토리지의 디폴트 모드를 따릅니다.
	ByStorageDefault BinDataCompress = iota
	// Compress는 무조건 압축을 의미합니다.
	Compress
	// NoCompress는 무조건 압축하지 않음을 의미합니다.
	NoCompress
)

// Value는 BinDataCompress의 정수값을 반환합니다.
func (bdc BinDataCompress) Value() byte {
	return byte(bdc)
}

// ValueOf는 주어진 정수값에 해당하는 BinDataCompress enum 값을 반환합니다.
func CompressValueOf(value byte) BinDataCompress {
	switch value {
	case byte(Compress):
		return Compress
	case byte(NoCompress):
		return NoCompress
	default:
		return ByStorageDefault
	}
}
