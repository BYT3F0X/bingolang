package util

import (
	"errors"
	"unsafe"
)

type Endianness int

const (
	BigEndian    Endianness = 0
	LittleEndian Endianness = 1
	UndefinedEndianness = 100
)

func DetermineEndianness() (Endianness, error) {
	i := 0x0100
	ptr := unsafe.Pointer(&i)

	if 0x01 == *(*byte)(ptr) {
		return BigEndian, nil
	} else if 0x00 == *(*byte)(ptr) {
		return LittleEndian, nil
	} else {
		return UndefinedEndianness, errors.New("unable to determine endianness of the host")
	}
}
