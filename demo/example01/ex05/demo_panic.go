package main

import "fmt"

func f() {
	defer func(){
		str := recover()
		fmt.Println(str)
	}()
	x := []int{1, 2, 3}
	fmt.Println(x[10])
}

func main() {
	f()
	fmt.Println("TODO NEXT")
}