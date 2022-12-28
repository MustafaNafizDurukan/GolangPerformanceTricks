package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func BenchmarkUnbufferedFileWrite(b *testing.B) {
	file, err := os.Create("unbuffered.test")
	if err != nil {
		b.Fatalf("Unable to create file: %v", err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	for i := 0; i < b.N; i++ {
		fmt.Fprintln(file, "Hello world")
	}
}

func BenchmarkBufferedFileWrite(b *testing.B) {
	file, err := os.Create("buffered.test")
	if err != nil {
		b.Fatalf("Unable to create file: %v", err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < b.N; i++ {
		fmt.Fprintln(writer, "Hello world")
	}
}

/*
BenchmarkUnbufferedFileWrite-8           1928764              3149 ns/op               0 B/op          0 allocs/op
BenchmarkBufferedFileWrite-8            87643570                66.21 ns/op            0 B/op          0 allocs/op
*/
