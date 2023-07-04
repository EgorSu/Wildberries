package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p1 *Point) distance(p2 *Point) float64 { // метод для нахождения расстояния
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func newPoint(x, y float64) *Point { //конструктор структуры Point
	return &Point{x: x, y: y}
}

func main() {
	point1 := newPoint(1, 2)
	point2 := newPoint(5, 2)
	fmt.Println(point1.distance(point2))
}
