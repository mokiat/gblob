package gblob_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gblob"
)

var _ = Describe("LittleEndianWriter", func() {
	var (
		buffer *bytes.Buffer
		writer gblob.TypedWriter
	)

	BeforeEach(func() {
		buffer = new(bytes.Buffer)
		writer = gblob.NewLittleEndianWriter(buffer)
	})

	Specify("WriteUint8", func() {
		writer.WriteUint8(0x34)
		writer.WriteUint8(0x65)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x34, 0x65,
		}))
	})

	Specify("WriteInt8", func() {
		writer.WriteInt8(0x34)
		writer.WriteInt8(0x65)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x34, 0x65,
		}))
	})

	Specify("WriteUint16", func() {
		writer.WriteUint16(0x3421)
		writer.WriteUint16(0x6555)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x21, 0x34,
			0x55, 0x65,
		}))
	})

	Specify("WriteInt16", func() {
		writer.WriteInt16(0x3421)
		writer.WriteInt16(0x6555)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x21, 0x34,
			0x55, 0x65,
		}))
	})

	Specify("WriteUint32", func() {
		writer.WriteUint32(0x34217123)
		writer.WriteUint32(0x65554461)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x23, 0x71, 0x21, 0x34,
			0x61, 0x44, 0x55, 0x65,
		}))
	})

	Specify("WriteInt32", func() {
		writer.WriteInt32(0x34217123)
		writer.WriteInt32(0x65554461)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x23, 0x71, 0x21, 0x34,
			0x61, 0x44, 0x55, 0x65,
		}))
	})

	Specify("WriteUint64", func() {
		writer.WriteUint64(0x3421712398567211)
		writer.WriteUint64(0x6555446167854304)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x11, 0x72, 0x56, 0x98, 0x23, 0x71, 0x21, 0x34,
			0x04, 0x43, 0x85, 0x67, 0x61, 0x44, 0x55, 0x65,
		}))
	})

	Specify("WriteInt64", func() {
		writer.WriteInt64(0x3421712398567211)
		writer.WriteInt64(0x6555446167854304)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x11, 0x72, 0x56, 0x98, 0x23, 0x71, 0x21, 0x34,
			0x04, 0x43, 0x85, 0x67, 0x61, 0x44, 0x55, 0x65,
		}))
	})

	Specify("WriteFloat32", func() {
		writer.WriteFloat32(5.4)
		writer.WriteFloat32(1.2)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0xCD, 0xCC, 0xAC, 0x40,
			0x9A, 0x99, 0x99, 0x3F,
		}))
	})

	Specify("WriteFloat64", func() {
		writer.WriteFloat64(5.4)
		writer.WriteFloat64(1.2)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x9A, 0x99, 0x99, 0x99, 0x99, 0x99, 0x15, 0x40,
			0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0xF3, 0x3F,
		}))
	})

	Specify("WriteBytes", func() {
		writer.WriteBytes([]uint8{0x12, 0x34})
		writer.WriteBytes([]uint8{0x98, 0x76})
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x12, 0x34,
			0x98, 0x76,
		}))
	})
})

var _ = Describe("BigEndianWriter", func() {
	var (
		buffer *bytes.Buffer
		writer gblob.TypedWriter
	)

	BeforeEach(func() {
		buffer = new(bytes.Buffer)
		writer = gblob.NewBigEndianWriter(buffer)
	})

	Specify("WriteUint8", func() {
		writer.WriteUint8(0x34)
		writer.WriteUint8(0x65)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x34, 0x65,
		}))
	})

	Specify("WriteInt8", func() {
		writer.WriteInt8(0x34)
		writer.WriteInt8(0x65)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x34, 0x65,
		}))
	})

	Specify("WriteUint16", func() {
		writer.WriteUint16(0x3421)
		writer.WriteUint16(0x6555)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x34, 0x21,
			0x65, 0x55,
		}))
	})

	Specify("WriteInt16", func() {
		writer.WriteInt16(0x3421)
		writer.WriteInt16(0x6555)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x34, 0x21,
			0x65, 0x55,
		}))
	})

	Specify("WriteUint32", func() {
		writer.WriteUint32(0x34217123)
		writer.WriteUint32(0x65554461)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x34, 0x21, 0x71, 0x23,
			0x65, 0x55, 0x44, 0x61,
		}))
	})

	Specify("WriteInt32", func() {
		writer.WriteInt32(0x34217123)
		writer.WriteInt32(0x65554461)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x34, 0x21, 0x71, 0x23,
			0x65, 0x55, 0x44, 0x61,
		}))
	})

	Specify("WriteUint64", func() {
		writer.WriteUint64(0x3421712398567211)
		writer.WriteUint64(0x6555446167854304)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x34, 0x21, 0x71, 0x23, 0x98, 0x56, 0x72, 0x11,
			0x65, 0x55, 0x44, 0x61, 0x67, 0x85, 0x43, 0x04,
		}))
	})

	Specify("WriteInt64", func() {
		writer.WriteInt64(0x3421712398567211)
		writer.WriteInt64(0x6555446167854304)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x34, 0x21, 0x71, 0x23, 0x98, 0x56, 0x72, 0x11,
			0x65, 0x55, 0x44, 0x61, 0x67, 0x85, 0x43, 0x04,
		}))
	})

	Specify("WriteFloat32", func() {
		writer.WriteFloat32(5.4)
		writer.WriteFloat32(1.2)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x40, 0xAC, 0xCC, 0xCD,
			0x3F, 0x99, 0x99, 0x9A,
		}))
	})

	Specify("WriteFloat64", func() {
		writer.WriteFloat64(5.4)
		writer.WriteFloat64(1.2)
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x40, 0x15, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9A,
			0x3F, 0xF3, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33,
		}))
	})

	Specify("WriteBytes", func() {
		writer.WriteBytes([]uint8{0x12, 0x34})
		writer.WriteBytes([]uint8{0x98, 0x76})
		Expect(buffer.Bytes()).To(Equal([]uint8{
			0x12, 0x34,
			0x98, 0x76,
		}))
	})
})
