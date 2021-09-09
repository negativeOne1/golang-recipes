package main

import (
	"context"
	"testing"
	"time"
)

func BenchmarkToChansTimedContextGoroutines(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		a := make(chan string, b.N)
		c := make(chan string, b.N)
		ctx := context.Background()
		ToChansTimedContextGoroutines(ctx, time.Millisecond, "Hello", a, c)
		close(a)
		close(c)
	}
}

func BenchmarkToChansTimedTimerSelect(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		a := make(chan string, b.N)
		c := make(chan string, b.N)
		ToChansTimedTimerSelect(time.Millisecond, "Hello", a, c)
		close(a)
		close(c)
	}
}
