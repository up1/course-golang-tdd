package main

import "fmt"

func average(data []int) (result float64) {
	total := 0.0
	for _, v := range data {
		total += float64(v)
	}
	result = total/float64(len(data))
	return
}

func main() {
	data := []int{10, 2, 3, 4, 5}
	fmt.Println(average(data))
}
