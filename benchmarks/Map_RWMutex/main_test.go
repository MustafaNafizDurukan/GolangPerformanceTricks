package main

import (
	"sync"
	"testing"
)

type User struct {
	Name  string
	Email string
}

func BenchmarkMap(b *testing.B) {
	var m sync.Map
	m.Store("john", &User{Name: "John", Email: "john@example.com"})
	m.Store("jane", &User{Name: "Jane", Email: "jane@example.com"})
	m.Store("bob", &User{Name: "Bob", Email: "bob@example.com"})

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = m.Load("john")
			_, _ = m.Load("jane")
			_, _ = m.Load("bob")
			m.Store("john", &User{Name: "John", Email: "john@example.com"})
			m.Delete("jane")
		}
	})
}

func BenchmarkMutex(b *testing.B) {
	m := map[string]*User{
		"john": {Name: "John", Email: "john@example.com"},
		"jane": {Name: "Jane", Email: "jane@example.com"},
		"bob":  {Name: "Bob", Email: "bob@example.com"},
	}

	b.ResetTimer()

	var mu sync.RWMutex
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			_ = m["john"]
			_ = m["jane"]
			_ = m["bob"]

			m["john"] = &User{Name: "John", Email: "john@example.com"}
			delete(m, "jane")

			mu.Unlock()
		}
	})
}

/*
BenchmarkMap-8          122814218              101.0 ns/op            48 B/op          2 allocs/op
BenchmarkMutex-8        75717422               160.4 ns/op            32 B/op          1 allocs/op
*/
