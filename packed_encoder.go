package gblob

import (
	"fmt"
	"io"
	"reflect"
)

// NewLittleEndianPackedEncoder creates a new PackedEncoder that is configured
// to write its output in Little Endian order.
func NewLittleEndianPackedEncoder(out io.Writer) *PackedEncoder {
	return &PackedEncoder{
		out: NewLittleEndianWriter(out),
	}
}

// NewBigEndianPackedEncoder creates a new PackedEncoder that is configured
// to write its output in Big Endian order.
func NewBigEndianPackedEncoder(out io.Writer) *PackedEncoder {
	return &PackedEncoder{
		out: NewBigEndianWriter(out),
	}
}

// PackedEncoder encodes arbitrary Go objects in binary form by going through
// each field in sequence and serializing it without any padding.
type PackedEncoder struct {
	out TypedWriter
}

// Encode encodes the specified source value into the Writer.
func (e *PackedEncoder) Encode(source any) error {
	return e.encodeValue(reflect.ValueOf(source))
}

func (e *PackedEncoder) encodeValue(value reflect.Value) error {
	switch kind := value.Kind(); kind {
	case reflect.Pointer:
		return e.encodeValue(value.Elem())
	case reflect.Uint8:
		return e.out.WriteUint8(uint8(value.Uint()))
	case reflect.Int8:
		return e.out.WriteInt8(int8(value.Int()))
	case reflect.Uint16:
		return e.out.WriteUint16(uint16(value.Uint()))
	case reflect.Int16:
		return e.out.WriteInt16(int16(value.Int()))
	case reflect.Uint32:
		return e.out.WriteUint32(uint32(value.Uint()))
	case reflect.Int32:
		return e.out.WriteInt32(int32(value.Int()))
	case reflect.Uint64:
		return e.out.WriteUint64(uint64(value.Uint()))
	case reflect.Int64:
		return e.out.WriteInt64(int64(value.Int()))
	default:
		return fmt.Errorf("unsupported type: %v", kind)
	}
}
