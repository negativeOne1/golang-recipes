package main

import "sync"

type Once chan struct{}

func NewOnce() Once {
	o := make(Once, 1)

	o <- struct{}{}
	return o
}

func (o Once) Do(f func()) {
	_, ok := <-o
	if !ok {
		return
	}

	f()
	close(o)
}

func main() {
	o := NewOnce()
	wg := sync.WaitGroup{}
	wg.Add(3)
	//will only run once
	go func() { o.Do(func() { println("HELLO") }); wg.Done() }()
	go func() { o.Do(func() { println("HELLO") }); wg.Done() }()
	go func() { o.Do(func() { println("HELLO") }); wg.Done() }()
	wg.Wait()
}
