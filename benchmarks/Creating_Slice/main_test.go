package main

import (
	"math/rand"
	"testing"
)

func BenchmarkMakeSliceWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := make([]int, 0, 1024)
		for i := 0; i < 1024; i++ {
			numbers = append(numbers, rand.Int())
		}
	}
}

func BenchmarkMakeSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := []int{}
		for i := 0; i < 1024; i++ {
			numbers = append(numbers, rand.Int())
		}
	}
}

/*
BenchmarkBigEndianPutUint32-8           1000000000               4.966 ns/op           0 B/op          0 allocs/op
BenchmarkWrite-8                        179831037               33.50 ns/op            8 B/op          2 allocs/op
*/
