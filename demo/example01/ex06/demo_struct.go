package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) someFunc() float64 {
	return math.Sqrt(p.x)
}

func main() {
	p := Point{x: 1.0, y: 2.0}
	fmt.Println(p.someFunc())
}
