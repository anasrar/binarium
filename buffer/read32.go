package buffer

import (
	"encoding/binary"
	"io"
)

func ReadInt32(reader io.Reader, endian binary.ByteOrder) (int32, error) {
	result := int32(0)
	err := binary.Read(reader, endian, &result)
	return result, err
}

func ReadUint32(reader io.Reader, endian binary.ByteOrder) (uint32, error) {
	result := uint32(0)
	err := binary.Read(reader, endian, &result)
	return result, err
}

func ReadInt32LE(reader io.Reader) (int32, error) {
	return ReadInt32(reader, binary.LittleEndian)
}

func ReadUint32LE(reader io.Reader) (uint32, error) {
	return ReadUint32(reader, binary.LittleEndian)
}

func ReadInt32BE(reader io.Reader) (int32, error) {
	return ReadInt32(reader, binary.BigEndian)
}

func ReadUint32BE(reader io.Reader) (uint32, error) {
	return ReadUint32(reader, binary.BigEndian)
}

// alias for `ReadInt32`
//
// concept from Java
func ReadInt(reader io.Reader, endian binary.ByteOrder) (int32, error) {
	return ReadInt32(reader, endian)
}

// alias for `ReadUint32`
//
// concept from Java
func ReadUInt(reader io.Reader, endian binary.ByteOrder) (uint32, error) {
	return ReadUint32(reader, endian)
}

func ReadIntLE(reader io.Reader) (int32, error) {
	return ReadInt(reader, binary.LittleEndian)
}

func ReadUIntLE(reader io.Reader) (uint32, error) {
	return ReadUInt(reader, binary.LittleEndian)
}

func ReadIntBE(reader io.Reader) (int32, error) {
	return ReadInt(reader, binary.BigEndian)
}

func ReadUIntBE(reader io.Reader) (uint32, error) {
	return ReadUInt(reader, binary.BigEndian)
}

func ReadFloat32(reader io.Reader, endian binary.ByteOrder) (float32, error) {
	result := float32(0)
	err := binary.Read(reader, endian, &result)
	return result, err
}

func ReadFloat32LE(reader io.Reader, endian binary.ByteOrder) (float32, error) {
	return ReadFloat32(reader, binary.LittleEndian)
}

func ReadFloat32BE(reader io.Reader, endian binary.ByteOrder) (float32, error) {
	return ReadFloat32(reader, binary.BigEndian)
}

// alias for `ReadFloat32`
//
// concept from Java
func ReadFloat(reader io.Reader, endian binary.ByteOrder) (float32, error) {
	return ReadFloat32(reader, endian)
}

func ReadFloatLE(reader io.Reader, endian binary.ByteOrder) (float32, error) {
	return ReadFloat(reader, binary.LittleEndian)
}

func ReadFloatBE(reader io.Reader, endian binary.ByteOrder) (float32, error) {
	return ReadFloat(reader, binary.BigEndian)
}
