package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func (p *point) DistanceToAnother(anotherPoint *point) float64 {
	xDelta, yDelta := anotherPoint.x-p.x, anotherPoint.y-p.y
	return math.Sqrt(xDelta*xDelta + yDelta*yDelta)
}

func NewPoint(x, y float64) *point {
	return &point{x: x, y: y}
}

func main() {
	p1 := NewPoint(-1, 1)
	p2 := NewPoint(1, 1)

	fmt.Println("Distance:", p1.DistanceToAnother(p2))
}
