package main

import (
	"errors"
	"fmt"
	"hw5/pipe"
	"hw5/wp"
	"sync/atomic"
	"time"
)

func Stage1A(in pipe.In) pipe.Out {
	out := make(pipe.Bi)
	go func() {
		defer func() {
			defer close(out)
			if r := recover(); r != nil {
				out <- r
			}
		}()
		for val := range in {
			time.Sleep(time.Millisecond * 100)
			out <- val.(string) + "A"
		}
	}()

	return out
}

func Stage2B(in pipe.In) pipe.Out {
	out := make(pipe.Bi)
	go func() {
		defer func() {
			defer close(out)
			if r := recover(); r != nil {
				out <- r
			}
		}()
		for val := range in {
			time.Sleep(time.Millisecond * 100)
			out <- val.(string) + "B"
		}
	}()

	return out
}

func Stage3C(in pipe.In) pipe.Out {
	out := make(pipe.Bi)
	go func() {
		defer func() {
			defer close(out)
			if r := recover(); r != nil {
				out <- r
			}
		}()
		for val := range in {
			time.Sleep(time.Millisecond * 100)
			out <- val.(string) + "C"
		}
	}()

	return out
}

var toChange int = 0
var doneWorks int32 = 0

var tChange wp.Task = func() error {
	atomic.AddInt32(&doneWorks, 1)
	return nil
}

var tErr wp.Task = func() error {
	atomic.AddInt32(&doneWorks, 1)
	return errors.New("error sample")
}

var t1 wp.Task = func() error {
	atomic.AddInt32(&doneWorks, 1)
	return nil
}

func main() {
	//wp
	err := wp.Run([]wp.Task{t1, t1, t1, t1, t1, t1, t1, t1, t1, t1}, 5, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(doneWorks)

	//pipe
	done := make(pipe.Bi)
	input := pipe.NewPipeInput("_", "_", "_", "_", "_")
	go func() {
		time.Sleep(time.Second * 100)
		close(done)
	}()
	stages := []pipe.Stage{Stage1A, Stage2B, Stage3C, Stage3C, Stage1A, Stage2B, Stage3C, Stage3C, Stage1A, Stage2B, Stage3C, Stage3C, Stage1A, Stage2B, Stage3C, Stage3C, Stage1A, Stage2B, Stage3C, Stage3C}
	timer := time.Now()
	output := pipe.ExecutePipeline(input, done, stages...)
	for result := range output {
		fmt.Println(result)
	}
	fmt.Println(float32(time.Since(timer).Seconds()))
}

// func send(chNum int, ch chan<- string) {
// 	for i := range 10 {
// 		ch <- "ch " + strconv.Itoa(chNum) + " " + strconv.Itoa(i)
// 	}
// 	close(ch)
// }
