package buffer

import (
	"encoding/binary"
	"io"
)

func ReadInt8(reader io.Reader) (int8, error) {
	result := int8(0)
	err := binary.Read(reader, binary.BigEndian, &result)
	return result, err
}

func ReadUint8(reader io.Reader) (uint8, error) {
	result := uint8(0)
	err := binary.Read(reader, binary.BigEndian, &result)
	return result, err
}

// alias for `ReadUint8`
//
// concept from Java
func ReadByte(reader io.Reader) (uint8, error) {
	return ReadUint8(reader)
}

// shortcut for `ReadUint8` true if value is not zero
//
// concept from Java
func ReadBool(reader io.Reader) (bool, error) {
	result, err := ReadUint8(reader)
	if err != nil {
		return false, err
	}

	b := false
	if result != 0 {
		b = true
	}
	return b, nil
}
