package main

import "fmt"

func main() {
	x := make([]int, 5)
	fmt.Println(len(x), cap(x))
	fmt.Printf("%T\n", x)

	y := make([]int, 5, 10)
	fmt.Println(len(y), cap(y))
	fmt.Printf("%T\n", y)
}
