package main

import "fmt"

func errorHappen() {
	str := recover()
	fmt.Println(str)
}

func f() {
	defer errorHappen()
	x := []int{1, 2, 3}
	fmt.Println(x[10])
}

func main() {
	f()
	fmt.Println("TODO NEXT")
}