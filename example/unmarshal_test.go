package example

import (
	"encoding/binary"
	"testing"

	"github.com/anasrar/binarium"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalExampleString(t *testing.T) {
	str := ExampleUnmarshalString{}
	err := binarium.Unmarshal(
		[]byte{0, 5, 104, 101, 108, 108, 111},
		binary.BigEndian,
		&str,
	)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		ExampleUnmarshalString{
			StringLength: 5,
			String:       "hello",
		},
		str,
	)
}

func TestUnmarshalExampleBytes(t *testing.T) {
	s := ExampleUnmarshalBytes{}
	err := binarium.Unmarshal(
		[]byte{0, 5, 104, 101, 108, 108, 111},
		binary.BigEndian,
		&s,
	)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		ExampleUnmarshalBytes{
			ByteLength: 5,
			Bytes:      []byte{104, 101, 108, 108, 111},
		},
		s,
	)
}

func TestUnmarshalExampleSliceFloat32(t *testing.T) {
	s := ExampleUnmarshalSliceFloat32{}
	err := binarium.Unmarshal(
		[]byte{3, 63, 128, 0, 0, 64, 0, 0, 0, 64, 64, 0, 0},
		binary.BigEndian,
		&s,
	)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		ExampleUnmarshalSliceFloat32{
			Float32Length: 3,
			Float32s:      []float32{1, 2, 3},
		},
		s,
	)
}

func TestUnmarshalExampleSliceStruct(t *testing.T) {
	s := ExampleUnmarshalSliceStruct{}
	err := binarium.Unmarshal(
		[]byte{
			3,

			63, 128, 0, 0,
			1,

			64, 0, 0, 0,
			2,

			64, 64, 0, 0,
			3,
		},
		binary.BigEndian,
		&s,
	)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		ExampleUnmarshalSliceStruct{
			StructLength: 3,
			Structs: []ExampleSliceStruct{
				{
					Float32: 1,
					Uint8:   1,
				},
				{
					Float32: 2,
					Uint8:   2,
				},
				{
					Float32: 3,
					Uint8:   3,
				},
			},
		},
		s,
	)
}
