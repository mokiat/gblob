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

Following are some benchmark results. They compare this library against Go's `binary` and `gob` packages, since those are closest in terms of features. Results are based on the following hardware:

```
goos: linux
goarch: amd64
pkg: github.com/mokiat/gblob
cpu: AMD Ryzen 7 3700X 8-Core Processor
```

You can find the benchmark tests in the `bench_test.go` file. You can run them with the following command:

```sh
go test -benchmem -run=^$ -bench ^Benchmark github.com/mokiat/gblob
```
Adjust the `^Benchmark` regex to focus on specific benchmark sets and use the `-cpuprofile bench.pprof` flag if profiling is desired.

Take the results that follow with a grain of salt. Do your own measurements for your particular use case.


### Block API

Following is a benchmark comparison between `Block` and Go's `binary.ByteOrder` API.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | -----------------: | -----------------------------: | ------------------: |
| `Block` | 1493 ns/op | 0 B/op | 0 allocs/op |
| `binary.ByteOrder` | 1496 ns/op | 0 B/op | 0 allocs/op |

> Both APIs provide similar performance so it boils down to ease of use of and personal preference.


### TypedWriter

Following is a benchmark comparison between `TypedWriter` and Go's `binary.Write` functions.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | -----------------: | -----------------------------: | ------------------: |
| `TypedWriter` | 7753 ns/op | 0 B/op | 0 allocs/op |
| `binary.Write` | 39768 ns/op | 4096 B/op | 1024 allocs/op |

> The `TypedWriter` does not allocate any memory per write operation and it also runs about `3-4 times` faster. In reality, it allocates an initial buffer of size `8 bytes` that it reuses to achieve these results.


### TypedReader

Following is a benchmark comparison between `TypedReader` and Go's `binary.Read` functions.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | -----------------: | -----------------------------: | ------------------: |
| `TypedReader` | 11595 ns/op | 0 B/op | 0 allocs/op |
| `binary.Read` | 53088 ns/op | 4096 B/op | 1024 allocs/op |

> The `TypedReader` does not allocate any memory per read and it also runs `4-5 times` faster. This is again achieved by having the `TypedReader` allocate an initial buffer of `8 bytes` that it reuses.


### PackedEncoder

When serializing larger types with `PackedEncoder` vs `binary.Write`, the performance difference is negligible, though the memory aspect remains the same as before. More interesting in this case is the comparison with `gob.Encoder`, due to a similar feature set.

**WARNING:** The scenario that is compared is the one where a new instance of an `Encoder` is created for each `Encode` performed (i.e. `gob.NewEncoder(out).Encode(&target)`). This is arguably the more common scenario (saving an asset / writing a response). When a `gob.Encoder` is reused to encode multiple sequences of items to a stream, it is significantly faster than `PackedEncoder`, likely due to caching.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | -----------------: | -----------------------------: | ------------------: |
| `PackedEncoder` | 7644113 ns/op | 179158 B/op | 3072 allocs/op |
| `gob.Encoder` | 11349161 ns/op | 2408841 B/op | 38912 allocs/op |

> Here the `gob.Encoder` performs worse, especially when memory is concerned.


### PackedDecoder

Following is the comparison between `PackedDecoder` and `gob.Decoder`.

**WARNING:** The scenario that is compared is the one where a new instance of a `Decoder` is created for each `Decode` performed (i.e. `gob.NewDecoder(in).Decode(&target)`). This is arguably the more common scenario (loading an asset / reading a request). When a `gob.Decoder` is reused to read multiple sequences of items from a stream, it is significantly faster than `PackedDecoder`, likely due to caching.

| Approach | Time per Operation | Allocated Memory per Operation | Allocation Count per Operation |
| -------- | -----------------: | -----------------------------: | ------------------: |
| `PackedDecoder` | 17483112 ns/op | 2261422 B/op | 5120 allocs/op |
| `gob.Decoder` | 45206563 ns/op | 11466272 B/op | 236565 allocs/op |

> The `gob.Decoder` performs werse, especially when memory is concerned.
