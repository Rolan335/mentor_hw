package wp

import (
	"errors"
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func BenchmarkWp(b *testing.B) {
	for range 1000 {
		err := Run([]Task{t1,t1,t1,t1,t1,t1,t1,t1}, 5, 1)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func BenchmarkWpErr(b *testing.B){
	for range 1000{
		err := Run([]Task{t1,tErr,t1,t1,t1,t1,t1,t1},5,1)
		if err != nil{
			fmt.Println(err)
		}
	}
}

var toChange int = 0
var doneWorks int = 0

var tChange Task = func() error {
	doneWorks++
	toChange = 1
	return nil
}

var tErr Task = func() error {
	doneWorks++
	fmt.Println("got error")
	return errors.New("error sample")
}

var t1 Task = func() error {
	doneWorks++
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(time.Millisecond * 1)
	fmt.Println("task complete")
	return nil
}

func TestWp(t *testing.T) {
	a := assert.New(t)
	tasks := []Task{tChange, t1, t1, t1, t1, t1, t1, t1, t1}
	err := Run(tasks, 5, 1)
	a.Equal(len(tasks), doneWorks, "all tasks done")
	a.NoError(err, "no error")
	a.Equal(1, toChange, "work done and \"toChange\" changed")
	a.Equal(2, runtime.NumGoroutine(), "goroutines stopped") // when testing 2 goroutines at start
}

func TestWpErrStop(t *testing.T) {
	a := assert.New(t)
	err := Run([]Task{tChange, tErr, t1, t1, t1, t1, t1, t1, t1}, 5, 1)
	a.Equal(err, ErrErrorsLimitExceeded, "error returned")
	a.Equal(2, doneWorks)
}