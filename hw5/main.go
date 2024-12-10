package main

import (
	"errors"
	"fmt"
	"hw5/wp"
	"runtime"
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
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("task complete")
	return nil
}

func main() {
	runtime.GOMAXPROCS(5)
	fmt.Println(runtime.NumGoroutine())
	err := wp.Run([]wp.Task{t1, t1, t1, t1, t1, t1, t1, t1, t1, tErr, t1, t1, t1, t1, t1, t1, t1, t1, t1, t1, t1, t1, t1, t1, t1}, 10, 1)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(runtime.NumGoroutine())
}

// type multiply struct {
// 	a, b int
// }

// func (m multiply) Do() (any, error) {
// 	time.Sleep(time.Millisecond * 200)
// 	return m.a * m.b, nil
// }

// type sum struct {
// 	a, b int
// }

// func (s sum) Do() (any, error) {
// 	time.Sleep(time.Millisecond * 100)
// 	return s.a + s.b, nil
// }

// type concat struct {
// 	a, b string
// }

// func (c concat) Do() (any, error) {
// 	time.Sleep(time.Millisecond * 7)
// 	newStr := make([]string, 0, len(c.b)+1)
// 	newStr = append(newStr, c.a)
// 	newStr = append(newStr, c.b)
// 	return strings.Join(newStr, ""), nil
// }
