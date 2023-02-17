package gblob_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gblob"
)

var _ = Describe("LittleEndianReader", func() {
	var (
		buffer *bytes.Buffer
		reader gblob.TypedReader
	)

	BeforeEach(func() {
		buffer = new(bytes.Buffer)
		reader = gblob.NewLittleEndianReader(buffer)
	})

	Specify("ReadUint8", func() {
		buffer.Write([]uint8{0x34, 0x65})

		value, err := reader.ReadUint8()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(uint8(0x34)))

		value, err = reader.ReadUint8()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(uint8(0x65)))
	})

	Specify("ReadInt8", func() {
		buffer.Write([]uint8{0x34, 0x65})

		value, err := reader.ReadInt8()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(int8(0x34)))

		value, err = reader.ReadInt8()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(int8(0x65)))
	})

	Specify("ReadUint16", func() {
		buffer.Write([]uint8{
			0x21, 0x34,
			0x55, 0x65,
		})

		value, err := reader.ReadUint16()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(uint16(0x3421)))

		value, err = reader.ReadUint16()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(uint16(0x6555)))
	})

	Specify("ReadInt16", func() {
		buffer.Write([]uint8{
			0x21, 0x34,
			0x55, 0x65,
		})

		value, err := reader.ReadInt16()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(int16(0x3421)))

		value, err = reader.ReadInt16()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(int16(0x6555)))
	})

	Specify("ReadUint32", func() {
		buffer.Write([]uint8{
			0x23, 0x71, 0x21, 0x34,
			0x61, 0x44, 0x55, 0x65,
		})

		value, err := reader.ReadUint32()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(uint32(0x34217123)))

		value, err = reader.ReadUint32()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(uint32(0x65554461)))
	})

	Specify("ReadInt32", func() {
		buffer.Write([]uint8{
			0x23, 0x71, 0x21, 0x34,
			0x61, 0x44, 0x55, 0x65,
		})

		value, err := reader.ReadInt32()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(int32(0x34217123)))

		value, err = reader.ReadInt32()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(int32(0x65554461)))
	})

	Specify("ReadUint64", func() {
		buffer.Write([]uint8{
			0x11, 0x72, 0x56, 0x98, 0x23, 0x71, 0x21, 0x34,
			0x04, 0x43, 0x85, 0x67, 0x61, 0x44, 0x55, 0x65,
		})

		value, err := reader.ReadUint64()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(uint64(0x3421712398567211)))

		value, err = reader.ReadUint64()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(uint64(0x6555446167854304)))
	})

	Specify("ReadInt64", func() {
		buffer.Write([]uint8{
			0x11, 0x72, 0x56, 0x98, 0x23, 0x71, 0x21, 0x34,
			0x04, 0x43, 0x85, 0x67, 0x61, 0x44, 0x55, 0x65,
		})

		value, err := reader.ReadInt64()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(int64(0x3421712398567211)))

		value, err = reader.ReadInt64()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(int64(0x6555446167854304)))
	})

	Specify("ReadFloat32", func() {
		buffer.Write([]uint8{
			0xCD, 0xCC, 0xAC, 0x40,
			0x9A, 0x99, 0x99, 0x3F,
		})

		value, err := reader.ReadFloat32()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(BeNumerically("~", float32(5.4), 0.0001))

		value, err = reader.ReadFloat32()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(BeNumerically("~", float32(1.2), 0.0001))
	})

	Specify("ReadFloat64", func() {
		buffer.Write([]uint8{
			0x9A, 0x99, 0x99, 0x99, 0x99, 0x99, 0x15, 0x40,
			0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0xF3, 0x3F,
		})

		value, err := reader.ReadFloat64()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(BeNumerically("~", 5.4, 0.00000001))

		value, err = reader.ReadFloat64()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(BeNumerically("~", 1.2, 0.00000001))
	})

	Specify("ReadBytes", func() {
		buffer.Write([]uint8{0x34, 0x65})

		target := make([]uint8, 2)
		Expect(reader.ReadBytes(target)).To(Succeed())
		Expect(target).To(Equal([]uint8{0x34, 0x65}))
	})
})

var _ = Describe("BigEndianReader", func() {
	var (
		buffer *bytes.Buffer
		reader gblob.TypedReader
	)

	BeforeEach(func() {
		buffer = new(bytes.Buffer)
		reader = gblob.NewBigEndianReader(buffer)
	})

	Specify("ReadUint8", func() {
		buffer.Write([]uint8{0x34, 0x65})

		value, err := reader.ReadUint8()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(uint8(0x34)))

		value, err = reader.ReadUint8()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(uint8(0x65)))
	})

	Specify("ReadInt8", func() {
		buffer.Write([]uint8{0x34, 0x65})

		value, err := reader.ReadInt8()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(int8(0x34)))

		value, err = reader.ReadInt8()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(int8(0x65)))
	})

	Specify("ReadUint16", func() {
		buffer.Write([]uint8{
			0x34, 0x21,
			0x65, 0x55,
		})

		value, err := reader.ReadUint16()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(uint16(0x3421)))

		value, err = reader.ReadUint16()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(uint16(0x6555)))
	})

	Specify("ReadInt16", func() {
		buffer.Write([]uint8{
			0x34, 0x21,
			0x65, 0x55,
		})

		value, err := reader.ReadInt16()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(int16(0x3421)))

		value, err = reader.ReadInt16()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(int16(0x6555)))
	})

	Specify("ReadUint32", func() {
		buffer.Write([]uint8{
			0x34, 0x21, 0x71, 0x23,
			0x65, 0x55, 0x44, 0x61,
		})

		value, err := reader.ReadUint32()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(uint32(0x34217123)))

		value, err = reader.ReadUint32()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(uint32(0x65554461)))
	})

	Specify("ReadInt32", func() {
		buffer.Write([]uint8{
			0x34, 0x21, 0x71, 0x23,
			0x65, 0x55, 0x44, 0x61,
		})

		value, err := reader.ReadInt32()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(int32(0x34217123)))

		value, err = reader.ReadInt32()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(int32(0x65554461)))
	})

	Specify("ReadUint64", func() {
		buffer.Write([]uint8{
			0x34, 0x21, 0x71, 0x23, 0x98, 0x56, 0x72, 0x11,
			0x65, 0x55, 0x44, 0x61, 0x67, 0x85, 0x43, 0x04,
		})

		value, err := reader.ReadUint64()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(uint64(0x3421712398567211)))

		value, err = reader.ReadUint64()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(uint64(0x6555446167854304)))
	})

	Specify("ReadInt64", func() {
		buffer.Write([]uint8{
			0x34, 0x21, 0x71, 0x23, 0x98, 0x56, 0x72, 0x11,
			0x65, 0x55, 0x44, 0x61, 0x67, 0x85, 0x43, 0x04,
		})

		value, err := reader.ReadInt64()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(int64(0x3421712398567211)))

		value, err = reader.ReadInt64()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(Equal(int64(0x6555446167854304)))
	})

	Specify("ReadFloat32", func() {
		buffer.Write([]uint8{
			0x40, 0xAC, 0xCC, 0xCD,
			0x3F, 0x99, 0x99, 0x9A,
		})

		value, err := reader.ReadFloat32()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(BeNumerically("~", float32(5.4), 0.0001))

		value, err = reader.ReadFloat32()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(BeNumerically("~", float32(1.2), 0.0001))
	})

	Specify("ReadFloat64", func() {
		buffer.Write([]uint8{
			0x40, 0x15, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9A,
			0x3F, 0xF3, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33,
		})

		value, err := reader.ReadFloat64()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(BeNumerically("~", 5.4, 0.00000001))

		value, err = reader.ReadFloat64()
		Expect(err).ToNot(HaveOccurred())
		Expect(value).To(BeNumerically("~", 1.2, 0.00000001))
	})

	Specify("ReadBytes", func() {
		buffer.Write([]uint8{0x34, 0x65})

		target := make([]uint8, 2)
		Expect(reader.ReadBytes(target)).To(Succeed())
		Expect(target).To(Equal([]uint8{0x34, 0x65}))
	})
})
