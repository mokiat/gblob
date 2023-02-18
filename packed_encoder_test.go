package gblob_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gblob"
	"github.com/mokiat/gog"
)

var _ = Describe("PackedEncoder", func() {
	var (
		buffer  *bytes.Buffer
		encoder *gblob.PackedEncoder
	)

	seq := func(values ...uint8) []uint8 {
		return values
	}

	BeforeEach(func() {
		buffer = new(bytes.Buffer)
		encoder = gblob.NewLittleEndianPackedEncoder(buffer)
	})

	DescribeTable("types",
		func(data any, expected []byte) {
			Expect(encoder.Encode(data)).To(Succeed())
			Expect(buffer.Bytes()).To(Equal(expected))
		},
		Entry("uint8",
			uint8(0x13),
			seq(0x13),
		),
		Entry("*uint8",
			gog.PtrOf(uint8(0x13)),
			seq(0x13),
		),
		Entry("int8",
			int8(0x13),
			seq(0x13),
		),
		Entry("*int8",
			gog.PtrOf(int8(0x13)),
			seq(0x13),
		),
		Entry("uint16",
			uint16(0xF1CA),
			seq(0xCA, 0xF1),
		),
		Entry("*uint16",
			gog.PtrOf(uint16(0xF1CA)),
			seq(0xCA, 0xF1),
		),
		Entry("int16",
			int16(0x31CA),
			seq(0xCA, 0x31),
		),
		Entry("*int16",
			gog.PtrOf(int16(0x31CA)),
			seq(0xCA, 0x31),
		),
		Entry("uint32",
			uint32(0xF1CA7632),
			seq(0x32, 0x76, 0xCA, 0xF1),
		),
		Entry("*uint32",
			gog.PtrOf(uint32(0xF1CA7632)),
			seq(0x32, 0x76, 0xCA, 0xF1),
		),
		Entry("int32",
			int32(0x31CA7632),
			seq(0x32, 0x76, 0xCA, 0x31),
		),
		Entry("*int32",
			gog.PtrOf(int32(0x31CA7632)),
			seq(0x32, 0x76, 0xCA, 0x31),
		),
		Entry("uint64",
			uint64(0xF1CA7632A3C47321),
			seq(0x21, 0x73, 0xC4, 0xA3, 0x32, 0x76, 0xCA, 0xF1),
		),
		Entry("*uint64",
			gog.PtrOf(uint64(0xF1CA7632A3C47321)),
			seq(0x21, 0x73, 0xC4, 0xA3, 0x32, 0x76, 0xCA, 0xF1),
		),
		Entry("int64",
			int64(0x31CA7632A3C47321),
			seq(0x21, 0x73, 0xC4, 0xA3, 0x32, 0x76, 0xCA, 0x31),
		),
		Entry("*int64",
			gog.PtrOf(int64(0x31CA7632A3C47321)),
			seq(0x21, 0x73, 0xC4, 0xA3, 0x32, 0x76, 0xCA, 0x31),
		),
	)
})
