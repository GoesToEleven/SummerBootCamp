package main

import (
	"testing"
	"fmt"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

// run this at command:
// go test -bench='.*'
