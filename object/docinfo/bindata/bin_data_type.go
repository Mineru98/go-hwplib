package bindata

// BinDataType은 바이너리 데이터의 타입을 나타냅니다.
type BinDataType byte

const (
	// Link는 그림 외부 파일 참조를 나타냅니다.
	Link BinDataType = iota
	// Embedding은 그림 파일 포함을 나타냅니다.
	Embedding
	// Storage는 OLE 포함을 나타냅니다.
	Storage
)

// Value는 파일에 저장되는 정수값을 반환합니다.
func (bdt BinDataType) Value() byte {
	return byte(bdt)
}

// ValueOf는 파일에 저장되는 정수값에 해당되는 enum 값을 반환합니다.
func TypeValueOf(value byte) BinDataType {
	switch value {
	case Link.Value():
		return Link
	case Embedding.Value():
		return Embedding
	case Storage.Value():
		return Storage
	default:
		return Link
	}
}
