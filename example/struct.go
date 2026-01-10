package example

type ExampleNestedStruct struct {
	Int32 int32
	Uint8 uint8
}

type ExampleSliceStruct struct {
	Float32 float32
	Uint8   uint8
}

type ExampleMarshalStruct struct {
	SliceLength  uint8
	ByteLength   uint8
	StringLength uint16
	SkipField    string               `skip:""`
	String       string               `length:"StringLength"`
	Byte         []byte               `length:"ByteLength"`
	Nested       ExampleNestedStruct  ``
	Slice        []ExampleSliceStruct `length:"SliceLength"`
}

type ExampleUnmarshalString struct {
	SkipField    string `skip:""`
	StringLength uint16
	String       string `length:"StringLength"`
}

type ExampleUnmarshalBytes struct {
	ByteLength uint16
	SkipField  string `skip:""`
	Bytes      []byte `length:"ByteLength"`
}

type ExampleUnmarshalSliceFloat32 struct {
	Float32Length uint8
	Float32s      []float32 `length:"Float32Length"`
}

type ExampleUnmarshalSliceStruct struct {
	StructLength uint8
	Structs      []ExampleSliceStruct `length:"StructLength"`
}
