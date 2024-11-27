package main

import (
	"flag"
	"fmt"
	"hw4/pi"
)

func main() {
	input := flag.Int("goroutines", 1, "enter num of goroutines")
	flag.Parse()

	// fmt.Println(pi.Pi(*input))

    fmt.Println(pi.Pi2(*input))
}
