package main

import (
	"testing"
)

func BenchmarkIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i == 0 {
			_ = 0
		} else if i == 1 {
			_ = 1
		} else if i == 2 {
			_ = 2
		} else if i == 3 {
			_ = 3
		}
	}
}

func BenchmarkSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch i {
		case 0:
			_ = 0
		case 1:
			_ = 1
		case 2:
			_ = 2
		case 3:
			_ = 3
		}
	}
}

/*
BenchmarkIf-8           1000000000               0.2813 ns/op          0 B/op          0 allocs/op
BenchmarkSwitch-8       1000000000               0.2631 ns/op          0 B/op          0 allocs/op
*/
