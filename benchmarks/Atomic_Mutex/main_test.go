package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkAtomic(b *testing.B) {
	var x int64
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&x, 1)
	}
}

func BenchmarkMutex(b *testing.B) {
	var x int64
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		mu.Lock()
		x++
		mu.Unlock()
	}
}

/*
BenchmarkAtomic-8       1000000000               4.610 ns/op           0 B/op          0 allocs/op
BenchmarkMutex-8        548132620               11.02 ns/op            0 B/op          0 allocs/op
*/
