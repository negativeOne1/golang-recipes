package main

import (
	"reflect"
	"sync"
)

//go:noinline
func firstComeFirstServedSelect(message string, a, b chan<- string) {
	for i := 0; i < 2; i++ {
		select {
		case a <- message:
			a = nil
		case b <- message:
			b = nil
		}
	}
}

//go:noinline
func firstComeFirstServedGoroutines(message string, a, b chan<- string) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() { a <- message; wg.Done() }()
	go func() { b <- message; wg.Done() }()
	wg.Wait()
}

//go:noinline
func firstComeFirstServedSelectVariadic(message string, chs ...chan<- string) {
	cases := make([]reflect.SelectCase, len(chs))
	for i, ch := range chs {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: reflect.ValueOf(message),
		}
	}

	for i := 0; i < len(chs); i++ {
		chosen, _, _ := reflect.Select(cases)
		cases[chosen].Chan = reflect.ValueOf(nil)
	}
}

//go:noinline
func firstComeFirstServedGoroutinesVariadic(message string, chs ...chan<- string) {
	var wg sync.WaitGroup
	wg.Add(len(chs))

	for _, c := range chs {
		c := c //we need a new variable for this, otherwise the loop can't proceed
		go func() { c <- message; wg.Done() }()
	}
	wg.Wait()
}

func mainFCFS() {
	a := make(chan string, 2)
	b := make(chan string, 2)

	firstComeFirstServedSelect("Hello", a, b)
	firstComeFirstServedGoroutines("Goroutines", a, b)
	close(a)
	close(b)

	for s := range a {
		println(s)
	}
	for s := range b {
		println(s)
	}
}
