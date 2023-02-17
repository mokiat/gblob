package gblob

import "io"

// TypedWriter represents a writer that can serialize specific Go types to
// a byte sequence.
//
// The endianness depends on the actual implementation.
type TypedWriter interface {

	// WriteUint8 writes a single uint8 to the target.
	WriteUint8(uint8) error

	// WriteInt8 writes a single int8 to the target.
	WriteInt8(int8) error

	// WriteUint16 writes a single uint16 to the target.
	WriteUint16(uint16) error

	// WriteInt16 writes a single int16 to the target.
	WriteInt16(int16) error

	// WriteUint32 writes a single uint32 to the target.
	WriteUint32(uint32) error

	// WriteInt32 writes a single int32 to the target.
	WriteInt32(int32) error

	// WriteUint64 writes a single uint64 to the target.
	WriteUint64(uint64) error

	// WriteInt64 writes a single int64 to the target.
	WriteInt64(int64) error

	// WriteFloat32 writes a single float32 to the target.
	WriteFloat32(float32) error

	// WriteFloat64 writes a single float64 to the target.
	WriteFloat64(float64) error

	// WriteBytes writes len(bytes) from source to the target.
	WriteBytes(source []byte) error
}

// NewLittleEndianWriter returns an implementation of TypedWriter that writes
// to the specified out Writer in Little Endian order.
func NewLittleEndianWriter(out io.Writer) TypedWriter {
	return &typedWriter[LittleEndianBlock]{
		out:    out,
		buffer: make(LittleEndianBlock, 8), // 64 bit max
	}
}

// NewBigEndianWriter returns an implementation of TypedWriter that writes
// to the specified out Writer in Big Endian order.
func NewBigEndianWriter(out io.Writer) TypedWriter {
	return &typedWriter[BigEndianBlock]{
		out:    out,
		buffer: make(BigEndianBlock, 8), // 64 bit max
	}
}

type typedWriter[T blockBuffer] struct {
	out    io.Writer
	buffer T
}

func (w *typedWriter[T]) WriteUint8(value uint8) error {
	w.buffer.SetUint8(0, value)
	return w.flushBuffer(1)
}

// WriteInt8 writes a single int8 to the target.
func (w *typedWriter[T]) WriteInt8(value int8) error {
	w.buffer.SetInt8(0, value)
	return w.flushBuffer(1)
}

// WriteUint16 writes a single uint16 to the target.
func (w *typedWriter[T]) WriteUint16(value uint16) error {
	w.buffer.SetUint16(0, value)
	return w.flushBuffer(2)
}

// WriteInt16 writes a single int16 to the target.
func (w *typedWriter[T]) WriteInt16(value int16) error {
	w.buffer.SetInt16(0, value)
	return w.flushBuffer(2)
}

// WriteUint32 writes a single uint32 to the target.
func (w *typedWriter[T]) WriteUint32(value uint32) error {
	w.buffer.SetUint32(0, value)
	return w.flushBuffer(4)
}

// WriteInt32 writes a single int32 to the target.
func (w *typedWriter[T]) WriteInt32(value int32) error {
	w.buffer.SetInt32(0, value)
	return w.flushBuffer(4)
}

// WriteUint64 writes a single uint64 to the target.
func (w *typedWriter[T]) WriteUint64(value uint64) error {
	w.buffer.SetUint64(0, value)
	return w.flushBuffer(8)
}

// WriteInt64 writes a single int64 to the target.
func (w *typedWriter[T]) WriteInt64(value int64) error {
	w.buffer.SetInt64(0, value)
	return w.flushBuffer(8)
}

// WriteFloat32 writes a single float32 to the target.
func (w *typedWriter[T]) WriteFloat32(value float32) error {
	w.buffer.SetFloat32(0, value)
	return w.flushBuffer(4)
}

// WriteFloat64 writes a single float64 to the target.
func (w *typedWriter[T]) WriteFloat64(value float64) error {
	w.buffer.SetFloat64(0, value)
	return w.flushBuffer(8)
}

// WriteBytes writes len(bytes) from source to the target.
func (w *typedWriter[T]) WriteBytes(source []byte) error {
	_, err := w.out.Write(source)
	return err
}

func (w *typedWriter[T]) flushBuffer(count int) error {
	return w.WriteBytes(w.buffer[:count])
}
