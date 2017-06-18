package main

import "fmt"

func call() (int, int) {
	return 1, 2
}

func main() {
	x, y := call()
	fmt.Print(x, y)
}
