package buffer

import (
	"encoding/binary"
	"io"
)

func ReadInt16(reader io.Reader, endian binary.ByteOrder) (int16, error) {
	result := int16(0)
	err := binary.Read(reader, endian, &result)
	return result, err
}

func ReadUint16(reader io.Reader, endian binary.ByteOrder) (uint16, error) {
	result := uint16(0)
	err := binary.Read(reader, endian, &result)
	return result, err
}

func ReadInt16LE(reader io.Reader) (int16, error) {
	return ReadInt16(reader, binary.LittleEndian)
}

func ReadUint16LE(reader io.Reader) (uint16, error) {
	return ReadUint16(reader, binary.LittleEndian)
}

func ReadInt16BE(reader io.Reader) (int16, error) {
	return ReadInt16(reader, binary.BigEndian)
}

func ReadUint16BE(reader io.Reader) (uint16, error) {
	return ReadUint16(reader, binary.BigEndian)
}

// alias for `ReadInt16`
//
// concept from Java
func ReadShort(reader io.Reader, endian binary.ByteOrder) (int16, error) {
	return ReadInt16(reader, endian)
}

// alias for `ReadUint16`
//
// concept from Java
func ReadUShort(reader io.Reader, endian binary.ByteOrder) (uint16, error) {
	return ReadUint16(reader, endian)
}

func ReadShortLE(reader io.Reader) (int16, error) {
	return ReadShort(reader, binary.LittleEndian)
}

func ReadUShortLE(reader io.Reader) (uint16, error) {
	return ReadUShort(reader, binary.LittleEndian)
}

func ReadShortBE(reader io.Reader) (int16, error) {
	return ReadShort(reader, binary.BigEndian)
}

func ReadUShortBE(reader io.Reader) (uint16, error) {
	return ReadUShort(reader, binary.BigEndian)
}
