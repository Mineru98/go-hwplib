package object

// RecordHeader represents a record header.
type RecordHeader struct {
	tagID uint16
	level uint16
	size  uint64
}

// NewRecordHeader creates a new RecordHeader instance.
func NewRecordHeader() *RecordHeader {
	return &RecordHeader{}
}

// GetTagID returns the tag ID.
func (rh *RecordHeader) GetTagID() uint16 {
	return rh.tagID
}

// SetTagID sets the tag ID.
func (rh *RecordHeader) SetTagID(tagID uint16) {
	rh.tagID = tagID
}

// GetLevel returns the level.
func (rh *RecordHeader) GetLevel() uint16 {
	return rh.level
}

// SetLevel sets the level.
func (rh *RecordHeader) SetLevel(level uint16) {
	rh.level = level
}

// GetSize returns the size.
func (rh *RecordHeader) GetSize() uint64 {
	return rh.size
}

// SetSize sets the size.
func (rh *RecordHeader) SetSize(size uint64) {
	rh.size = size
}

// Clone creates a new RecordHeader and copies the values.
func (rh *RecordHeader) Clone() *RecordHeader {
	return &RecordHeader{
		tagID: rh.tagID,
		level: rh.level,
		size:  rh.size,
	}
}
