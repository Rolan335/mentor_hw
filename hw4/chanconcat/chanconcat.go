package chanconcat

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	resCh := make(chan interface{})
	if len(channels) == 0 {
		close(resCh)
		return resCh
	}
	ctx, cancel := context.WithCancel(context.Background())
	for _, ch := range channels {
		go func() {
			select {
			case <-ctx.Done():
				return
			case _, ok := <-ch:
				if !ok {
					close(resCh)
					cancel()
				}
			}
		}()
	}
	return resCh
}

// Пример использования функции:

func Concat() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	fmt.Println(runtime.NumGoroutine())
	go func() {
		time.Sleep(time.Second)
		fmt.Println(runtime.NumGoroutine())
	}()
	start := time.Now()
	_, ok := <-or(
		sig(2*time.Second),
		sig(2*time.Second),
		sig(2*time.Second),
		sig(2*time.Second),
		sig(2*time.Second),
		sig(2*time.Second),
		sig(2*time.Second),
		sig(10*time.Second),
		sig(10*time.Second),
		sig(10*time.Second),
		sig(10*time.Second),
		sig(10*time.Second),
		sig(10*time.Second),
		sig(10*time.Second),
		sig(10*time.Second),
		sig(10*time.Second),
	)
	fmt.Printf("done after %v \n", time.Since(start)) // ~2 second
	fmt.Println("\"or\" open: ", ok)
	time.Sleep(time.Second)
	fmt.Println(runtime.NumGoroutine())
}
