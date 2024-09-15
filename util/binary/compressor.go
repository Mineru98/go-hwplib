package binary

import (
	"bytes"
	"compress/flate"
	"encoding/binary"
	"io"
)

// Compress compresses the input data using DEFLATE algorithm.
func Compress(original []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer, err := flate.NewWriter(&buf, flate.DefaultCompression)
	if err != nil {
		return nil, err
	}
	_, err = writer.Write(original)
	if err != nil {
		return nil, err
	}
	writer.Close()

	// Append zero bytes
	buf.Write([]byte{0, 0, 0, 0})

	// Append original length in little-endian
	length := make([]byte, 4)
	binary.LittleEndian.PutUint32(length, uint32(len(original)))
	buf.Write(length)

	return buf.Bytes(), nil
}

// Decompress decompresses the input data using DEFLATE algorithm.
func Decompress(compressed []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(compressed))
	defer reader.Close()

	var buf bytes.Buffer
	_, err := io.Copy(&buf, reader)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
