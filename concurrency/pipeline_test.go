package main

import (
	"math/rand"
	"testing"
)

var (
	pipelineNumbers = rand.Perm(1e2)
)

func BenchmarkPipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := gen(2, 3)

		sq(in)
		sq(in)

	}
}

func BenchmarkPipelineMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := gen(2, 3)

		c1 := sq(in)
		c2 := sq(in)

		merge(c1, c2)
	}
}
