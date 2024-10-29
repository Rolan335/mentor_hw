package parsepoint

import (
	p "point/point"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Parse(input *string) (p.Point, error) {
	if input == nil {
		return p.Point{}, errors.New("bad input. nil pointer in input")
	}
	pSlice := strings.Split(*input, ",")
	if len(pSlice) < 2 {
		return p.Point{}, errors.New("bad input. expected - x,y got - " + *input)
	}
	x, err := strconv.ParseFloat(pSlice[0], 32)
	if err != nil {
		return p.Point{}, fmt.Errorf("error parsing x - %v", err)
	}
	y, err := strconv.ParseFloat(pSlice[1], 32)
	if err != nil {
		return p.Point{}, fmt.Errorf("error parsing y - %v", err)
	}
	point := p.Point{X: float32(x), Y: float32(y)}
	return point, nil
}
