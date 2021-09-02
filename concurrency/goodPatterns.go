package main

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

func ToChansTimedTimerSelect(d time.Duration, message string, a, b chan string) (written int) {
	t := time.NewTimer(d)
	for i := 0; i < 2; i++ {
		select {
		case a <- message:
			a = nil
		case b <- message:
			b = nil
		case <-t.C:
			return i
		}
	}
	t.Stop()
	return 2
}

func ToChansTimedContextGoroutines(ctx context.Context, d time.Duration, message string, ch ...chan string) (written int) {
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()
	var (
		wr int32
		wg sync.WaitGroup
	)
	wg.Add(len(ch))
	for _, c := range ch {
		c := c
		go func() {
			defer wg.Done()
			select {
			case c <- message:
				atomic.AddInt32(&wr, 1)
			case <-ctx.Done():
			}
		}()
	}
	wg.Wait()
	return int(wr)
}
