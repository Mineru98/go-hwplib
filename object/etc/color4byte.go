package etc

import "github.com/Mineru98/go-hwplib/util/binary"

// Color4Byte는 4바이트 색상 객체입니다. Windows API의 COLORREF에 해당합니다.
type Color4Byte struct {
	value uint64
}

// NewColor4Byte는 새로운 Color4Byte 객체를 생성하는 생성자입니다.
func NewColor4Byte() *Color4Byte {
	return &Color4Byte{}
}

// NewColor4ByteWithRGBA는 RGBA 값을 받아 Color4Byte 객체를 생성하는 생성자입니다.
func NewColor4ByteWithRGBA(r, g, b, a int) *Color4Byte {
	color := &Color4Byte{}
	color.SetR(uint16(r))
	color.SetG(uint16(g))
	color.SetB(uint16(b))
	color.SetA(uint16(a))
	return color
}

// NewColor4ByteWithRGB는 RGB 값을 받아 Color4Byte 객체를 생성하는 생성자입니다.
func NewColor4ByteWithRGB(r, g, b int) *Color4Byte {
	color := &Color4Byte{}
	color.SetR(uint16(r))
	color.SetG(uint16(g))
	color.SetB(uint16(b))
	color.SetA(0)
	return color
}

// GetValue는 unsigned 4 byte color 값을 반환합니다.
func (c *Color4Byte) GetValue() uint64 {
	return c.value
}

// SetValue는 unsigned 4 byte color 값을 설정합니다.
func (c *Color4Byte) SetValue(value uint64) {
	c.value = value
}

// GetR는 색상의 red 값을 반환합니다.
func (c *Color4Byte) GetR() uint16 {
	return uint16(binary.GetRange(c.value, 0, 7))
}

// SetR는 색상의 red 값을 설정합니다.
func (c *Color4Byte) SetR(r uint16) {
	c.value = binary.SetRange(c.value, 0, 7, int(r))
}

// GetG는 색상의 green 값을 반환합니다.
func (c *Color4Byte) GetG() uint16 {
	return uint16(binary.GetRange(c.value, 8, 15))
}

// SetG는 색상의 green 값을 설정합니다.
func (c *Color4Byte) SetG(g uint16) {
	c.value = binary.SetRange(c.value, 8, 15, int(g))
}

// GetB는 색상의 blue 값을 반환합니다.
func (c *Color4Byte) GetB() uint16 {
	return uint16(binary.GetRange(c.value, 16, 23))
}

// SetB는 색상의 blue 값을 설정합니다.
func (c *Color4Byte) SetB(b uint16) {
	c.value = binary.SetRange(c.value, 16, 23, int(b))
}

// GetA는 색상의 alpha 값을 반환합니다.
func (c *Color4Byte) GetA() uint16 {
	return uint16(binary.GetRange(c.value, 24, 31))
}

// SetA는 색상의 alpha 값을 설정합니다.
func (c *Color4Byte) SetA(a uint16) {
	c.value = binary.SetRange(c.value, 24, 31, int(a))
}

// Clone은 Color4Byte 객체를 복제합니다.
func (c *Color4Byte) Clone() *Color4Byte {
	cloned := NewColor4Byte()
	cloned.Copy(c)
	return cloned
}

// Copy는 다른 Color4Byte 객체의 값을 복사합니다.
func (c *Color4Byte) Copy(from *Color4Byte) {
	c.value = from.value
}
