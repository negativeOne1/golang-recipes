package main

//Writing a timesensitive message to a channel using context or timer

import (
	"context"
	"fmt"
	"time"
)

func toChanTimedContext(ctx context.Context, d time.Duration, message string, c chan<- string) (written bool) {
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()

	select {
	case c <- message:
		return true
	case <-ctx.Done():
		return false
	}
}

func toChanTimedTimer(d time.Duration, message string, c chan<- string) (written bool) {
	t := time.NewTimer(d)
	defer t.Stop()

	select {
	case c <- message:
		return true
	case <-t.C:
		return false
	}
}

func mainTimed() {
	c := make(chan string, 1)

	ctx := context.Background()
	println(toChanTimedContext(ctx, time.Millisecond, "SendViaContext", c))

	println(toChanTimedTimer(time.Millisecond, "SendViaTimer", c))

	close(c)

	for s := range c {
		fmt.Println(s)
	}
}
