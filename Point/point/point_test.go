package point_test

import (
	"Point/point"
	"testing"
)

func TestPointCreating(t *testing.T) {
	p := point.Point{3.54824, 4.64828}
	if p.X != 3.54824 && p.Y != 4.64828 {
		t.Error("Error creating struct with params X: 3.54824, Y: 4.64828")
	}
}
func TestCalcDistance(t *testing.T) {
	p1 := point.Point{10, 13}
	p2 := point.Point{6, 13}
	res := p1.CalcDistance(p2)
	if res != 4 {
		t.Error("Expected: 4, Got: ", res)
	}
	p1_1 := point.Point{0, 0}
	p2_1 := point.Point{0, 0}
	res_1 := p1_1.CalcDistance(p2_1)
	if res_1 != 0 {
		t.Error("Expected: 0, Got: ", res_1)
	}
}

func TestIsInRadius(t *testing.T) {
	p1 := point.Point{4, 4}
	p2 := point.Point{4, 9}
	var n float64 = 6
	res := p1.IsInRadius(p2, n)
	if !res {
		t.Error("Expected: true, Got: ", res)
	}
}
