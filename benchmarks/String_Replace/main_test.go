package main

import (
	"bytes"
	"strings"
	"testing"
)

var (
	testData = "performanceExamples"
	old      = "e"
	new      = "z"
)

func BenchmarkReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strings.Replace(testData, old, new, -1)
		_ = s
	}
}

func BenchmarkBytesReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := []byte(testData)
		b = bytes.Replace(b, []byte(old), []byte(new), -1)
		_ = string(b)
	}
}

/*
BenchmarkReplace-8              68897886                79.88 ns/op           24 B/op          1 allocs/op
BenchmarkBytesReplace-8         62435028                94.14 ns/op           24 B/op          1 allocs/op
*/
