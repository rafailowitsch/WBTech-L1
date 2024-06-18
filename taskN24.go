package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func Distance(p1, p2 *Point) float64 {
	result := math.Sqrt(math.Pow(math.Abs(p2.x-p1.x), 2) + math.Pow(math.Abs(p2.y-p1.y), 2))

	return result
}

func taskN24() {
	p1 := NewPoint(4, 5)
	p2 := NewPoint(0, 0)

	dis := Distance(p1, p2)

	fmt.Printf("Distance between p1 and p2 is %.2f", dis)
}
