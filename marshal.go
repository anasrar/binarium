package binarium

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"

	"github.com/anasrar/binarium/buffer"
)

func Marshal(endian binary.ByteOrder, s any) ([]byte, error) {
	buff := bytes.NewBuffer([]byte{})
	writer := bufio.NewWriter(buff)

	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return []byte{}, fmt.Errorf("s is not struct")
	}

	if err := marshalstruct(writer, endian, v); err != nil {
		return []byte{}, err
	}

	if err := writer.Flush(); err != nil {
		return []byte{}, err
	}

	return buff.Bytes(), nil
}

func marshalstruct(writer io.Writer, endian binary.ByteOrder, v reflect.Value) error {
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldType := t.Field(i)
		fieldValue := v.Field(i)

		if _, skip := fieldType.Tag.Lookup("skip"); skip {
			continue
		}

		if err := writekind(writer, endian, fieldType, fieldValue); err != nil {
			return err
		}
	}
	return nil
}

func writekind(writer io.Writer, endian binary.ByteOrder, field reflect.StructField, value reflect.Value) error {
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return fmt.Errorf("field %s is nil pointer", field.Name)
		}
		return writekind(writer, endian, field, value.Elem())
	}

	switch value.Kind() {

	case reflect.Struct:
		return marshalstruct(writer, endian, value)

	case reflect.Slice:
		if value.Type().Elem().Kind() == reflect.Uint8 {
			b := value.Bytes()
			_, err := writer.Write(b)
			return err
		}

		length := value.Len()
		for i := 0; i < length; i++ {
			elem := value.Index(i)

			elemField := field
			elemField.Name = fmt.Sprintf("%s[%d]", field.Name, i)

			if err := writekind(writer, endian, elemField, elem); err != nil {
				return err
			}
		}
		return nil

	case reflect.Bool:
		return buffer.WriteBool(writer, value.Bool())

	case reflect.Int8:
		return buffer.WriteInt8(writer, int8(value.Int()))

	case reflect.Uint8:
		return buffer.WriteUint8(writer, uint8(value.Uint()))

	case reflect.Int16:
		return buffer.WriteInt16(writer, endian, int16(value.Int()))

	case reflect.Uint16:
		return buffer.WriteUint16(writer, endian, uint16(value.Uint()))

	case reflect.Int32:
		return buffer.WriteInt32(writer, endian, int32(value.Int()))

	case reflect.Uint32:
		return buffer.WriteUint32(writer, endian, uint32(value.Uint()))

	case reflect.Int64:
		return buffer.WriteInt64(writer, endian, value.Int())

	case reflect.Uint64:
		return buffer.WriteUInt64(writer, endian, value.Uint())

	case reflect.Float32:
		return buffer.WriteFloat32(writer, endian, float32(value.Float()))

	case reflect.Float64:
		return buffer.WriteFloat64(writer, endian, value.Float())

	case reflect.String:
		return buffer.WriteString(writer, value.String())

	default:
		return fmt.Errorf(
			"field %s (%s) is not supported for marshal",
			field.Name,
			value.Kind())
	}
}
