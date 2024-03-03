package gblob_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gblob"
	"github.com/mokiat/gog"
)

var _ = Describe("PackedDecoder", func() {
	var (
		buffer  *bytes.Buffer
		decoder *gblob.PackedDecoder
	)

	seq := func(values ...uint8) []uint8 {
		return values
	}

	type CustomArray [3]uint16

	type CustomSlice []uint16

	type CustomString string

	type CustomStruct struct {
		A uint16
		B *uint8
		C uint8
	}

	BeforeEach(func() {
		buffer = new(bytes.Buffer)
		decoder = gblob.NewLittleEndianPackedDecoder(buffer)
	})

	DescribeTable("types",
		func(data []byte, target any, expected any) {
			buffer.Write(data)
			Expect(decoder.Decode(target)).To(Succeed())
			Expect(target).To(Equal(expected))
		},
		Entry("bool",
			seq(0x01),
			gog.PtrOf(false),
			gog.PtrOf(true),
		),
		Entry("*bool",
			seq(0x01),
			gog.PtrOf((*bool)(nil)),
			gog.PtrOf(gog.PtrOf(true)),
		),
		Entry("uint8",
			seq(0x3F),
			gog.PtrOf(uint8(0x00)),
			gog.PtrOf(uint8(0x3F)),
		),
		Entry("*uint8",
			seq(0x3F),
			gog.PtrOf((*uint8)(nil)),
			gog.PtrOf(gog.PtrOf(uint8(0x3F))),
		),
		Entry("int8",
			seq(0x3F),
			gog.PtrOf(int8(0x00)),
			gog.PtrOf(int8(0x3F)),
		),
		Entry("*int8",
			seq(0x3F),
			gog.PtrOf((*int8)(nil)),
			gog.PtrOf(gog.PtrOf(int8(0x3F))),
		),
		Entry("uint16",
			seq(0xCA, 0xF1),
			gog.PtrOf(uint16(0x0000)),
			gog.PtrOf(uint16(0xF1CA)),
		),
		Entry("*uint16",
			seq(0xCA, 0xF1),
			gog.PtrOf((*uint16)(nil)),
			gog.PtrOf(gog.PtrOf(uint16(0xF1CA))),
		),
		Entry("int16",
			seq(0xCA, 0x31),
			gog.PtrOf(int16(0x0000)),
			gog.PtrOf(int16(0x31CA)),
		),
		Entry("*int16",
			seq(0xCA, 0x31),
			gog.PtrOf((*int16)(nil)),
			gog.PtrOf(gog.PtrOf(int16(0x31CA))),
		),
		Entry("uint32",
			seq(0x32, 0x76, 0xCA, 0xF1),
			gog.PtrOf(uint32(0x00000000)),
			gog.PtrOf(uint32(0xF1CA7632)),
		),
		Entry("*uint32",
			seq(0x32, 0x76, 0xCA, 0xF1),
			gog.PtrOf((*uint32)(nil)),
			gog.PtrOf(gog.PtrOf(uint32(0xF1CA7632))),
		),
		Entry("int32",
			seq(0x32, 0x76, 0xCA, 0x51),
			gog.PtrOf(int32(0x00000000)),
			gog.PtrOf(int32(0x51CA7632)),
		),
		Entry("*int32",
			seq(0x32, 0x76, 0xCA, 0x51),
			gog.PtrOf((*int32)(nil)),
			gog.PtrOf(gog.PtrOf(int32(0x51CA7632))),
		),
		Entry("uint64",
			seq(0x21, 0x73, 0xC4, 0xA3, 0x32, 0x76, 0xCA, 0xF1),
			gog.PtrOf(uint64(0x0000000000000000)),
			gog.PtrOf(uint64(0xF1CA7632A3C47321)),
		),
		Entry("*uint64",
			seq(0x21, 0x73, 0xC4, 0xA3, 0x32, 0x76, 0xCA, 0xF1),
			gog.PtrOf((*uint64)(nil)),
			gog.PtrOf(gog.PtrOf(uint64(0xF1CA7632A3C47321))),
		),
		Entry("int64",
			seq(0x21, 0x73, 0xC4, 0xA3, 0x32, 0x76, 0xCA, 0x31),
			gog.PtrOf(int64(0x0000000000000000)),
			gog.PtrOf(int64(0x31CA7632A3C47321)),
		),
		Entry("*int64",
			seq(0x21, 0x73, 0xC4, 0xA3, 0x32, 0x76, 0xCA, 0x31),
			gog.PtrOf((*int64)(nil)),
			gog.PtrOf(gog.PtrOf(int64(0x31CA7632A3C47321))),
		),
		Entry("float32",
			seq(0xCD, 0xCC, 0x6C, 0x40),
			gog.PtrOf(float32(0.0)),
			gog.PtrOf(float32(3.7)),
		),
		Entry("*float32",
			seq(0xCD, 0xCC, 0x6C, 0x40),
			gog.PtrOf((*float32)(nil)),
			gog.PtrOf(gog.PtrOf(float32(3.7))),
		),
		Entry("float64",
			seq(0x9A, 0x99, 0x99, 0x99, 0x99, 0x99, 0x0D, 0x40),
			gog.PtrOf(float64(0.0)),
			gog.PtrOf(float64(3.7)),
		),
		Entry("*float64",
			seq(0x9A, 0x99, 0x99, 0x99, 0x99, 0x99, 0x0D, 0x40),
			gog.PtrOf((*float64)(nil)),
			gog.PtrOf(gog.PtrOf(float64(3.7))),
		),
		Entry("array",
			seq(0xFA, 0x31, 0xAC, 0x45, 0x21, 0x5F),
			gog.PtrOf([3]uint16{}),
			gog.PtrOf([3]uint16{0x31FA, 0x45AC, 0x5F21}),
		),
		Entry("*array",
			seq(0xFA, 0x31, 0xAC, 0x45, 0x21, 0x5F),
			gog.PtrOf((*[3]uint16)(nil)),
			gog.PtrOf(gog.PtrOf([3]uint16{0x31FA, 0x45AC, 0x5F21})),
		),
		Entry("CustomArray",
			seq(0xFA, 0x31, 0xAC, 0x45, 0x21, 0x5F),
			gog.PtrOf(CustomArray{}),
			gog.PtrOf(CustomArray{0x31FA, 0x45AC, 0x5F21}),
		),
		Entry("*CustomArray",
			seq(0xFA, 0x31, 0xAC, 0x45, 0x21, 0x5F),
			gog.PtrOf((*CustomArray)(nil)),
			gog.PtrOf(gog.PtrOf(CustomArray{0x31FA, 0x45AC, 0x5F21})),
		),
		Entry("bytes",
			seq(
				0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // length
				0x31, 0xFA, 0x45, 0xAC, 0x5F, 0x21, // items
			),
			gog.PtrOf([]uint8(nil)),
			gog.PtrOf([]uint8{0x31, 0xFA, 0x45, 0xAC, 0x5F, 0x21}),
		),
		Entry("*bytes",
			seq(
				0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // length
				0x31, 0xFA, 0x45, 0xAC, 0x5F, 0x21, // items
			),
			gog.PtrOf((*[]uint8)(nil)),
			gog.PtrOf(gog.PtrOf([]uint8{0x31, 0xFA, 0x45, 0xAC, 0x5F, 0x21})),
		),
		Entry("slice",
			seq(
				0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // length
				0xFA, 0x31, 0xAC, 0x45, 0x21, 0x5F, // items
			),
			gog.PtrOf([]uint16(nil)),
			gog.PtrOf([]uint16{0x31FA, 0x45AC, 0x5F21}),
		),
		Entry("*slice",
			seq(
				0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // length
				0xFA, 0x31, 0xAC, 0x45, 0x21, 0x5F, // items
			),
			gog.PtrOf((*[]uint16)(nil)),
			gog.PtrOf(gog.PtrOf([]uint16{0x31FA, 0x45AC, 0x5F21})),
		),
		Entry("CustomSlice",
			seq(
				0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // length
				0xFA, 0x31, 0xAC, 0x45, 0x21, 0x5F, // items
			),
			gog.PtrOf(CustomSlice(nil)),
			gog.PtrOf(CustomSlice{0x31FA, 0x45AC, 0x5F21}),
		),
		Entry("*CustomSlice",
			seq(
				0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // length
				0xFA, 0x31, 0xAC, 0x45, 0x21, 0x5F, // items
			),
			gog.PtrOf((*CustomSlice)(nil)),
			gog.PtrOf(gog.PtrOf(CustomSlice{0x31FA, 0x45AC, 0x5F21})),
		),
		Entry("map",
			seq(
				0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // length
				0xAC, 0xFA, 0x37, // first entry
				0x05, 0xA2, 0x51, // second entry
			),
			gog.PtrOf(map[uint8]uint16(nil)),
			gog.PtrOf(map[uint8]uint16{
				0xAC: 0x37FA,
				0x05: 0x51A2,
			}),
		),
		Entry("*map",
			seq(
				0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // length
				0xAC, 0xFA, 0x37, // first entry
				0x05, 0xA2, 0x51, // second entry
			),
			gog.PtrOf((*map[uint8]uint16)(nil)),
			gog.PtrOf(gog.PtrOf(map[uint8]uint16{
				0xAC: 0x37FA,
				0x05: 0x51A2,
			})),
		),
		Entry("string",
			seq(
				0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // length
				0x68, 0x65, 0x6C, 0x6C, 0x6F, // items
			),
			gog.PtrOf(""),
			gog.PtrOf("hello"),
		),
		Entry("*string",
			seq(
				0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // length
				0x68, 0x65, 0x6C, 0x6C, 0x6F, // items
			),
			gog.PtrOf((*string)(nil)),
			gog.PtrOf(gog.PtrOf("hello")),
		),
		Entry("CustomString",
			seq(
				0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // length
				0x68, 0x65, 0x6C, 0x6C, 0x6F, // items
			),
			gog.PtrOf(CustomString("")),
			gog.PtrOf(CustomString("hello")),
		),
		Entry("*CustomString",
			seq(
				0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // length
				0x68, 0x65, 0x6C, 0x6C, 0x6F, // items
			),
			gog.PtrOf((*CustomString)(nil)),
			gog.PtrOf(gog.PtrOf(CustomString("hello"))),
		),
		Entry("struct",
			seq(0x66, 0x55, 0xFF, 0x01),
			&CustomStruct{},
			&CustomStruct{
				A: 0x5566,
				B: gog.PtrOf(uint8(0xFF)),
				C: 0x01,
			},
		),
		Entry("*struct",
			seq(0x66, 0x55, 0xFF, 0x01),
			gog.PtrOf((*CustomStruct)(nil)),
			gog.PtrOf(&CustomStruct{
				A: 0x5566,
				B: gog.PtrOf(uint8(0xFF)),
				C: 0x01,
			}),
		),
		Entry("Decodable",
			seq(0x39),
			&testDecodable{},
			&testDecodable{
				value: 0x39,
			},
		),
		Entry("*Decodable",
			seq(0x39),
			gog.PtrOf((*testDecodable)(nil)),
			gog.PtrOf(&testDecodable{
				value: 0x39,
			}),
		),
	)
})

type testDecodable struct {
	value byte
}

var _ gblob.PackedDecodable = (*testDecodable)(nil)

func (d *testDecodable) DecodePacked(reader gblob.TypedReader) error {
	v, err := reader.ReadUint8()
	if err != nil {
		return err
	}
	d.value = v
	return nil
}
