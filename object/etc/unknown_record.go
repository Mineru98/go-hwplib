package etc

import (
	"github.com/Mineru98/go-hwplib/object"
)

// UnknownRecord represents an unknown record with a header and body.
type UnknownRecord struct {
	header *object.RecordHeader
	body   []byte
}

// NewUnknownRecord creates a new UnknownRecord instance.
func NewUnknownRecord() *UnknownRecord {
	return &UnknownRecord{}
}

// NewUnknownRecordWithHeader creates a new UnknownRecord instance with a given header.
func NewUnknownRecordWithHeader(header *object.RecordHeader) *UnknownRecord {
	return &UnknownRecord{
		header: header.Clone(),
	}
}

// GetHeader returns the record header.
func (ur *UnknownRecord) GetHeader() *object.RecordHeader {
	return ur.header
}

// SetHeader sets the record header.
func (ur *UnknownRecord) SetHeader(header *object.RecordHeader) {
	ur.header = header.Clone()
}

// GetBody returns the record body data.
func (ur *UnknownRecord) GetBody() []byte {
	return ur.body
}

// SetBody sets the record body data.
func (ur *UnknownRecord) SetBody(body []byte) {
	ur.body = make([]byte, len(body))
	copy(ur.body, body)
}

// Clone creates a copy of the UnknownRecord instance.
func (ur *UnknownRecord) Clone() *UnknownRecord {
	cloned := &UnknownRecord{}
	if ur.header != nil {
		cloned.header = ur.header.Clone()
	}
	if ur.body != nil {
		cloned.body = make([]byte, len(ur.body))
		copy(cloned.body, ur.body)
	}
	return cloned
}
