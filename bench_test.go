package gblob_test

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"testing"

	"github.com/mokiat/gblob"
)

//
// Block API comparison follows.
//

func Benchmark_BlockAPI_Block(b *testing.B) {
	const (
		itemCount   = 1024
		expectedSum = itemCount * (0 + itemCount - 1) / 2
	)

	data := make(gblob.LittleEndianBlock, itemCount*4)

	b.ResetTimer()

	for range b.N {
		for i := range itemCount {
			data.SetUint32(i*4, uint32(i))
		}
		sum := uint32(0)
		for i := range itemCount {
			value := data.Uint32(i * 4)
			sum += value
		}
		if sum != expectedSum {
			b.Errorf("Sum %d is not equal to %d", sum, expectedSum)
		}
	}
}

func Benchmark_BlockAPI_ByteOrder(b *testing.B) {
	const (
		itemCount   = 1024
		expectedSum = itemCount * (0 + itemCount - 1) / 2
	)

	data := make([]byte, itemCount*4)

	b.ResetTimer()

	for range b.N {
		for i := range itemCount {
			binary.LittleEndian.PutUint32(data[i*4:], uint32(i))
		}
		sum := uint32(0)
		for i := range itemCount {
			value := binary.LittleEndian.Uint32(data[i*4:])
			sum += value
		}
		if sum != expectedSum {
			b.Errorf("Sum %d is not equal to %d", sum, expectedSum)
		}
	}
}

//
// Writer API comparison follows.
//

func Benchmark_Writer_TypedWriter(b *testing.B) {
	const itemCount = 1024

	data := bytes.NewBuffer(make([]byte, itemCount*4))

	writer := gblob.NewLittleEndianWriter(data)

	b.ResetTimer()

	for range b.N {
		data.Reset()
		for i := range itemCount {
			writer.WriteUint32(uint32(i))
		}
		if data.Len() != itemCount*4 {
			b.Errorf("Length %d is not equal to %d", data.Len(), itemCount)
		}
	}
}

func Benchmark_Writer_BinaryWrite(b *testing.B) {
	const itemCount = 1024

	data := bytes.NewBuffer(make([]byte, itemCount*4))

	b.ResetTimer()

	for range b.N {
		data.Reset()
		for i := range itemCount {
			binary.Write(data, binary.LittleEndian, uint32(i))
		}
		if data.Len() != itemCount*4 {
			b.Errorf("Length %d is not equal to %d", data.Len(), itemCount)
		}
	}
}

//
// Reader API comparison follows.
//

func Benchmark_Reader_TypedReader(b *testing.B) {
	const itemCount = 1024

	data := make([]byte, itemCount*4)
	for i := range data {
		data[i] = (byte(i % 256))
	}
	seeker := bytes.NewReader(data)

	reader := gblob.NewLittleEndianReader(seeker)

	b.ResetTimer()

	for range b.N {
		seeker.Reset(data)
		sum := uint32(0)
		for range itemCount {
			val, _ := reader.ReadUint32()
			sum += val
		}
		if sum <= 0 {
			b.Errorf("Sum %d is not positive", sum)
		}
	}
}

func Benchmark_Reader_BinaryRead(b *testing.B) {
	const itemCount = 1024

	data := make([]byte, itemCount*4)
	for i := range data {
		data[i] = (byte(i % 256))
	}
	seeker := bytes.NewReader(data)

	b.ResetTimer()

	for range b.N {
		seeker.Reset(data)
		sum := uint32(0)
		for range itemCount {
			var val uint32
			binary.Read(seeker, binary.LittleEndian, &val)
			sum += val
		}
		if sum <= 0 {
			b.Errorf("Sum %d is not positive", sum)
		}
	}
}

//
// Encode API comparison follows.
//

func Benchmark_Encoder_PackedEncoder(b *testing.B) {
	const itemCount = 1024

	type encodeStruct struct {
		A uint32
		B int16
		C float64
		D float32
		E [32]byte
		F struct {
			G byte
		}
		H []uint64
	}
	template := encodeStruct{
		A: 0,
		B: 10,
		C: 32.0,
		D: 100.0,
		E: [32]byte{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		},
		F: struct{ G byte }{
			G: 255,
		},
		H: make([]uint64, 256),
	}

	data := bytes.NewBuffer(make([]byte, 0, itemCount*1024))

	b.ResetTimer()

	for range b.N {
		data.Reset()

		for range itemCount {
			if err := gblob.NewLittleEndianPackedEncoder(data).Encode(template); err != nil {
				panic(err)
			}
		}
		if len := data.Len(); len <= 0 {
			b.Errorf("Length %d is not positive", data.Len())
		}
	}
}

func Benchmark_Encoder_GobEncoder(b *testing.B) {
	const itemCount = 1024

	type encodeStruct struct {
		A uint32
		B int16
		C float64
		D float32
		E [32]byte
		F struct {
			G byte
		}
		H []uint64
	}
	template := encodeStruct{
		A: 0,
		B: 10,
		C: 32.0,
		D: 100.0,
		E: [32]byte{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		},
		F: struct{ G byte }{
			G: 255,
		},
		H: make([]uint64, 256),
	}

	data := bytes.NewBuffer(make([]byte, 0, itemCount*1024))

	b.ResetTimer()

	for range b.N {
		data.Reset()

		for range itemCount {
			if err := gob.NewEncoder(data).Encode(template); err != nil {
				panic(err)
			}
		}
		if len := data.Len(); len <= 0 {
			b.Errorf("Length %d is not positive", data.Len())
		}
	}
}

//
// Decode API comparison follows.
//

func Benchmark_Decoder_PackedDecoder(b *testing.B) {
	const itemCount = 1024

	type encodeStruct struct {
		A uint32
		B int16
		C float64
		D float32
		E [32]byte
		F struct {
			G byte
		}
		H []uint64
	}

	data := bytes.NewBuffer(make([]byte, 0, itemCount*1024))
	template := encodeStruct{
		A: 0,
		B: 10,
		C: 32.0,
		D: 100.0,
		E: [32]byte{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		},
		F: struct{ G byte }{
			G: 255,
		},
		H: make([]uint64, 256),
	}
	for range itemCount {
		if err := gblob.NewLittleEndianPackedEncoder(data).Encode(template); err != nil {
			panic(err)
		}
	}

	seeker := bytes.NewReader(data.Bytes())

	b.ResetTimer()

	for range b.N {
		seeker.Reset(data.Bytes())

		for range itemCount {
			var template encodeStruct
			if err := gblob.NewLittleEndianPackedDecoder(seeker).Decode(&template); err != nil {
				panic(err)
			}
			if template.B != 10 {
				b.Errorf("Field B %d is not equal to 10", template.B)
			}
		}
	}
}

func Benchmark_Decoder_GobDecoder(b *testing.B) {
	const itemCount = 1024

	type encodeStruct struct {
		A uint32
		B int16
		C float64
		D float32
		E [32]byte
		F struct {
			G byte
		}
		H []uint64
	}

	data := bytes.NewBuffer(make([]byte, 0, itemCount*1024))
	template := encodeStruct{
		A: 0,
		B: 10,
		C: 32.0,
		D: 100.0,
		E: [32]byte{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		},
		F: struct{ G byte }{
			G: 255,
		},
		H: make([]uint64, 256),
	}
	for range itemCount {
		if err := gob.NewEncoder(data).Encode(template); err != nil {
			panic(err)
		}
	}

	seeker := bytes.NewReader(data.Bytes())

	b.ResetTimer()

	for range b.N {
		seeker.Reset(data.Bytes())

		for range itemCount {
			var template encodeStruct
			if err := gob.NewDecoder(seeker).Decode(&template); err != nil {
				panic(err)
			}
			if template.B != 10 {
				b.Errorf("Field B %d is not equal to 10", template.B)
			}
		}
	}
}
