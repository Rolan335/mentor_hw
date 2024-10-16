package parsepoint

import (
	p "Point/point"
	"errors"
	"strconv"
	"strings"
)

func Parse(input *string) (p.Point, error) {
	//check if nil
	pSlice := strings.Split(*input, ",")
	if len(pSlice) < 2 {
		return p.Point{}, errors.New("bad input. expected - x,y got - " + *input)
	}
	//err should be unique
	x, err := strconv.ParseFloat(pSlice[0], 32)
	if err != nil {
		return p.Point{}, err
	}
	y, err := strconv.ParseFloat(pSlice[1], 32)
	if err != nil {
		return p.Point{}, err
	}
	point := p.Point{X: float32(x), Y: float32(y)}
	return point, nil
}
