package gblob

import "math"

// LittleEndianBlock represents a fixed-size block of bytes that holds
// values encoded in Little Endian order.
type LittleEndianBlock []byte

// Uint8 returns the uint8 value at the specified offset.
func (b LittleEndianBlock) Uint8(offset int) uint8 {
	return b[offset]
}

// SetUint8 places the uint8 value at the specified offset.
func (b LittleEndianBlock) SetUint8(offset int, value uint8) {
	b[offset] = value
}

// Int8 returns the int8 value at the specified offset.
func (b LittleEndianBlock) Int8(offset int) int8 {
	return int8(b.Uint8(offset))
}

// SetInt8 places the int8 value at the specified offset.
func (b LittleEndianBlock) SetInt8(offset int, value int8) {
	b.SetUint8(offset, uint8(value))
}

// Uint16 returns the uint16 value at the specified offset.
func (b LittleEndianBlock) Uint16(offset int) uint16 {
	return uint16(b[offset+0])<<0 |
		uint16(b[offset+1])<<8
}

// SetUint16 places the uint16 value at the specified offset.
func (b LittleEndianBlock) SetUint16(offset int, value uint16) {
	b[offset+0] = byte(value >> 0)
	b[offset+1] = byte(value >> 8)
}

// Int16 returns the int16 value at the specified offset.
func (b LittleEndianBlock) Int16(offset int) int16 {
	return int16(b.Uint16(offset))
}

// SetInt16 places the int16 value at the specified offset.
func (b LittleEndianBlock) SetInt16(offset int, value int16) {
	b.SetUint16(offset, uint16(value))
}

// Uint32 returns the uint32 value at the specified offset.
func (b LittleEndianBlock) Uint32(offset int) uint32 {
	return uint32(b[offset+0])<<0 |
		uint32(b[offset+1])<<8 |
		uint32(b[offset+2])<<16 |
		uint32(b[offset+3])<<24
}

// SetUint32 places the uint32 value at the specified offset.
func (b LittleEndianBlock) SetUint32(offset int, value uint32) {
	b[offset+0] = byte(value >> 0)
	b[offset+1] = byte(value >> 8)
	b[offset+2] = byte(value >> 16)
	b[offset+3] = byte(value >> 24)
}

// Int32 returns the int32 value at the specified offset.
func (b LittleEndianBlock) Int32(offset int) int32 {
	return int32(b.Uint32(offset))
}

// SetInt32 places the int32 value at the specified offset.
func (b LittleEndianBlock) SetInt32(offset int, value int32) {
	b.SetUint32(offset, uint32(value))
}

// Uint64 returns the uint64 value at the specified offset.
func (b LittleEndianBlock) Uint64(offset int) uint64 {
	return uint64(b[offset+0])<<0 |
		uint64(b[offset+1])<<8 |
		uint64(b[offset+2])<<16 |
		uint64(b[offset+3])<<24 |
		uint64(b[offset+4])<<32 |
		uint64(b[offset+5])<<40 |
		uint64(b[offset+6])<<48 |
		uint64(b[offset+7])<<56
}

// SetUint64 places the uint64 value at the specified offset.
func (b LittleEndianBlock) SetUint64(offset int, value uint64) {
	b[offset+0] = byte(value >> 0)
	b[offset+1] = byte(value >> 8)
	b[offset+2] = byte(value >> 16)
	b[offset+3] = byte(value >> 24)
	b[offset+4] = byte(value >> 32)
	b[offset+5] = byte(value >> 40)
	b[offset+6] = byte(value >> 48)
	b[offset+7] = byte(value >> 56)
}

// Int64 returns the int64 value at the specified offset.
func (b LittleEndianBlock) Int64(offset int) int64 {
	return int64(b.Uint64(offset))
}

// SetInt64 places the int64 value at the specified offset.
func (b LittleEndianBlock) SetInt64(offset int, value uint64) {
	b.SetUint64(offset, uint64(value))
}

// Float32 returns the float32 value at the specified offset.
func (b LittleEndianBlock) Float32(offset int) float32 {
	return math.Float32frombits(b.Uint32(offset))
}

// SetFloat32 places the float32 value at the specified offset.
func (b LittleEndianBlock) SetFloat32(offset int, value float32) {
	b.SetUint32(offset, math.Float32bits(value))
}

// Float64 returns the float64 value at the specified offset.
func (b LittleEndianBlock) Float64(offset int) float64 {
	return math.Float64frombits(b.Uint64(offset))
}

// SetFloat64 places the float64 value at the specified offset.
func (b LittleEndianBlock) SetFloat64(offset int, value float64) {
	b.SetUint64(offset, math.Float64bits(value))
}

// BigEndianBlock represents a fixed-size block of bytes that holds
// values encoded in Big Endian order.
type BigEndianBlock []byte

// Uint8 returns the uint8 value at the specified offset.
func (b BigEndianBlock) Uint8(offset int) uint8 {
	return b[offset]
}

// SetUint8 places the uint8 value at the specified offset.
func (b BigEndianBlock) SetUint8(offset int, value uint8) {
	b[offset] = value
}

// Int8 returns the int8 value at the specified offset.
func (b BigEndianBlock) Int8(offset int) int8 {
	return int8(b.Uint8(offset))
}

// SetInt8 places the int8 value at the specified offset.
func (b BigEndianBlock) SetInt8(offset int, value int8) {
	b.SetUint8(offset, uint8(value))
}

// Uint16 returns the uint16 value at the specified offset.
func (b BigEndianBlock) Uint16(offset int) uint16 {
	return uint16(b[offset+1])<<0 |
		uint16(b[offset+0])<<8
}

// SetUint16 places the uint16 value at the specified offset.
func (b BigEndianBlock) SetUint16(offset int, value uint16) {
	b[offset+1] = byte(value >> 0)
	b[offset+0] = byte(value >> 8)
}

// Int16 returns the int16 value at the specified offset.
func (b BigEndianBlock) Int16(offset int) int16 {
	return int16(b.Uint16(offset))
}

// SetInt16 places the int16 value at the specified offset.
func (b BigEndianBlock) SetInt16(offset int, value int16) {
	b.SetUint16(offset, uint16(value))
}

// Uint32 returns the uint32 value at the specified offset.
func (b BigEndianBlock) Uint32(offset int) uint32 {
	return uint32(b[offset+3])<<0 |
		uint32(b[offset+2])<<8 |
		uint32(b[offset+1])<<16 |
		uint32(b[offset+0])<<24
}

// SetUint32 places the uint32 value at the specified offset.
func (b BigEndianBlock) SetUint32(offset int, value uint32) {
	b[offset+3] = byte(value >> 0)
	b[offset+2] = byte(value >> 8)
	b[offset+1] = byte(value >> 16)
	b[offset+0] = byte(value >> 24)
}

// Int32 returns the int32 value at the specified offset.
func (b BigEndianBlock) Int32(offset int) int32 {
	return int32(b.Uint32(offset))
}

// SetInt32 places the int32 value at the specified offset.
func (b BigEndianBlock) SetInt32(offset int, value int32) {
	b.SetUint32(offset, uint32(value))
}

// Uint64 returns the uint64 value at the specified offset.
func (b BigEndianBlock) Uint64(offset int) uint64 {
	return uint64(b[offset+7])<<0 |
		uint64(b[offset+6])<<8 |
		uint64(b[offset+5])<<16 |
		uint64(b[offset+4])<<24 |
		uint64(b[offset+3])<<32 |
		uint64(b[offset+2])<<40 |
		uint64(b[offset+1])<<48 |
		uint64(b[offset+0])<<56
}

// SetUint64 places the uint64 value at the specified offset.
func (b BigEndianBlock) SetUint64(offset int, value uint64) {
	b[offset+7] = byte(value >> 0)
	b[offset+6] = byte(value >> 8)
	b[offset+5] = byte(value >> 16)
	b[offset+4] = byte(value >> 24)
	b[offset+3] = byte(value >> 32)
	b[offset+2] = byte(value >> 40)
	b[offset+1] = byte(value >> 48)
	b[offset+0] = byte(value >> 56)
}

// Int64 returns the int64 value at the specified offset.
func (b BigEndianBlock) Int64(offset int) int64 {
	return int64(b.Uint64(offset))
}

// SetInt64 places the int64 value at the specified offset.
func (b BigEndianBlock) SetInt64(offset int, value uint64) {
	b.SetUint64(offset, uint64(value))
}

// Float32 returns the float32 value at the specified offset.
func (b BigEndianBlock) Float32(offset int) float32 {
	return math.Float32frombits(b.Uint32(offset))
}

// SetFloat32 places the float32 value at the specified offset.
func (b BigEndianBlock) SetFloat32(offset int, value float32) {
	b.SetUint32(offset, math.Float32bits(value))
}

// Float64 returns the float64 value at the specified offset.
func (b BigEndianBlock) Float64(offset int) float64 {
	return math.Float64frombits(b.Uint64(offset))
}

// SetFloat64 places the float64 value at the specified offset.
func (b BigEndianBlock) SetFloat64(offset int, value float64) {
	b.SetUint64(offset, math.Float64bits(value))
}