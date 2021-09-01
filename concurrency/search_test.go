package main

import (
	"runtime"
	"testing"
)

var (
	docs = generateList(1e3)
)

func BenchmarkSearchSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		find("Go", docs)
	}
}

func BenchmarkSearchConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findConcurrent(runtime.NumCPU(), "Go", docs)
	}
}
