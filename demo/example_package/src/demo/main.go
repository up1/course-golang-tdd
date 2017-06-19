package main

import (
	"demo/math"
	"fmt"
)

func main() {
	datas := []float64{1, 2, 3, 4, 5}
	average :=  math.Average(datas)
	fmt.Println(average)
	if false {
		fmt.Println("False")
	}
}
