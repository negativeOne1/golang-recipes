package main

import (
	"sync"
	"sync/atomic"
)

func add(numbers []int) int {
	var v int
	for _, n := range numbers {
		v += n
	}
	return v
}

func addConcurrent(goroutines int, numbers []int) int {
	var v int64
	totalNumbers := len(numbers)
	lastGoroutine := goroutines - 1
	stride := totalNumbers / goroutines

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func(g int) {
			start := g * stride
			end := start + stride

			if g == lastGoroutine {
				end = totalNumbers
			}

			atomic.AddInt64(&v, int64(add(numbers[start:end])))
			wg.Done()
		}(g)
	}

	wg.Wait()
	return int(v)
}
