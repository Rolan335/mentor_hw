package main

import (
	"fmt"
	l "hw3/lrucache_refactor"
)

func main() {
	cache := l.New(2)
	cache.Set("a", "a")
	fmt.Println(cache)
	cache.Set("b", "b")
	cache.Set("c", "c")
	fmt.Println(cache)
	cache.Set("d", "d")
	fmt.Println(cache)
}
