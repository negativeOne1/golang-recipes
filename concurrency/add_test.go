package main

import (
	"math/rand"
	"runtime"
	"testing"
)

var (
	thousand = 1000
	million  = thousand * thousand
	numbers  = rand.Perm(10 * million)
)

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(numbers)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addConcurrent(runtime.NumCPU(), numbers)
	}
}
