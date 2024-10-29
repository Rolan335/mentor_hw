package main

import (
	"point/parsepoint"
	"flag"
	"fmt"
)

func main() {
	action := flag.Bool("action", true, "true - calcDistance false - IsInRadius")
	p1 := flag.String("point1", "0,0", "type (x,y) for first point")
	p2 := flag.String("point2", "0,0", "type (x,y) for second point")
	n := flag.Float64("n", 0.0, "n for IsInRadius")
	flag.Parse()

	p1Struct, err := parsepoint.Parse(p1)
	if err != nil {
		fmt.Println(err)
	}
	p2Struct, err := parsepoint.Parse(p2)
	if err != nil {
		fmt.Println(err)
	}
	if *action {
		fmt.Println(p1Struct.CalcDistance(p2Struct))
	} else {
		fmt.Println(p1Struct.IsInRadius(p2Struct, *n))
	}

}
