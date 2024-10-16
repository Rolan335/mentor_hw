package point

import (
	"math"
)

type Point struct {
	X, Y float32
}

func (p Point) CalcDistance(p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p.X)-float64(p2.X), 2) + math.Pow(float64(p.Y)-float64(p2.Y), 2))
}

func (p Point) IsInRadius(p2 Point, n float64) bool {
	return p.CalcDistance(p2) <= n
}

type Polygon struct {
	Point []Point
}

//---
type Polygon2 []Point

func (p Polygon2) Test(){
	
}

