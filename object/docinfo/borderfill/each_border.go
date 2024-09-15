package borderfill

import (
	"github.com/Mineru98/go-hwplib/object/etc"
)

// BorderType은 선의 종류를 나타냅니다.
type BorderType int

// BorderThickness는 선의 두께를 나타냅니다.
type BorderThickness int

// EachBorder는 테두리/배경 객체에서 4방향의 각각의 선을 나타내는 객체입니다.
type EachBorder struct {
	// 선 종류
	Type BorderType
	// 두께
	Thickness BorderThickness
	// 색상
	Color *etc.Color4Byte
}

// NewEachBorder는 새로운 EachBorder 객체를 생성하는 생성자입니다.
func NewEachBorder() *EachBorder {
	return &EachBorder{
		Color: etc.NewColor4Byte(),
	}
}

// GetType은 선의 종류를 반환합니다.
func (e *EachBorder) GetType() BorderType {
	return e.Type
}

// SetType은 선의 종류를 설정합니다.
func (e *EachBorder) SetType(t BorderType) {
	e.Type = t
}

// GetThickness는 선의 두께를 반환합니다.
func (e *EachBorder) GetThickness() BorderThickness {
	return e.Thickness
}

// SetThickness는 선의 두께를 설정합니다.
func (e *EachBorder) SetThickness(thickness BorderThickness) {
	e.Thickness = thickness
}

// GetColor는 선의 색상 객체를 반환합니다.
func (e *EachBorder) GetColor() *etc.Color4Byte {
	return e.Color
}

// Copy는 다른 EachBorder 객체의 값을 복사합니다.
func (e *EachBorder) Copy(from *EachBorder) {
	e.Type = from.Type
	e.Thickness = from.Thickness
	e.Color.Copy(from.Color)
}
