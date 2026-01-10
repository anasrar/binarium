package buffer

import (
	"encoding/binary"
	"io"
)

func WriteInt64(writer io.Writer, endian binary.ByteOrder, v int64) error {
	return binary.Write(writer, endian, v)
}

func WriteUInt64(writer io.Writer, endian binary.ByteOrder, v uint64) error {
	return binary.Write(writer, endian, v)
}

func WriteInt64LE(writer io.Writer, v int64) error {
	return WriteInt64(writer, binary.LittleEndian, v)
}

func WriteUint64LE(writer io.Writer, v uint64) error {
	return WriteUInt64(writer, binary.LittleEndian, v)
}

func WriteInt64BE(writer io.Writer, v int64) error {
	return WriteInt64(writer, binary.BigEndian, v)
}

func WriteUInt64BE(writer io.Writer, v uint64) error {
	return WriteUInt64(writer, binary.BigEndian, v)
}

// alias for `WriteInt64`
//
// concept from Java
func WriteLong(writer io.Writer, endian binary.ByteOrder, v int64) error {
	return WriteInt64(writer, endian, v)
}

// alias for `WriteUint64`
//
// concept from Java
func WriteULong(writer io.Writer, endian binary.ByteOrder, v uint64) error {
	return WriteUInt64(writer, endian, v)
}

func WriteLongLE(writer io.Writer, v int64) error {
	return WriteInt64(writer, binary.LittleEndian, v)
}

func WriteULongLE(writer io.Writer, v uint64) error {
	return WriteUInt64(writer, binary.LittleEndian, v)
}

func WriteLongBE(writer io.Writer, v int64) error {
	return WriteInt64(writer, binary.BigEndian, v)
}

func WriteULongBE(writer io.Writer, v uint64) error {
	return WriteUInt64(writer, binary.BigEndian, v)
}

func WriteFloat64(writer io.Writer, endian binary.ByteOrder, v float64) error {
	return binary.Write(writer, endian, v)
}

func WriteFloat64LE(writer io.Writer, endian binary.ByteOrder, v float64) error {
	return WriteFloat64(writer, binary.LittleEndian, v)
}

func WriteFloat64BE(writer io.Writer, endian binary.ByteOrder, v float64) error {
	return WriteFloat64(writer, binary.BigEndian, v)
}

// alias for `WriteFloat64`
//
// concept from Java
func WriteDouble(writer io.Writer, endian binary.ByteOrder, v float64) error {
	return WriteFloat64(writer, endian, v)
}

func WriteDoubleLE(writer io.Writer, endian binary.ByteOrder, v float64) error {
	return WriteDouble(writer, binary.LittleEndian, v)
}

func WriteDoubleBE(writer io.Writer, endian binary.ByteOrder, v float64) error {
	return WriteDouble(writer, binary.BigEndian, v)
}
