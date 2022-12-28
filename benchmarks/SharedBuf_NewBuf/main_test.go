package main

import (
	"math/rand"
	"testing"
)

func BenchmarkNewBuffers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1024)
		buf := make([]byte, n)

		// Do something with buffer
		for j := 0; j < n; j++ {
			buf[j] = 1
		}
	}
}

func BenchmarkReuseBuffers(b *testing.B) {
	sharedBuf := make([]byte, 1024)

	for i := 0; i < b.N; i++ {
		n := rand.Intn(1024)
		buf := sharedBuf[0:n]

		// Do something with buffer
		for j := 0; j < n; j++ {
			buf[j] = 1
		}
	}
}

/*
BenchmarkNewBuffers-8           12385022               445.7 ns/op           540 B/op          0 allocs/op
BenchmarkReuseBuffers-8         39079291               155.3 ns/op             0 B/op          0 allocs/op
*/
