package main

import (
	"StringUnpack/unpack"
	"fmt"
)

func main(){
	//TODO: zero validation and another tasks.
	//`qwe\\\\45`
	input := "4a4b3cd4e\n4"
	fmt.Println(unpack.Unpack(input,false))
	// fmt.Println(unpack.Unpack(input,true))
}