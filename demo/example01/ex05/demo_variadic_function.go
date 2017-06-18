package main

import "fmt"

func add(ops ...int) int {
	total := 0
	for _, v := range ops {
		total += v
	}
	return total
}

func main() {
	fmt.Println(add())
	fmt.Println(add(1))
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3))
}
