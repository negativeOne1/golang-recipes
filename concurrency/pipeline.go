package main

import (
	"fmt"
	"sync"
)

// Pipelines as a series of stages connected by channels

// gen converts a list of numbers to a channel
func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// sq receives integers from channel and returns a channel that emits the
// square of each received
func sq(in <-chan int) <-chan int {
	out := make(chan int, 1)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func simplePipeline() {
	c := gen(2, 3)
	out := sq(c)

	fmt.Println(<-out) //4
	fmt.Println(<-out) //9

	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)
	}
}

func withMerge() {
	in := gen(2, 3)

	c1 := sq(in)
	c2 := sq(in)

	fmt.Println(<-merge(c1, c2))
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	simplePipeline()
	withMerge()
}
