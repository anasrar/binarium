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

func Unmarshal(b []byte, endian binary.ByteOrder, s any) error {
	buff := bytes.NewBuffer(b)
	reader := bufio.NewReader(buff)

	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Pointer || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("s must be pointer to struct")
	}

	v = v.Elem()

	if err := unmarshalstruct(reader, endian, v); err != nil {
		return err
	}

	return nil
}

func unmarshalstruct(reader io.Reader, endian binary.ByteOrder, v reflect.Value) error {

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldType := t.Field(i)
		fieldValue := v.Field(i)

		if _, skip := fieldType.Tag.Lookup("skip"); skip {
			continue
		}

		if err := readkind(reader, endian, v, t, fieldType, fieldValue); err != nil {
			return err
		}
	}

	return nil
}

func readkind(reader io.Reader, endian binary.ByteOrder, structValue reflect.Value, structType reflect.Type, field reflect.StructField, value reflect.Value) error {
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return fmt.Errorf("field %s is nil pointer", field.Name)
		}
		return readkind(reader, endian, structValue, structType, field, value.Elem())
	}

	switch value.Kind() {

	case reflect.Struct:
		return unmarshalstruct(reader, endian, value)

	case reflect.Slice:
		key, ok := field.Tag.Lookup("length")
		if !ok {
			return fmt.Errorf(
				"field %s (%s) required tag length",
				field.Name,
				value.Kind(),
			)
		}

		_, ok = structType.FieldByName(key)
		if !ok {
			return fmt.Errorf(
				"field %s (%s) tag length %s field not found",
				field.Name,
				value.Kind(),
				key,
			)
		}
		lengthvalue := structValue.FieldByName(key)

		if !lengthvalue.CanUint() && !lengthvalue.CanInt() {
			return fmt.Errorf("field %s is not uint or int", key)
		}

		if value.Type().Elem().Kind() == reflect.Uint8 {
			b := make([]byte, lengthvalue.Uint())
			if _, err := reader.Read(b); err != nil {
				return err
			}
			value.SetBytes(b)
			return nil
		}

		length := lengthvalue.Uint()

		slice := reflect.MakeSlice(value.Type(), int(length), int(length))
		value.Set(slice)

		for i := 0; i < int(length); i++ {
			elem := value.Index(i)

			elemField := field
			elemField.Name = fmt.Sprintf("%s[%d]", field.Name, i)

			if err := readkind(reader, endian, structValue, structType, elemField, elem); err != nil {
				return err
			}
		}

		return nil

	case reflect.String:
		key, ok := field.Tag.Lookup("length")
		if !ok {
			return fmt.Errorf(
				"field %s (%s) required tag length",
				field.Name,
				value.Kind(),
			)
		}

		_, ok = structType.FieldByName(key)
		if !ok {
			return fmt.Errorf(
				"field %s (%s) tag length %s field not found",
				field.Name,
				value.Kind(),
				key,
			)
		}
		lengthvalue := structValue.FieldByName(key)

		if !lengthvalue.CanUint() && !lengthvalue.CanInt() {
			return fmt.Errorf("field %s is not uint or int", key)
		}

		b := make([]byte, lengthvalue.Uint())
		if _, err := reader.Read(b); err != nil {
			return err
		}

		value.SetString(string(b))

		return nil

	case reflect.Bool:
		v, err := buffer.ReadBool(reader)
		value.SetBool(v)
		return err

	case reflect.Int8:
		v, err := buffer.ReadInt8(reader)
		value.SetInt(int64(v))
		return err

	case reflect.Uint8:
		v, err := buffer.ReadUint8(reader)
		value.SetUint(uint64(v))
		return err

	case reflect.Int16:
		v, err := buffer.ReadInt16(reader, endian)
		value.SetInt(int64(v))
		return err

	case reflect.Uint16:
		v, err := buffer.ReadUint16(reader, endian)
		value.SetUint(uint64(v))
		return err

	case reflect.Int32:
		v, err := buffer.ReadInt32(reader, endian)
		value.SetInt(int64(v))
		return err

	case reflect.Uint32:
		v, err := buffer.ReadUint32(reader, endian)
		value.SetUint(uint64(v))
		return err

	case reflect.Int64:
		v, err := buffer.ReadInt64(reader, endian)
		value.SetInt(int64(v))
		return err

	case reflect.Uint64:
		v, err := buffer.ReadUint64(reader, endian)
		value.SetUint(uint64(v))
		return err

	case reflect.Float32:
		v, err := buffer.ReadFloat32(reader, endian)
		value.SetFloat(float64(v))
		return err

	case reflect.Float64:
		v, err := buffer.ReadFloat64(reader, endian)
		value.SetFloat(v)
		return err

	default:
		return fmt.Errorf(
			"field %s (%s) is not supported for unmarshal",
			field.Name,
			value.Kind(),
		)
	}
}
