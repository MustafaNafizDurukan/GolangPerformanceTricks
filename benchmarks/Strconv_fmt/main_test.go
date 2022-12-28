package main

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", 1234567890)
	}
}

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(1234567890)
	}
}

/*
BenchmarkSprintf-8      44634056                82.51 ns/op           16 B/op          1 allocs/op
BenchmarkItoa-8         92875183                35.08 ns/op           16 B/op          1 allocs/op
*/
