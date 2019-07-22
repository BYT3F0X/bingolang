// streamio this package contains utilities for dealing with binary streams in go
package streamio

import (
	"errors"
	"io"
)

// readBytesPrimitive the primitive function used by this library to read streams in chunks which may be offset.
// returns bytes read, number of bytes read and any errors that may have occurred.
func readBytesPrimitive(reader io.Reader, length int, offset int) ([]byte, int, error) {
	bytes := make([]byte, length)

	if offset != 0 {
		switch r := reader.(type) {
		case io.Seeker:
			_, err := r.Seek(0, offset)
			if err != nil {
				return bytes, -1, err
			}
		default:
			return bytes, -1, errors.New("the stream is not seekable")
		}
	}

	// From this point on if a error is returned the bytes array may not be empty
	retNum, err := reader.Read(bytes)

	if err != nil {
		return bytes, retNum, err
	}

	return bytes, retNum, nil
}

func ReadBytes(reader io.Reader, size int) ([]byte, int, error) {
	return readBytesPrimitive(reader, size, 0)
}

func ReadBytesOffset(reader io.Reader, size int, offset int) ([]byte, int, error) {
	return readBytesPrimitive(reader, size, offset)
}

