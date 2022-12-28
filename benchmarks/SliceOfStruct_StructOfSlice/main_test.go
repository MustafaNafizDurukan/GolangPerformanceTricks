package main

import (
	"testing"
)

const (
	count = 10000000
)

type User struct {
	Name  string
	Email string
}

type SliceOfStructs struct {
	Users []User
}

type StructOfSlices struct {
	Names  []string
	Emails []string
}

func BenchmarkSliceOfStructs(b *testing.B) {
	s := SliceOfStructs{
		Users: make([]User, count),
	}
	for i := 0; i < count; i++ {
		s.Users[i] = User{Name: "John", Email: "john@example.com"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Users[i%count]
	}
}

func BenchmarkStructOfSlices(b *testing.B) {
	s := StructOfSlices{
		Names:  make([]string, count),
		Emails: make([]string, count),
	}
	for i := 0; i < count; i++ {
		s.Names[i] = "John"
		s.Emails[i] = "john@example.com"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = User{
			Name:  s.Names[i%count],
			Email: s.Emails[i%count],
		}
	}
}

/*
BenchmarkSliceOfStructs-8       1000000000               0.8914 ns/op          0 B/op          0 allocs/op
BenchmarkStructOfSlices-8       836656719                1.338 ns/op           0 B/op          0 allocs/op
*/
