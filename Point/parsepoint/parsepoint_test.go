package parsepoint

import (
	p "point/point"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	input_1 := "1.2,3.4"
	res, err := Parse(&input_1)
	if res != (p.Point{X: 1.2, Y: 3.4}) && err != nil {
		t.Error("Expected: 1.2, 3.4 . Got: ", fmt.Sprintf("%v", res))
	}
}

func TestErrorHandlers(t *testing.T) {
	input_1 := "1.4"
	res, err := Parse(&input_1)
	if err == nil {
		t.Error("expected err. Got: ", res, err)
	}
	input_2 := "1.4,.."
	res, err = Parse(&input_2)
	if err == nil {
		t.Error("expected err. Got: ", res, err)
	}
}
