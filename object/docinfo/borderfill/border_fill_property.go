package borderfill

import "github.com/Mineru98/go-hwplib/util/binary"

// BorderFillProperty represents the properties of a border/fill object.
type BorderFillProperty struct {
	// The integer value saved in the file (unsigned 2 bytes)
	value uint64
}

// NewBorderFillProperty creates a new instance of BorderFillProperty.
func NewBorderFillProperty() *BorderFillProperty {
	return &BorderFillProperty{}
}

// GetValue returns the integer value saved in the file.
func (b *BorderFillProperty) GetValue() uint64 {
	return b.value
}

// SetValue sets the integer value saved in the file.
func (b *BorderFillProperty) SetValue(value uint64) {
	b.value = value
}

// Is3DEffect returns whether the 3D effect is enabled (0 bit).
func (b *BorderFillProperty) Is3DEffect() bool {
	return binary.Get(b.value, 0)
}

// Set3DEffect sets whether the 3D effect is enabled (0 bit).
func (b *BorderFillProperty) Set3DEffect(effect bool) {
	b.value = binary.Set(b.value, 0, effect)
}

// IsShadowEffect returns whether the shadow effect is enabled (1 bit).
func (b *BorderFillProperty) IsShadowEffect() bool {
	return binary.Get(b.value, 1)
}

// SetShadowEffect sets whether the shadow effect is enabled (1 bit).
func (b *BorderFillProperty) SetShadowEffect(effect bool) {
	b.value = binary.Set(b.value, 1, effect)
}

// GetSlashDiagonalShape returns the Slash diagonal shape (2-4 bits).
func (b *BorderFillProperty) GetSlashDiagonalShape() SlashDiagonalShape {
	return SlashDiagonalShape(binary.GetRange(b.value, 2, 4))
}

// SetSlashDiagonalShape sets the Slash diagonal shape (2-4 bits).
func (b *BorderFillProperty) SetSlashDiagonalShape(shape SlashDiagonalShape) {
	b.value = binary.SetRange(b.value, 2, 4, int(shape))
}

// GetBackSlashDiagonalShape returns the BackSlash diagonal shape (5-7 bits).
func (b *BorderFillProperty) GetBackSlashDiagonalShape() BackSlashDiagonalShape {
	return BackSlashDiagonalShape(binary.GetRange(b.value, 5, 7))
}

// SetBackSlashDiagonalShape sets the BackSlash diagonal shape (5-7 bits).
func (b *BorderFillProperty) SetBackSlashDiagonalShape(shape BackSlashDiagonalShape) {
	b.value = binary.SetRange(b.value, 5, 7, int(shape))
}

// IsBrokenSlashDiagonal returns whether the Slash diagonal is broken (8-9 bits).
func (b *BorderFillProperty) IsBrokenSlashDiagonal() bool {
	return binary.Get(b.value, 8) || binary.Get(b.value, 9)
}

// SetBrokenSlashDiagonal sets whether the Slash diagonal is broken (8-9 bits).
func (b *BorderFillProperty) SetBrokenSlashDiagonal(broken bool) {
	b.value = binary.Set(b.value, 8, broken)
	b.value = binary.Set(b.value, 9, broken)
}

// IsBrokenBackSlashDiagonal returns whether the BackSlash diagonal is broken (10 bit).
func (b *BorderFillProperty) IsBrokenBackSlashDiagonal() bool {
	return binary.Get(b.value, 10)
}

// SetBrokenBackSlashDiagonal sets whether the BackSlash diagonal is broken (10 bit).
func (b *BorderFillProperty) SetBrokenBackSlashDiagonal(broken bool) {
	b.value = binary.Set(b.value, 10, broken)
}

// IsRotateSlashDiagonal180 returns whether the Slash diagonal is rotated by 180 degrees (11 bit).
func (b *BorderFillProperty) IsRotateSlashDiagonal180() bool {
	return binary.Get(b.value, 11)
}

// SetRotateSlashDiagonal180 sets whether the Slash diagonal is rotated by 180 degrees (11 bit).
func (b *BorderFillProperty) SetRotateSlashDiagonal180(rotated bool) {
	b.value = binary.Set(b.value, 11, rotated)
}

// IsRotateBackSlashDiagonal180 returns whether the BackSlash diagonal is rotated by 180 degrees (12 bit).
func (b *BorderFillProperty) IsRotateBackSlashDiagonal180() bool {
	return binary.Get(b.value, 12)
}

// SetRotateBackSlashDiagonal180 sets whether the BackSlash diagonal is rotated by 180 degrees (12 bit).
func (b *BorderFillProperty) SetRotateBackSlashDiagonal180(rotated bool) {
	b.value = binary.Set(b.value, 12, rotated)
}

// HasCenterLine returns whether the center line is present (13 bit).
func (b *BorderFillProperty) HasCenterLine() bool {
	return binary.Get(b.value, 13)
}

// SetHasCenterLine sets whether the center line is present (13 bit).
func (b *BorderFillProperty) SetHasCenterLine(hasLine bool) {
	b.value = binary.Set(b.value, 13, hasLine)
}

// GetCenterLineSort returns the type of center line (8-9 bits).
func (b *BorderFillProperty) GetCenterLineSort() CenterLineSort {
	return CenterLineSort(binary.GetRange(b.value, 8, 9))
}

// SetCenterLineSort sets the type of center line (8-9 bits).
func (b *BorderFillProperty) SetCenterLineSort(sort CenterLineSort) {
	b.value = binary.SetRange(b.value, 8, 9, int(sort))
}

// Copy copies the properties from another BorderFillProperty instance.
func (b *BorderFillProperty) Copy(from *BorderFillProperty) {
	b.value = from.value
}
