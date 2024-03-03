# Go Blob

[![Go Reference](https://pkg.go.dev/badge/github.com/mokiat/gblob.svg)](https://pkg.go.dev/github.com/mokiat/gblob)
![Build Status](https://github.com/mokiat/gblob/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/mokiat/gblob)](https://goreportcard.com/report/github.com/mokiat/gblob)

A package that provides utilities for writing and reading primitive types to and from byte sequences.

## User's Guide

For a more complete documentation, check the [GoDocs](https://pkg.go.dev/github.com/mokiat/gblob).

### Block API

The **Block** API allows one to place values of concrete primitive types at specific offsets inside a byte slice. It is up to the caller to ensure that the slice has the necessary size.

**Example:**

```go
offset := 4
block := make(gblob.LittleEndianBlock, 32)
block.SetFloat32(offset, 3.5)
fmt.Println(block.Float32(offset))
```

This API is similar to Go's built-in `binary.ByteOrder`. The difference here is that the gblob API is slightly more compact, it has helper functions for more primitive types and allows one to pass an offset into the byte slice for better readability.

There are two implementations available - **LittleEndianBlock** and **BigEndianBlock**, depending on the desired byte order.

### TypedWriter / TypedReader API

The **TypedWriter** API allows one to write concrete primitive types to an `io.Writer`. No padding or alignment is added to the output.

**Example:**

```go
var buffer bytes.Buffer
writer := gblob.NewLittleEndianWriter(&buffer)
writer.WriteUint64(0x13743521FA954321)
```

There are two implementations available - **NewLittleEndianWriter** and **NewBigEndianWriter**, depending on the desired byte order.

The **TypedReader** API allows one to read concrete primitive types from an `io.Reader`.

**Example:**

```go
buffer := bytes.NewBuffer([]uint8{
  0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20,
})
reader := gblob.NewLittleEndianReader(buffer)
value, err := reader.ReadUint64()
```

There are two implementations available - **NewLittleEndianReader** and **NewBigEndianReader**, depending on the desired byte order.

### PackedEncoder / PackedDecoder API

The **PackedEncoder** API allows one to marshal a data structure to a binary
sequence.

**Example:**

```go
var buffer bytes.Buffer
gblob.NewLittleEndianPackedEncoder(&buffer).Encode(struct{
  A float32
  B uint64
}{
  A: 3.14,
  B: 0xFF003344FF003344,
})
```

There are two implementations available - **NewLittleEndianPackedEncoder** and **NewBigEndianPackedEncoder**, depending on the desired byte order.

This is similar to Go's `bytes.Write`, except that it supports slices, maps and strings.

The **PackedDecoder** API allows one to unmarshal a data structure from a
binary sequence that was previously marshaled through the PackedEncoder API.

**Example:**

```go
buffer := bytes.NewBuffer([]uint8{
  0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20,
})
var target uint64
gblob.NewLittleEndianPackedDecoder(buffer).Decode(&target)
```

There are two implementations available - **NewLittleEndianPackedDecoder** and **NewBigEndianPackedDecoder**, depending on the desired byte order.

This is similar to Go's `bytes.Read`, except that it supports slices, maps and strings.

## Performance

Following are some benchmark results. They compare this library against Go's `binary` and `gob`, since those are closest in terms of features.

Nevertheless, there is not a complete feature parity between the packages, hence the `glob` library is not a direct replacement nor are the test exhaustive.

Take the results that follow with a grain of salt. Do your own measurements for your particular use case if you consider using the library.

### Block API

Following is a benchmark comparison between `Block` and Go's `binary.ByteOrder` API.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | ------------------ | ------------------------------ | ------------------- |
| `Block` | 0.2327 ns/op | 0 B/op | 0 allocs/op |
| `binary.ByteOrder` | 0.2281 ns/op | 0 B/op | 0 allocs/op |

> Both APIs provide similar performance so it boils down to ease of use of and personal preference.

### TypedWriter

Following is a benchmark comparison between `TypedWriter` and Go's `binary.Write` functions.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | ------------------ | ------------------------------ | ------------------- |
| `TypedWriter` | 56.97 ns/op | 0 B/op | 0 allocs/op |
| `binary.Write` | 175.5 ns/op | 32 B/op | 6 allocs/op |

> The `TypedWriter` does not allocate any memory per write operation and it also runs about `3 times` faster. In reality, it allocates an initial buffer of size `8 bytes` that it reuses to achieve thes results.

### TypedReader

Following is a benchmark comparison between `TypedReader` and Go's `binary.Read` functions.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | ------------------ | ------------------------------ | ------------------- |
| `TypedReader` | 61.69 ns/op | 0 B/op | 0 allocs/op |
| `binary.Read` | 226.3 ns/op | 56 B/op | 12 allocs/op |

> The `TypedReader` does not allocate any memory per read and it also runs nearly `4 times` faster. This is again achieved by having the `TypedReader` allocate an initial buffer of `8 bytes` that it reuses.

### PackedEncoder

It is difficult to compare the `PackedEncoder` API with `binary.Write` or `gob.NewEncoder`, since the former has less flexibility to the types it can read and the latter is much more flexible - it allows for forward compatible serialization.

Nevertheless, following is a benchmark comparison between `PackedEncoder` and Go's `gob.Encoder`.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | ------------------ | ------------------------------ | ------------------- |
| `PackedEncoder` | 242.6 ns/op | 72 B/op | 2 allocs/op |
| `gob.Encoder` | 408.7 ns/op | 72 B/op | 2 allocs/op |

> `PackedEncoder` is almost twice faster. Memory allocation is equal.

And following is a benchmark comparison between `PackedEncoder` and Go's `binary.Write` function.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | ------------------ | ------------------------------ | ------------------- |
| `PackedEncoder` | 133.7 ns/op | 32 B/op | 1 allocs/op |
| `binary.Write` | 442.9 ns/op | 112 B/op | 8 allocs/op |

> The `PackedEncoder` is significantly quicker and it allocates less memory.

### PackedDecoder

It is difficult to compare the `PackedDecoder` API with `binary.Read` or `gob.Decoder`, since the former has less flexibility to the types it can read and the latter is much more flexible - it allows for backward compatible deserialization.

Nevertheless, following is a benchmark comparison between `PackedDecoder` and Go's `gob.Decoder`.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | ------------------ | ------------------------------ | ------------------- |
| `PackedDecoder` | 434.6 ns/op | 152 B/op | 5 allocs/op |
| `gob.Decoder` | 19480 ns/op | 7928 B/op | 212 allocs/op |

> As can be seen, the `PackedDecoder` is much faster and barely allocates memory. On the downside, a change to the target type would break decoding, which is not the case with `gob.Decoder`.

And following is a benchmark comparison between `PackedDecoder` and Go's `binary.Read` function.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | ------------------ | ------------------------------ | ------------------- |
| `PackedDecoder` | 187.7 ns/op | 32 B/op | 1 allocs/op |
| `binary.Read` | 173.5 ns/op | 64 B/op | 2 allocs/op |

> The `PackedDecoder` is marginably slower, though it allocates slightly less memory and has the added benefit of supporing types like slice and map.
