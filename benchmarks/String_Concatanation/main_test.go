package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var (
	testData = []string{"a", "b", "c", "d", "e"}
)

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strings.Join(testData, ":")
		_ = s
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s:%s:%s:%s:%s", testData[0], testData[1], testData[2], testData[3], testData[4])
		_ = s
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := testData[0] + ":" + testData[1] + ":" + testData[2] + ":" + testData[3] + ":" + testData[4]
		_ = s
	}
}

func BenchmarkStringsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for _, s := range testData {
			builder.WriteString(s)
			builder.WriteString(":")
		}
		s := builder.String()
		_ = s
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		for _, s := range testData {
			buffer.WriteString(s)
			buffer.WriteString(":")
		}
		_ = buffer.String()
	}
}

/*
BenchmarkStringsJoin-8          68165312                52.27 ns/op           16 B/op          1 allocs/op
BenchmarkSprintf-8              11581468               305.7 ns/op            96 B/op          6 allocs/op
BenchmarkConcat-8               66846781                51.51 ns/op            0 B/op          0 allocs/op
BenchmarkStringsBuilder-8       48861591                73.24 ns/op           24 B/op          2 allocs/op
BenchmarkBytesBuffer-8          42273943                79.86 ns/op           64 B/op          1 allocs/op
*/
