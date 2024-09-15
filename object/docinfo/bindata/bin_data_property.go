package bindata

import (
	"github.com/Mineru98/go-hwplib/util/binary"
)

// BinDataType, BinDataCompress, BinDataState 는 별도로 정의해 주세요.

type BinDataProperty struct {
	value uint16 // 2 byte unsigned integer
}

// NewBinDataProperty 생성자
func NewBinDataProperty() *BinDataProperty {
	return &BinDataProperty{}
}

// GetValue 파일에 저장되는 정수값을 반환한다.
func (b *BinDataProperty) GetValue() uint16 {
	return b.value
}

// SetValue 파일에 저장되는 정수값을 설정한다.
func (b *BinDataProperty) SetValue(value uint16) {
	b.value = value
}

// GetType 바이너리 데이터의 타입을 반환한다. (0~3 BitFlag)
func (b *BinDataProperty) GetType() BinDataType {
	return BinDataType(binary.GetRange(uint64(b.value), 0, 3))
}

// SetType 바이너리 데이터의 타입을 설정한다. (0~3 BitFlag)
func (b *BinDataProperty) SetType(t BinDataType) {
	b.value = uint16(binary.SetRange(uint64(b.value), 0, 3, int(t)))
}

// GetCompress 바이너리 데이터의 압축 방법을 반환한다. (4~5 BitFlag)
func (b *BinDataProperty) GetCompress() BinDataCompress {
	return BinDataCompress(binary.GetRange(uint64(b.value), 4, 5))
}

// SetCompress 바이너리 데이터의 압축 방법을 설정한다. (4~5 BitFlag)
func (b *BinDataProperty) SetCompress(c BinDataCompress) {
	b.value = uint16(binary.SetRange(uint64(b.value), 4, 5, int(c)))
}

// GetState 바이너리 데이터의 상태를 반환한다. (8~9 BitFlag)
func (b *BinDataProperty) GetState() BinDataState {
	return BinDataState(binary.GetRange(uint64(b.value), 8, 9))
}

// SetState 바이너리 데이터의 상태를 설정한다. (8~9 BitFlag)
func (b *BinDataProperty) SetState(s BinDataState) {
	b.value = uint16(binary.SetRange(uint64(b.value), 8, 9, int(s)))
}

// Copy 다른 BinDataProperty 객체의 값을 복사한다.
func (b *BinDataProperty) Copy(from *BinDataProperty) {
	b.value = from.value
}
