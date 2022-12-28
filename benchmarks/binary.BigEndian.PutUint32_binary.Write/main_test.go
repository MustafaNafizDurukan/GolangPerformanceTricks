package main

import (
	"bytes"
	"encoding/binary"
	"testing"
)

var (
	testData = uint32(1234567890)
	buf      = new(bytes.Buffer)
)

func BenchmarkBigEndianPutUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf.Reset()
		var data [4]byte
		binary.BigEndian.PutUint32(data[:], testData)
		buf.Write(data[:])
	}
}

func BenchmarkWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf.Reset()
		binary.Write(buf, binary.BigEndian, testData)
	}
}

/*
BenchmarkBigEndianPutUint32-8           1000000000               4.947 ns/op           0 B/op          0 allocs/op
BenchmarkWrite-8                        173993296               34.48 ns/op            8 B/op          2 allocs/op
*/
