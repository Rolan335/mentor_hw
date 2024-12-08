package main

import (
	"fmt"
	"hw5/wp"
	"strings"
	"time"
)

type multiply struct {
	a, b int
}

func (m multiply) Do() (any, error) {
	time.Sleep(time.Millisecond * 200)
	return m.a * m.b, nil
}

type sum struct {
	a, b int
}

func (s sum) Do() (any, error) {
	time.Sleep(time.Millisecond * 100)
	return s.a + s.b, nil
}

type concat struct {
	a, b string
}

func (c concat) Do() (any, error) {
	time.Sleep(time.Millisecond * 7)
	newStr := make([]string, 0, len(c.b)+1)
	newStr = append(newStr, c.a)
	newStr = append(newStr, c.b)
	return strings.Join(newStr, ""), nil
}

func main() {
	pool := wp.NewWorkerPool(5, 5)
	res, err := pool.ProceedTasks(
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
		sum{2, 3},
		sum{4, 5},
		multiply{3, 5},
		concat{"ff", "gg"},
		concat{"fss", "cpxcg"},
	)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}
