package buffer

import (
	"encoding/binary"
	"io"
)

func ReadInt64(reader io.Reader, endian binary.ByteOrder) (int64, error) {
	result := int64(0)
	err := binary.Read(reader, endian, &result)
	return result, err
}

func ReadUint64(reader io.Reader, endian binary.ByteOrder) (uint64, error) {
	result := uint64(0)
	err := binary.Read(reader, endian, &result)
	return result, err
}

func ReadInt64LE(reader io.Reader) (int64, error) {
	return ReadInt64(reader, binary.LittleEndian)
}

func ReadUint64LE(reader io.Reader) (uint64, error) {
	return ReadUint64(reader, binary.LittleEndian)
}

func ReadInt64BE(reader io.Reader) (int64, error) {
	return ReadInt64(reader, binary.BigEndian)
}

func ReadUint64BE(reader io.Reader) (uint64, error) {
	return ReadUint64(reader, binary.BigEndian)
}

// alias for `ReadInt64`
//
// concept from Java
func ReadLong(reader io.Reader, endian binary.ByteOrder) (int64, error) {
	return ReadInt64(reader, endian)
}

// alias for `ReadUint64`
//
// concept from Java
func ReadULong(reader io.Reader, endian binary.ByteOrder) (uint64, error) {
	return ReadUint64(reader, endian)
}

func ReadLongLE(reader io.Reader) (int64, error) {
	return ReadLong(reader, binary.LittleEndian)
}

func ReadULongLE(reader io.Reader) (uint64, error) {
	return ReadULong(reader, binary.LittleEndian)
}

func ReadLongBE(reader io.Reader) (int64, error) {
	return ReadLong(reader, binary.BigEndian)
}

func ReadULongBE(reader io.Reader) (uint64, error) {
	return ReadULong(reader, binary.BigEndian)
}

func ReadFloat64(reader io.Reader, endian binary.ByteOrder) (float64, error) {
	result := float64(0)
	err := binary.Read(reader, endian, &result)
	return result, err
}

func ReadFloat64LE(reader io.Reader, endian binary.ByteOrder) (float64, error) {
	return ReadFloat64(reader, binary.LittleEndian)
}

func ReadFloat64BE(reader io.Reader, endian binary.ByteOrder) (float64, error) {
	return ReadFloat64(reader, binary.BigEndian)
}

// alias for `ReadFloat64`
//
// concept from Java
func ReadDouble(reader io.Reader, endian binary.ByteOrder) (float64, error) {
	return ReadFloat64(reader, endian)
}

func ReadDoubleLE(reader io.Reader, endian binary.ByteOrder) (float64, error) {
	return ReadDouble(reader, binary.LittleEndian)
}

func ReadDoubleBE(reader io.Reader, endian binary.ByteOrder) (float64, error) {
	return ReadDouble(reader, binary.BigEndian)
}
