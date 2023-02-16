package gblob_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gblob"
)

var _ = Describe("LittleEndianBlock", func() {
	// pad is just used to test offsets.
	const pad = uint8(0x00)

	Specify("Uint8", func() {
		block := gblob.LittleEndianBlock{pad, pad, 0x13}
		Expect(block.Uint8(2)).To(Equal(uint8(0x13)))
	})

	Specify("SetUint8", func() {
		block := gblob.LittleEndianBlock{pad, pad, pad}
		block.SetUint8(2, 0x13)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x13}))
	})

	Specify("Int8", func() {
		block := gblob.LittleEndianBlock{pad, pad, 0x13}
		Expect(block.Int8(2)).To(Equal(int8(0x13)))
	})

	Specify("SetInt8", func() {
		block := gblob.LittleEndianBlock{pad, pad, pad}
		block.SetInt8(2, 0x13)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x13}))
	})

	Specify("Uint16", func() {
		block := gblob.LittleEndianBlock{pad, pad, 0x52, 0x13}
		Expect(block.Uint16(2)).To(Equal(uint16(0x1352)))
	})

	Specify("SetUint16", func() {
		block := gblob.LittleEndianBlock{pad, pad, pad, pad}
		block.SetUint16(2, 0x1352)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x52, 0x13}))
	})

	Specify("Int16", func() {
		block := gblob.LittleEndianBlock{pad, pad, 0x52, 0x13}
		Expect(block.Int16(2)).To(Equal(int16(0x1352)))
	})

	Specify("SetInt16", func() {
		block := gblob.LittleEndianBlock{pad, pad, pad, pad}
		block.SetInt16(2, 0x1352)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x52, 0x13}))
	})

	Specify("Uint32", func() {
		block := gblob.LittleEndianBlock{pad, pad, 0x52, 0x13, 0x44, 0x37}
		Expect(block.Uint32(2)).To(Equal(uint32(0x37441352)))
	})

	Specify("SetUint32", func() {
		block := gblob.LittleEndianBlock{pad, pad, pad, pad, pad, pad}
		block.SetUint32(2, 0x37441352)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x52, 0x13, 0x44, 0x37}))
	})

	Specify("Int32", func() {
		block := gblob.LittleEndianBlock{pad, pad, 0x52, 0x13, 0x44, 0x37}
		Expect(block.Int32(2)).To(Equal(int32(0x37441352)))
	})

	Specify("SetInt32", func() {
		block := gblob.LittleEndianBlock{pad, pad, pad, pad, pad, pad}
		block.SetInt32(2, 0x37441352)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x52, 0x13, 0x44, 0x37}))
	})

	Specify("Uint64", func() {
		block := gblob.LittleEndianBlock{pad, pad, 0x52, 0x13, 0x44, 0x37, 0x11, 0x71, 0x91, 0x01}
		Expect(block.Uint64(2)).To(Equal(uint64(0x0191711137441352)))
	})

	Specify("SetUint64", func() {
		block := gblob.LittleEndianBlock{pad, pad, pad, pad, pad, pad, pad, pad, pad, pad}
		block.SetUint64(2, 0x0191711137441352)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x52, 0x13, 0x44, 0x37, 0x11, 0x71, 0x91, 0x01}))
	})

	Specify("Int64", func() {
		block := gblob.LittleEndianBlock{pad, pad, 0x52, 0x13, 0x44, 0x37, 0x11, 0x71, 0x91, 0x01}
		Expect(block.Int64(2)).To(Equal(int64(0x0191711137441352)))
	})

	Specify("SetInt64", func() {
		block := gblob.LittleEndianBlock{pad, pad, pad, pad, pad, pad, pad, pad, pad, pad}
		block.SetInt64(2, 0x0191711137441352)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x52, 0x13, 0x44, 0x37, 0x11, 0x71, 0x91, 0x01}))
	})

	Specify("Float32", func() {
		block := gblob.LittleEndianBlock{pad, pad, 0x8F, 0xC2, 0xBD, 0x40}
		Expect(block.Float32(2)).To(BeNumerically("~", 5.93, 0.0001))
	})

	Specify("SetFloat32", func() {
		block := gblob.LittleEndianBlock{pad, pad, pad, pad, pad, pad}
		block.SetFloat32(2, 5.93)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x8F, 0xC2, 0xBD, 0x40}))
	})

	Specify("Float64", func() {
		block := gblob.LittleEndianBlock{pad, pad, 0xB8, 0x1E, 0x85, 0xEB, 0x51, 0xB8, 0x17, 0x40}
		Expect(block.Float64(2)).To(BeNumerically("~", 5.93, 0.0000000001))
	})

	Specify("SetFloat64", func() {
		block := gblob.LittleEndianBlock{pad, pad, pad, pad, pad, pad, pad, pad, pad, pad}
		block.SetFloat64(2, 5.93)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0xB8, 0x1E, 0x85, 0xEB, 0x51, 0xB8, 0x17, 0x40}))
	})
})

var _ = Describe("BigEndianBlock", func() {
	// pad is just used to test offsets.
	const pad = uint8(0x00)

	Specify("Uint8", func() {
		block := gblob.BigEndianBlock{pad, pad, 0x13}
		Expect(block.Uint8(2)).To(Equal(uint8(0x13)))
	})

	Specify("SetUint8", func() {
		block := gblob.BigEndianBlock{pad, pad, pad}
		block.SetUint8(2, 0x13)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x13}))
	})

	Specify("Int8", func() {
		block := gblob.BigEndianBlock{pad, pad, 0x13}
		Expect(block.Int8(2)).To(Equal(int8(0x13)))
	})

	Specify("SetInt8", func() {
		block := gblob.BigEndianBlock{pad, pad, pad}
		block.SetInt8(2, 0x13)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x13}))
	})

	Specify("Uint16", func() {
		block := gblob.BigEndianBlock{pad, pad, 0x13, 0x52}
		Expect(block.Uint16(2)).To(Equal(uint16(0x1352)))
	})

	Specify("SetUint16", func() {
		block := gblob.BigEndianBlock{pad, pad, pad, pad}
		block.SetUint16(2, 0x1352)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x13, 0x52}))
	})

	Specify("Int16", func() {
		block := gblob.BigEndianBlock{pad, pad, 0x13, 0x52}
		Expect(block.Int16(2)).To(Equal(int16(0x1352)))
	})

	Specify("SetInt16", func() {
		block := gblob.BigEndianBlock{pad, pad, pad, pad}
		block.SetInt16(2, 0x1352)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x13, 0x52}))
	})

	Specify("Uint32", func() {
		block := gblob.BigEndianBlock{pad, pad, 0x37, 0x44, 0x13, 0x52}
		Expect(block.Uint32(2)).To(Equal(uint32(0x37441352)))
	})

	Specify("SetUint32", func() {
		block := gblob.BigEndianBlock{pad, pad, pad, pad, pad, pad}
		block.SetUint32(2, 0x37441352)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x37, 0x44, 0x13, 0x52}))
	})

	Specify("Int32", func() {
		block := gblob.BigEndianBlock{pad, pad, 0x37, 0x44, 0x13, 0x52}
		Expect(block.Int32(2)).To(Equal(int32(0x37441352)))
	})

	Specify("SetInt32", func() {
		block := gblob.BigEndianBlock{pad, pad, pad, pad, pad, pad}
		block.SetInt32(2, 0x37441352)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x37, 0x44, 0x13, 0x52}))
	})

	Specify("Uint64", func() {
		block := gblob.BigEndianBlock{pad, pad, 0x01, 0x91, 0x71, 0x11, 0x37, 0x44, 0x13, 0x52}
		Expect(block.Uint64(2)).To(Equal(uint64(0x0191711137441352)))
	})

	Specify("SetUint64", func() {
		block := gblob.BigEndianBlock{pad, pad, pad, pad, pad, pad, pad, pad, pad, pad}
		block.SetUint64(2, 0x0191711137441352)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x01, 0x91, 0x71, 0x11, 0x37, 0x44, 0x13, 0x52}))
	})

	Specify("Int64", func() {
		block := gblob.BigEndianBlock{pad, pad, 0x01, 0x91, 0x71, 0x11, 0x37, 0x44, 0x13, 0x52}
		Expect(block.Int64(2)).To(Equal(int64(0x0191711137441352)))
	})

	Specify("SetInt64", func() {
		block := gblob.BigEndianBlock{pad, pad, pad, pad, pad, pad, pad, pad, pad, pad}
		block.SetInt64(2, 0x0191711137441352)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x01, 0x91, 0x71, 0x11, 0x37, 0x44, 0x13, 0x52}))
	})

	Specify("Float32", func() {
		block := gblob.BigEndianBlock{pad, pad, 0x40, 0xBD, 0xC2, 0x8F}
		Expect(block.Float32(2)).To(BeNumerically("~", 5.93, 0.0001))
	})

	Specify("SetFloat32", func() {
		block := gblob.BigEndianBlock{pad, pad, pad, pad, pad, pad}
		block.SetFloat32(2, 5.93)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x40, 0xBD, 0xC2, 0x8F}))
	})

	Specify("Float64", func() {
		block := gblob.BigEndianBlock{pad, pad, 0x40, 0x17, 0xB8, 0x51, 0xEB, 0x85, 0x1E, 0xB8}
		Expect(block.Float64(2)).To(BeNumerically("~", 5.93, 0.0000000001))
	})

	Specify("SetFloat64", func() {
		block := gblob.BigEndianBlock{pad, pad, pad, pad, pad, pad, pad, pad, pad, pad}
		block.SetFloat64(2, 5.93)
		Expect([]uint8(block[2:])).To(Equal([]uint8{0x40, 0x17, 0xB8, 0x51, 0xEB, 0x85, 0x1E, 0xB8}))
	})
})
