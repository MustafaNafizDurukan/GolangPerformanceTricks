package main

import (
	"os"
	"strconv"
	"testing"
)

const (
	count = 10000000
)

var (
	stringMap = make(map[string]int)
	intMap    = make(map[int]int)
)

func TestMain(m *testing.M) {
	for i := 0; i < count; i++ {
		stringMap[strconv.Itoa(i)] = i
	}

	for i := 0; i < count; i++ {
		intMap[i] = i
	}

	code := m.Run()

	os.Exit(code)
}

func BenchmarkFindIntKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, ok := intMap[i]
		if !ok {
		}
	}
}

func BenchmarkFindStringKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, ok := stringMap[strconv.Itoa(i)]
		if !ok {
		}
	}
}

func BenchmarkWriteIntKey(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < count; i++ {
		m[i] = i
	}
}

func BenchmarkWriteStringKey(b *testing.B) {
	m := make(map[string]int)
	for i := 0; i < count; i++ {
		m[strconv.Itoa(i)] = i
	}
}

/*
BenchmarkFindIntKey-8           100000000               55.13 ns/op            0 B/op          0 allocs/op
BenchmarkFindStringKey-8        47351510               145.0 ns/op             7 B/op          0 allocs/op
BenchmarkInsertIntKey-8         1000000000               0.6194 ns/op          0 B/op          0 allocs/op
BenchmarkInsertStringKey-8      1000000000               1.458 ns/op           0 B/op          0 allocs/op
*/
