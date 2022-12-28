package main

import (
	"encoding/json"
	"testing"
)

type User struct {
	Name  string
	Email string
}

func BenchmarkStruct(b *testing.B) {
	u := User{Name: "John", Email: "john@example.com"}
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(u)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMap(b *testing.B) {
	m := map[string]string{"Name": "John", "Email": "john@example.com"}
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(m)
		if err != nil {
			b.Error(err)
		}
	}
}

/*
BenchmarkStruct-8       22557823               236.3 ns/op            80 B/op          2 allocs/op
BenchmarkMap-8           7919541               768.2 ns/op           376 B/op          9 allocs/op
*/
