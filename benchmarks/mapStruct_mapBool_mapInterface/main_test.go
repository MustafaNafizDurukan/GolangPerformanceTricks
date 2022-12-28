package main

import (
	"testing"
)

const it = uint(1 << 21)

func BenchmarkSetWithBoolValueWrite(b *testing.B) {
	m := make(map[uint]bool)

	for i := uint(0); i < it; i++ {
		m[i] = true
	}
}

func BenchmarkSetWithStructValueWrite(b *testing.B) {
	m := make(map[uint]struct{})

	for i := uint(0); i < it; i++ {
		m[i] = struct{}{}
	}
}

func BenchmarkSetWithInterfaceValueWrite(b *testing.B) {
	m := make(map[uint]interface{})

	for i := uint(0); i < it; i++ {
		m[i] = struct{}{}
	}
}

/*
BenchmarkSetWithBoolValueWrite-8                1000000000               0.2552 ns/op          0 B/op          0 allocs/op
BenchmarkSetWithStructValueWrite-8              1000000000               0.2419 ns/op          0 B/op          0 allocs/op
BenchmarkSetWithInterfaceValueWrite-8           1000000000               0.3546 ns/op          0 B/op          0 allocs/op
*/
