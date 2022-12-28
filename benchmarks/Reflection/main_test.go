package main

import (
	"reflect"
	"testing"
)

// Avoid using reflection:
// Reflection can be a useful feature in Go, but it can also be a performance bottleneck. Avoid using reflection if possible,
// or use it sparingly in performance-critical sections of your code.

type User struct {
	Name  string
	Email string
}

func BenchmarkReflection(b *testing.B) {
	u := &User{}
	val := reflect.ValueOf(u).Elem()
	field := val.FieldByName("Name")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field.SetString("John")
	}
}

func BenchmarkNoReflection(b *testing.B) {
	u := &User{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.Name = "John"
	}
}
