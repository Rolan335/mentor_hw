package point

import (
	"fmt"
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

type Polygon []Point

func (p *Polygon) AddPoint(slice ...Point) {
	for _, v := range slice {
		*p = append(*p, v)
	}
}

func (p Polygon) GetPerimeter() float64 {
	perimeter := p[0].CalcDistance(p[len(p) - 1])
	for i := 0; i < len(p) - 1; i++{
		fmt.Println(perimeter)
		perimeter += p[i].CalcDistance(p[i+1])
	}
	return perimeter
}
