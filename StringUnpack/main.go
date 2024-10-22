package main

import (
	"StringUnpack/unpack"
	"fmt"
)

func main(){
	//TODO: zero validation and another tasks.
	//`qwe\\\\45`
	input := `qwe45`
	fmt.Println(unpack.Unpack(input,true))
	// fmt.Println(unpack.Unpack(input,true))
}