# Go Blob ![Build Status](https://github.com/mokiat/gblob/workflows/Go/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/mokiat/gblob)](https://goreportcard.com/report/github.com/mokiat/gblob) [![GoDoc](https://godoc.org/github.com/mokiat/gblob?status.svg)](https://godoc.org/github.com/mokiat/gblob)

A package that provides utilities for writing and reading primitive types to and from byte sequences.

## User's Guide

For a more complete documentation, check the [GoDocs](https://pkg.go.dev/github.com/mokiat/gblob).

### Block

The **Block** API allows one to place values of concrete primitive types at specific offsets inside a byte slice.

**Example:**

```go
block := make(gblob.LittleEndianBlock, 32)
block.SetFloat32(4, 3.5)
```

There are two implementations available - **LittleEndianBlock** and **BigEndianBlock**.

Following is a benchmark comparison between `Block` and Go's `binary.ByteOrder` API.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | ------------------ | ------------------------------ | ------------------- |
| `Block` | 0.2327 ns/op | 0 B/op | 0 allocs/op |
| `binary.ByteOrder` | 0.2281 ns/op | 0 B/op | 0 allocs/op |

Both APIs provide similar performance.

> As with any benchmark result, take it with a grain of salt and do your own measurements according to your own requirements.


### TypedWriter

The **TypedWriter** API allows one to write concrete primitive types to an `io.Writer`.

**Example:**

```go
var buffer bytes.Buffer
writer := gblob.NewLittleEndianWriter(&buffer)
writer.WriteUint64(0x13743521FA954321)
```

There are two implementations available - **NewLittleEndianWriter** and **NewBigEndianWriter**.

Following is a benchmark comparison between `TypedWriter` and Go's `binary.Write` functions.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | ------------------ | ------------------------------ | ------------------- |
| `TypedWriter` | 56.97 ns/op | 0 B/op | 0 allocs/op |
| `binary.Write` | 175.5 ns/op | 32 B/op | 6 allocs/op |

As can be seen, not only does the TypedWriter not allocate any memory, but it also runs about `3 times` faster.

> As with any benchmark result, take it with a grain of salt and do your own measurements according to your own requirements.

**NOTE:** In reality the `TypedWriter` does allocate around `8 bytes` but only when first created (uses it for a write buffer). It does not allocate during write operations.

### TypedReader

The **TypedReader** API allows one to read concrete primitive types from an `io.Reader`.

**Example:**

```go
buffer := bytes.NewBuffer([]uint8{
  0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20,
})
reader := gblob.NewLittleEndianReader(buffer)
value, err := reader.ReadUint64()
```

There are two implementations available - **NewLittleEndianReader** and **NewBigEndianReader**.

Following is a benchmark comparison between `TypedReader` and Go's `binary.Read` functions.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | ------------------ | ------------------------------ | ------------------- |
| `TypedReader` | 61.69 ns/op | 0 B/op | 0 allocs/op |
| `binary.Read` | 226.3 ns/op | 56 B/op | 12 allocs/op |

As can be seen, not only does the TypedReader not allocate any memory, but it also runs nearly `4 times` faster.

> As with any benchmark result, take it with a grain of salt and do your own measurements according to your own requirements.

**NOTE:** In reality the `TypedReader` does allocate around `8 bytes` but only when first created (uses it for a read buffer). It does not allocate during read operations.

### PackedEncoder

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

There are two implementations available - **NewLittleEndianPackedEncoder** and **NewBigEndianPackedEncoder**.

It is difficult to compare the `PackedEncoder` API with `binary.Write` or `gob.NewEncoder`, since the former has less
flexibility to the types it can read and the latter is much more flexible - it allows for forward compatible
serialization.

Nevertheless, following is a benchmark comparison between `PackedEncoder` and Go's `gob.NewEncoder` functions.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | ------------------ | ------------------------------ | ------------------- |
| `PackedEncoder` | 242.6 ns/op | 72 B/op | 2 allocs/op |
| `gob.NewEncoder` | 408.7 ns/op | 72 B/op | 2 allocs/op |

The `PackedEncoder` is almost twice faster. Memory allocation is equal.

And following is a benchmark comparison between `PackedEncoder` and Go's `binary.Write` functions.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | ------------------ | ------------------------------ | ------------------- |
| `PackedEncoder` | 133.7 ns/op | 32 B/op | 1 allocs/op |
| `binary.Write` | 442.9 ns/op | 112 B/op | 8 allocs/op |

The `PackedEncoder` is significantly quicker and it allocates less memory.

> As with any benchmark result, take it with a grain of salt and do your own measurements according to your own requirements.

### PackedDecoder

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

There are two implementations available - **NewLittleEndianPackedDecoder** and **NewBigEndianPackedDecoder**.

It is difficult to compare the `PackedDecoder` API with `binary.Read` or `gob.NewDecoder`, since the former has less
flexibility to the types it can read and the latter is much more flexible - it allows for backward compatible
deserialization.

Nevertheless, following is a benchmark comparison between `PackedDecoder` and Go's `gob.NewDecoder` functions.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | ------------------ | ------------------------------ | ------------------- |
| `PackedDecoder` | 434.6 ns/op | 152 B/op | 5 allocs/op |
| `gob.NewDecoder` | 19480 ns/op | 7928 B/op | 212 allocs/op |

As can be seen, the `PackedDecoder` is much faster and barely allocates memory. On the downside, a change to the target
type would break decoding, which is not the case with `gob.NewDecoder`.

And following is a benchmark comparison between `PackedDecoder` and Go's `binary.Read` functions.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | ------------------ | ------------------------------ | ------------------- |
| `PackedDecoder` | 187.7 ns/op | 32 B/op | 1 allocs/op |
| `binary.Read` | 173.5 ns/op | 64 B/op | 2 allocs/op |

The `PackedDecoder` is marginably slower, though it allocates slightly less memory and has the added benefit of
supporing types like slice and map.

> As with any benchmark result, take it with a grain of salt and do your own measurements according to your own requirements.
