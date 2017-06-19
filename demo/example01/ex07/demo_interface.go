package main

import (
	"math"
	"fmt"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	w, h float64
}

func (r *Rectangle) area() float64 {
	return r.w * r.h
}

func main() {
	c := Circle{1, 1, 10}
	r := Rectangle{5, 10}
	shapes := []Shape{&c, &r}
	for _, s := range shapes {
		fmt.Println(s.area())
	}
}
