package main

import (
	"errors"
	"fmt"
	"hw5/wp"
	"runtime"
	"sync/atomic"
	"time"
)

var toChange int = 0

var tChange wp.Task = func() error {
	toChange = 1
	return nil
}

var tErr wp.Task = func() error {
	fmt.Println("got error")
	return errors.New("error sample")
}

var t1 wp.Task = func() error {
	fmt.Println("task work")
	return nil
}

func main() {
	fmt.Println(runtime.NumGoroutine())
	var errC int32 = 0
	for range 10000 {
		err := wp.Run([]wp.Task{t1, t1, t1, t1, t1, t1, t1, t1, t1, t1, t1, t1, t1, t1, t1, t1, t1, t1, tErr}, 5, 1)
		if err != nil {
			atomic.AddInt32(&errC, 1)
		}
	}

	time.Sleep(time.Second * 3)
	fmt.Println(runtime.NumGoroutine(), errC)
}
