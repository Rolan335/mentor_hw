package wp

import (
	"errors"
	"fmt"
	"runtime"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func BenchmarkWp(b *testing.B) {
	for range 1000 {
		err := Run([]Task{t1, t1, t1, t1, t1, t1, t1, t1}, 5, 1)
		if err != nil {
			fmt.Println(err)
		}
	}
}

var errC int32 = 0

func BenchmarkWpErr(b *testing.B) {
	for range 1000 {
		err := Run([]Task{t1, t1, t1, t1, t1, t1, t1, tErr}, 10, 1)
		if err != nil {
			atomic.AddInt32(&errC, 1)
		}
	}
	fmt.Println("end", runtime.NumGoroutine(), errC)
}

var toChange int = 0
var doneWorks int32 = 0

var tChange Task = func() error {
	atomic.AddInt32(&doneWorks, 1)
	toChange = 1
	return nil
}

var tErr Task = func() error {
	atomic.AddInt32(&doneWorks, 1)
	return errors.New("error sample")
}

var t1 Task = func() error {
	atomic.AddInt32(&doneWorks, 1)
	return nil
}

func TestWp(t *testing.T) {
	toChange = 0
	atomic.SwapInt32(&doneWorks, 0)
	a := assert.New(t)
	tasks := []Task{tChange, t1, t1, t1, t1, t1, t1, t1, t1}
	err := Run(tasks, 5, 1)
	a.EqualValues(len(tasks), doneWorks, "all tasks done")
	a.NoError(err, "no error")
	a.Equal(1, toChange, "work done and \"toChange\" changed")
	a.Equal(2, runtime.NumGoroutine() < 5, "goroutines stopped")
}

func TestWpErrStop(t *testing.T) {
	toChange = 0
	a := assert.New(t)
	atomic.SwapInt32(&doneWorks, 0)
	err := Run([]Task{tErr, tChange, t1, t1, t1, t1, t1, t1, t1, t1}, 1, 1)
	a.Equal(err, ErrErrorsLimitExceeded, "error returned")
	a.Equal(true, doneWorks < 10, "work stopped")
	atomic.SwapInt32(&doneWorks, 0)
	_ = Run([]Task{tChange, tErr, t1, t1, t1, t1, t1, t1, t1, t1}, 5, -1)
	a.Equal(true, doneWorks >= 10, "works done despite all errrors")
	atomic.SwapInt32(&doneWorks, 0)
	_ = Run([]Task{tErr, tChange, t1, t1, t1, t1, t1, t1, t1, t1}, 5, 0)
	a.Equal(true, doneWorks <= 5, "works hasn't been done")
	time.Sleep(time.Second)
	a.Equal(2, runtime.NumGoroutine(), "goroutines stopped")
}

func TestWpTwoErr(t *testing.T) {
	toChange = 0
	atomic.SwapInt32(&doneWorks, 0)
	a := assert.New(t)
	err := Run([]Task{tChange, tErr, t1, t1, t1, t1, t1, t1, t1, t1}, 5, 2)
	a.NoError(err, "error returned")
	a.EqualValues(10, doneWorks)
}

// func TestConcurrency(t *testing.T) {
// 	a := assert.New(t)
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func() {
// 		_ = Run([]Task{tChange, tErr, t1, t1, t1, t1, t1, t1, t1, t1}, 5, 2)
// 		wg.Done()
// 	}()
//     runtime.Gosched()
// 	goroutines := runtime.NumGoroutine()
// 	wg.Wait()
// 	a.Equal(true, goroutines >= 7, "testing goroutines + workers")
// }
