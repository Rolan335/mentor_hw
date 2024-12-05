package main

import (
	"context"
	"fmt"
	"hw4/semaphore"
	"time"
)

func main() {
	semaphore := semaphore.NewSemaphoreCond(2)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	timer := time.Now()
	done := 0
	for i := range 10 {
		err := semaphore.Acquire(ctx, 1)
		fmt.Println(err)
		if err == nil {
			go func() {
				defer semaphore.Release(1)
				fmt.Println("goroutine ", i, " start")
				time.Sleep(time.Second)
				fmt.Println("goroutine ", i, " stop")
				done++
			}()
		}
	}
	for done != 10 {

	}
	fmt.Println(time.Since(timer))
	for {

	}
}
