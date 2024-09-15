package etc

import (
	"bytes"
	"encoding/binary"
	"unicode/utf16"
)

type HWPString struct {
	bytes []byte
}

func NewHWPString() *HWPString {
	return &HWPString{bytes: nil}
}

func NewHWPStringWithBytes(b []byte) *HWPString {
	return &HWPString{bytes: b}
}

func (h *HWPString) GetBytes() []byte {
	return h.bytes
}

func (h *HWPString) SetBytes(b []byte) {
	h.bytes = b
}

func (h *HWPString) ToUTF16LEString() string {
	if h.bytes == nil {
		return ""
	}

	runes := utf16.Decode(h.byteSliceToUint16Slice(h.bytes))
	return string(runes)
}

func (h *HWPString) FromUTF16LEString(utf16LE string) {
	if utf16LE != "" {
		h.bytes = h.uint16SliceToByteSlice(utf16.Encode([]rune(utf16LE)))
	}
}

func (h *HWPString) Clone() *HWPString {
	cloned := NewHWPString()
	cloned.Copy(h)
	return cloned
}

func (h *HWPString) Copy(from *HWPString) {
	h.bytes = from.bytes
}

func (h *HWPString) GetWCharsSize() int {
	if h.bytes != nil {
		return 2 + len(h.bytes)
	}
	return 2
}

func (h *HWPString) Equals(other *HWPString) bool {
	return bytes.Equal(h.bytes, other.bytes)
}

// Helper functions
func (h *HWPString) byteSliceToUint16Slice(b []byte) []uint16 {
	u16s := make([]uint16, len(b)/2)
	for i := 0; i < len(u16s); i++ {
		u16s[i] = binary.LittleEndian.Uint16(b[i*2:])
	}
	return u16s
}

func (h *HWPString) uint16SliceToByteSlice(u16s []uint16) []byte {
	b := make([]byte, len(u16s)*2)
	for i, u16 := range u16s {
		binary.LittleEndian.PutUint16(b[i*2:], u16)
	}
	return b
}
