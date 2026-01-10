package buffer

import (
	"encoding/binary"
	"io"
)

func WriteInt16(writer io.Writer, endian binary.ByteOrder, v int16) error {
	return binary.Write(writer, endian, v)
}

func WriteUint16(writer io.Writer, endian binary.ByteOrder, v uint16) error {
	return binary.Write(writer, endian, v)
}

func WriteInt16LE(writer io.Writer, v int16) error {
	return WriteInt16(writer, binary.LittleEndian, v)
}

func WriteUint16LE(writer io.Writer, v uint16) error {
	return WriteUint16(writer, binary.LittleEndian, v)
}

func WriteInt16BE(writer io.Writer, v int16) error {
	return WriteInt16(writer, binary.BigEndian, v)
}

func WriteUint16BE(writer io.Writer, v uint16) error {
	return WriteUint16(writer, binary.BigEndian, v)
}

// alias for `WriteInt16`
//
// concept from Java
func WriteShort(writer io.Writer, endian binary.ByteOrder, v int16) error {
	return WriteInt16(writer, endian, v)
}

// alias for `WriteUint16`
//
// concept from Java
func WriteUShort(writer io.Writer, endian binary.ByteOrder, v uint16) error {
	return WriteUint16(writer, endian, v)
}

func WriteShortLE(writer io.Writer, v int16) error {
	return WriteShort(writer, binary.LittleEndian, v)
}

func WriteUShortLE(writer io.Writer, v uint16) error {
	return WriteUShort(writer, binary.LittleEndian, v)
}

func WriteShortBE(writer io.Writer, v int16) error {
	return WriteShort(writer, binary.BigEndian, v)
}

func WriteUShortBE(writer io.Writer, v uint16) error {
	return WriteUShort(writer, binary.BigEndian, v)
}
