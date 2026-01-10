package buffer

import (
	"encoding/binary"
	"io"
)

func WriteInt32(writer io.Writer, endian binary.ByteOrder, v int32) error {
	return binary.Write(writer, endian, v)
}

func WriteUint32(writer io.Writer, endian binary.ByteOrder, v uint32) error {
	return binary.Write(writer, endian, v)
}

func WriteInt32LE(writer io.Writer, v int32) error {
	return WriteInt32(writer, binary.LittleEndian, v)
}

func WriteUint32LE(writer io.Writer, v uint32) error {
	return WriteUint32(writer, binary.LittleEndian, v)
}

func WriteInt32BE(writer io.Writer, v int32) error {
	return WriteInt32(writer, binary.BigEndian, v)
}

func WriteUint32BE(writer io.Writer, v uint32) error {
	return WriteUint32(writer, binary.BigEndian, v)
}

// alias for `WriteInt32`
//
// concept from Java
func WriteInt(writer io.Writer, endian binary.ByteOrder, v int32) error {
	return WriteInt32(writer, endian, v)
}

// alias for `WriteUInt32`
//
// concept from Java
func WriteUInt(writer io.Writer, endian binary.ByteOrder, v uint32) error {
	return WriteUint32(writer, endian, v)
}

func WriteIntLE(writer io.Writer, v int32) error {
	return WriteInt(writer, binary.LittleEndian, v)
}

func WriteUIntLE(writer io.Writer, v uint32) error {
	return WriteUInt(writer, binary.LittleEndian, v)
}

func WriteIntBE(writer io.Writer, v int32) error {
	return WriteInt(writer, binary.BigEndian, v)
}

func WriteUIntBE(writer io.Writer, v uint32) error {
	return WriteUInt(writer, binary.BigEndian, v)
}

func WriteFloat32(writer io.Writer, endian binary.ByteOrder, v float32) error {
	return binary.Write(writer, endian, v)
}

func WriteFloat32LE(writer io.Writer, endian binary.ByteOrder, v float32) error {
	return WriteFloat32(writer, binary.LittleEndian, v)
}

func WriteFloat32BE(writer io.Writer, endian binary.ByteOrder, v float32) error {
	return WriteFloat32(writer, binary.BigEndian, v)
}

// alias for `WriteFloat32`
//
// concept from Java
func WriteFloat(writer io.Writer, endian binary.ByteOrder, v float32) error {
	return WriteFloat32(writer, endian, v)
}

func WriteFloatLE(writer io.Writer, endian binary.ByteOrder, v float32) error {
	return WriteFloat(writer, binary.LittleEndian, v)
}

func WriteFloatBE(writer io.Writer, endian binary.ByteOrder, v float32) error {
	return WriteFloat(writer, binary.BigEndian, v)
}
