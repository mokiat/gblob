package gblob

import (
	"fmt"
	"io"
	"reflect"
)

var (
	decodableType = reflect.TypeFor[PackedDecodable]()
)

// PackedDecodable is an interface that can be implemented by types that want to
// provide a custom packed decoding.
type PackedDecodable interface {

	// DecodePacked decodes the receiver from the specified reader.
	DecodePacked(reader TypedReader) error
}

// NewLittleEndianPackedDecoder creates a new PackedDecoder that is configured
// to read its input in Little Endian order.
func NewLittleEndianPackedDecoder(in io.Reader) *PackedDecoder {
	return &PackedDecoder{
		in: NewLittleEndianReader(in),
	}
}

// NewBigEndianPackedDecoder creates a new PackedDecoder that is configured
// to read its input in Big Endian order.
func NewBigEndianPackedDecoder(in io.Reader) *PackedDecoder {
	return &PackedDecoder{
		in: NewBigEndianReader(in),
	}
}

// PackedDecoder decodes arbitrary Go objects from binary form by going through
// each field in sequence and deserializing it without any padding.
type PackedDecoder struct {
	in TypedReader
}

// Decode decodes the specified target value from the Reader.
func (d *PackedDecoder) Decode(target interface{}) error {
	value := reflect.ValueOf(target)
	return d.decodeValue(value)
}

func (d *PackedDecoder) decodeValue(value reflect.Value) error {
	if value.Type().Implements(decodableType) {
		if value.Kind() == reflect.Pointer && value.IsNil() {
			value.Set(reflect.New(value.Type().Elem()))
		}
		decodable := value.Interface().(PackedDecodable)
		return decodable.DecodePacked(d.in)
	}
	switch kind := value.Kind(); kind {
	case reflect.Pointer:
		if value.IsNil() {
			value.Set(reflect.New(value.Type().Elem()))
		}
		return d.decodeValue(value.Elem())
	case reflect.Bool:
		v, err := d.in.ReadUint8()
		if err != nil {
			return err
		}
		value.SetBool(v > 0x00)
		return nil
	case reflect.Uint8:
		v, err := d.in.ReadUint8()
		if err != nil {
			return err
		}
		value.SetUint(uint64(v))
		return nil
	case reflect.Int8:
		v, err := d.in.ReadInt8()
		if err != nil {
			return err
		}
		value.SetInt(int64(v))
		return nil
	case reflect.Uint16:
		v, err := d.in.ReadUint16()
		if err != nil {
			return err
		}
		value.SetUint(uint64(v))
		return nil
	case reflect.Int16:
		v, err := d.in.ReadInt16()
		if err != nil {
			return err
		}
		value.SetInt(int64(v))
		return nil
	case reflect.Uint32:
		v, err := d.in.ReadUint32()
		if err != nil {
			return err
		}
		value.SetUint(uint64(v))
		return nil
	case reflect.Int32:
		v, err := d.in.ReadInt32()
		if err != nil {
			return err
		}
		value.SetInt(int64(v))
		return nil
	case reflect.Uint64:
		v, err := d.in.ReadUint64()
		if err != nil {
			return err
		}
		value.SetUint(uint64(v))
		return nil
	case reflect.Int64:
		v, err := d.in.ReadInt64()
		if err != nil {
			return err
		}
		value.SetInt(int64(v))
		return nil
	case reflect.Float32:
		v, err := d.in.ReadFloat32()
		if err != nil {
			return err
		}
		value.SetFloat(float64(v))
		return nil
	case reflect.Float64:
		v, err := d.in.ReadFloat64()
		if err != nil {
			return err
		}
		value.SetFloat(float64(v))
		return nil
	case reflect.Array:
		count := value.Len()
		for i := 0; i < count; i++ {
			if err := d.decodeValue(value.Index(i)); err != nil {
				return err
			}
		}
		return nil
	case reflect.Struct:
		fieldCount := value.NumField()
		for i := 0; i < fieldCount; i++ {
			field := value.Field(i)
			if err := d.decodeValue(field); err != nil {
				return err
			}
		}
		return nil
	case reflect.Slice:
		count, err := d.in.ReadUint64()
		if err != nil {
			return err
		}
		if value.Type().Elem().Kind() == reflect.Uint8 { // fast track
			data := make([]uint8, count)
			if err := d.in.ReadBytes(data); err != nil {
				return err
			}
			value.Set(reflect.ValueOf(data).Convert(value.Type()))
		} else {
			value.Set(reflect.MakeSlice(value.Type(), int(count), int(count)))
			for i := 0; i < int(count); i++ {
				if err := d.decodeValue(value.Index(i)); err != nil {
					return err
				}
			}
		}
		return nil
	case reflect.Map:
		count, err := d.in.ReadUint64()
		if err != nil {
			return err
		}
		value.Set(reflect.MakeMapWithSize(value.Type(), int(count)))
		for i := 0; i < int(count); i++ {
			entryKey := reflect.New(value.Type().Key())
			if err := d.decodeValue(entryKey); err != nil {
				return err
			}
			entryValue := reflect.New(value.Type().Elem())
			if err := d.decodeValue(entryValue); err != nil {
				return err
			}
			value.SetMapIndex(entryKey.Elem(), entryValue.Elem())
		}
		return nil
	case reflect.String:
		count, err := d.in.ReadUint64()
		if err != nil {
			return err
		}
		data := make([]byte, count)
		if err := d.in.ReadBytes(data); err != nil {
			return err
		}
		value.SetString(string(data))
		return nil
	default:
		return fmt.Errorf("unsupported type: %v", kind)
	}
}
