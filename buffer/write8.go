package buffer

import (
	"encoding/binary"
	"io"
)

func WriteInt8(writer io.Writer, v int8) error {
	return binary.Write(writer, binary.BigEndian, v)
}

func WriteUint8(writer io.Writer, v uint8) error {
	return binary.Write(writer, binary.BigEndian, v)
}

// alias for `WriteUint8`
//
// concept from Java
func WriteByte(writer io.Writer, v uint8) error {
	return WriteUint8(writer, v)
}

// shortcut for true `WriteUint8(1)` and false `WriteUint(0)`
//
// concept from Java
func WriteBool(writer io.Writer, v bool) error {
	b := uint8(0)
	if v {
		b = 1
	}
	return WriteUint8(writer, b)
}
