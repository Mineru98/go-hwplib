package compoundfile

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io"
	"io/ioutil"

	"github.com/Mineru98/go-hwplib/object"
	"github.com/Mineru98/go-hwplib/object/fileheader"
)

type StreamReader struct {
	is              io.ReadCloser
	size            int64
	read            int64
	header          object.RecordHeader
	readAfterHeader int64
	fileVersion     fileheader.FileVersion
	docInfo         *DocInfo
}

func CreateStreamReader(de DocumentEntry, compress bool, fileVersion fileheader.FileVersion) (*StreamReader, error) {
	sr := &StreamReader{fileVersion: fileVersion}

	if !compress {
		sr.is = de.NewDocumentInputStream()
		sr.size = de.GetSize()
	} else {
		decompressed, err := getDecompressedBytes(de.NewDocumentInputStream())
		if err != nil {
			sr.is = de.NewDocumentInputStream()
			sr.size = de.GetSize()
		} else {
			sr.is = ioutil.NopCloser(bytes.NewReader(decompressed))
			sr.size = int64(len(decompressed))
		}
	}
	return sr, nil
}

func getDecompressedBytes(is io.Reader) ([]byte, error) {
	iis, err := zlib.NewReader(is)
	if err != nil {
		return nil, err
	}
	defer iis.Close()

	return ioutil.ReadAll(iis)
}

func (sr *StreamReader) ReadBytes(buffer []byte) error {
	_, err := sr.is.Read(buffer)
	if err != nil {
		return err
	}
	sr.forwardPosition(int64(len(buffer)))
	return nil
}

func (sr *StreamReader) forwardPosition(n int64) {
	sr.read += n
	sr.readAfterHeader += n
}

func (sr *StreamReader) ReadSInt1() (byte, error) {
	buffer := make([]byte, 1)
	err := sr.ReadBytes(buffer)
	return buffer[0], err
}

func (sr *StreamReader) ReadUInt1() (uint16, error) {
	buffer := make([]byte, 1)
	err := sr.ReadBytes(buffer)
	return uint16(buffer[0]), err
}

func (sr *StreamReader) ReadSInt2() (int16, error) {
	buffer := make([]byte, 2)
	err := sr.ReadBytes(buffer)
	return int16(binary.LittleEndian.Uint16(buffer)), err
}

func (sr *StreamReader) ReadUInt2() (uint16, error) {
	buffer := make([]byte, 2)
	err := sr.ReadBytes(buffer)
	return binary.LittleEndian.Uint16(buffer), err
}

func (sr *StreamReader) ReadSInt4() (int32, error) {
	buffer := make([]byte, 4)
	err := sr.ReadBytes(buffer)
	return int32(binary.LittleEndian.Uint32(buffer)), err
}

func (sr *StreamReader) ReadUInt4() (uint32, error) {
	buffer := make([]byte, 4)
	err := sr.ReadBytes(buffer)
	return binary.LittleEndian.Uint32(buffer), err
}

func (sr *StreamReader) ReadDouble() (float64, error) {
	buffer := make([]byte, 8)
	err := sr.ReadBytes(buffer)
	return float64(binary.LittleEndian.Uint64(buffer)), err
}

func (sr *StreamReader) ReadFloat() (float32, error) {
	buffer := make([]byte, 4)
	err := sr.ReadBytes(buffer)
	return float32(binary.LittleEndian.Uint32(buffer)), err
}

func (sr *StreamReader) Skip(n int) error {
	buffer := make([]byte, n)
	return sr.ReadBytes(buffer)
}

func (sr *StreamReader) Close() error {
	return sr.is.Close()
}

// ... 추가 메서드 변환 ...
