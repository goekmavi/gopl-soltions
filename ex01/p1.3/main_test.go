// go test -bench=.
// go test -bench=WithRange
// go test -bench=^BenchmarkWith
// go test -bench='(Range|Join)'
package main

import (
	"testing"
)

var testArgs = []string{"Hello", "World", "Benchmark"}

func BenchmarkWithRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withRange(testArgs)
	}
}

func BenchmarkWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withJoin(testArgs)
	}
}
