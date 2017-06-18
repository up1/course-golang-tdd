package main

import "fmt"

func main() {
	slice1 := []int {1, 2, 3}
	slice2 := append(slice1, 4, 5)

	fmt.Printf("%v\n", slice1)
	fmt.Printf("%v\n", slice2)
}
