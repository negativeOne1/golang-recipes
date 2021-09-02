package main

import "testing"

func BenchmarkFCFSSelect(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		a := make(chan string, b.N)
		c := make(chan string, b.N)
		firstComeFirstServedSelect("Hello", a, c)
		close(a)
		close(c)
	}
}

func BenchmarkFCFSGoroutines(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		a := make(chan string, b.N)
		c := make(chan string, b.N)
		firstComeFirstServedGoroutines("Hello", a, c)
		close(a)
		close(c)
	}
}

func BenchmarkFCFSSelectVariadic(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		a := make(chan string, b.N)
		c := make(chan string, b.N)
		firstComeFirstServedSelectVariadic("Hello", a, c)
		close(a)
		close(c)
	}
}

func BenchmarkFCFSGoroutinesVariadic(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		a := make(chan string, b.N)
		c := make(chan string, b.N)
		firstComeFirstServedGoroutinesVariadic("Hello", a, c)
		close(a)
		close(c)
	}
}
