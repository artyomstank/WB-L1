package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p1 *Point) distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2.0) + math.Pow(p2.y-p1.y, 2.0))
}

func (p1 *Point) print() string {
	return fmt.Sprintf("{x = %f; y = %f}", p1.x, p1.y)
}

func createPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}
func main() {
	p1 := createPoint(1, 5)
	p2 := createPoint(3, 10)
	fmt.Printf("p1 = %s\n", p1.print())
	fmt.Printf("p2 = %s\n", p2.print())

	fmt.Printf("Расстояние между точками = %f", p1.distance(p2))
}
