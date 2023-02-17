package gblob

import "io"

// TypedReader represents a reader that can parse specific Go types from
// a byte sequence.
//
// The endianness depends on the actual implementation.
type TypedReader interface {

	// ReadUint8 reads a single uint8 from the source.
	ReadUint8() (uint8, error)

	// ReadInt8 reads a single int8 from the source.
	ReadInt8() (int8, error)

	// ReadUint16 reads a single uint16 from the source.
	ReadUint16() (uint16, error)

	// ReadInt16 reads a single int16 from the source.
	ReadInt16() (int16, error)

	// ReadUint32 reads a single uint32 from the source.
	ReadUint32() (uint32, error)

	// ReadInt32 reads a single int32 from the source.
	ReadInt32() (int32, error)

	// ReadUint64 reads a single uint64 from the source.
	ReadUint64() (uint64, error)

	// ReadInt64 reads a single int64 from the source.
	ReadInt64() (int64, error)

	// ReadFloat32 reads a single float32 from the source.
	ReadFloat32() (float32, error)

	// ReadFloat64 reads a single float64 from the source.
	ReadFloat64() (float64, error)

	// ReadBytes reads exactly len(target) bytes from the source and places
	// them inside target.
	ReadBytes(target []byte) error
}

// NewLittleEndianReader returns an implementation of TypedReader that reads
// from the specified in Reader in Little Endian order.
func NewLittleEndianReader(in io.Reader) TypedReader {
	return &typedReader[LittleEndianBlock]{
		in:     in,
		buffer: make(LittleEndianBlock, 8), // 64 bit max
	}
}

// NewBigEndianReader returns an implementation of TypedReader that reads
// from the specified in Reader in Big Endian order.
func NewBigEndianReader(in io.Reader) TypedReader {
	return &typedReader[BigEndianBlock]{
		in:     in,
		buffer: make(BigEndianBlock, 8), // 64 bit max
	}
}

type typedReader[T blockBuffer] struct {
	in     io.Reader
	buffer T
}

func (r *typedReader[T]) ReadUint8() (uint8, error) {
	err := r.fillBuffer(1)
	return r.buffer.Uint8(0), err
}

func (r *typedReader[T]) ReadInt8() (int8, error) {
	err := r.fillBuffer(1)
	return r.buffer.Int8(0), err
}

func (r *typedReader[T]) ReadUint16() (uint16, error) {
	err := r.fillBuffer(2)
	return r.buffer.Uint16(0), err
}

func (r *typedReader[T]) ReadInt16() (int16, error) {
	err := r.fillBuffer(2)
	return r.buffer.Int16(0), err
}

func (r *typedReader[T]) ReadUint32() (uint32, error) {
	err := r.fillBuffer(4)
	return r.buffer.Uint32(0), err
}

func (r *typedReader[T]) ReadInt32() (int32, error) {
	err := r.fillBuffer(4)
	return r.buffer.Int32(0), err
}

func (r *typedReader[T]) ReadUint64() (uint64, error) {
	err := r.fillBuffer(8)
	return r.buffer.Uint64(0), err
}

func (r *typedReader[T]) ReadInt64() (int64, error) {
	err := r.fillBuffer(8)
	return r.buffer.Int64(0), err
}

func (r *typedReader[T]) ReadFloat32() (float32, error) {
	err := r.fillBuffer(4)
	return r.buffer.Float32(0), err
}

func (r *typedReader[T]) ReadFloat64() (float64, error) {
	err := r.fillBuffer(8)
	return r.buffer.Float64(0), err
}

func (r *typedReader[T]) ReadBytes(target []byte) error {
	_, err := io.ReadFull(r.in, target)
	return err
}

func (r *typedReader[T]) fillBuffer(count int) error {
	return r.ReadBytes(r.buffer[:count])
}
