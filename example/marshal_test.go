package example

import (
	"encoding/binary"
	"testing"

	"github.com/anasrar/binarium"
	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	example := ExampleMarshalStruct{
		SliceLength:  2,
		ByteLength:   3,
		StringLength: 11,
		String:       "hello world",
		Byte:         []byte{10, 10, 100},
		Nested: ExampleNestedStruct{
			Int32: 1000,
			Uint8: 20,
		},
		Slice: []ExampleSliceStruct{
			{
				Float32: 1,
				Uint8:   1,
			},
			{
				Float32: 0,
				Uint8:   2,
			},
		},
	}

	b, err := binarium.Marshal(binary.BigEndian, example)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		b,
		[]byte{
			2, 3, 0, 11, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 10, 10, 100, 0, 0, 3, 232, 20, 63, 128, 0, 0, 1, 0, 0, 0, 0, 2,
		},
	)
}
