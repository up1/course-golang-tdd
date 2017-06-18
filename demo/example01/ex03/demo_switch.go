package main

import "fmt"

func main() {
	i := 1

	switch i {
	case 0:
		fmt.Println("case 0")
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	case 3:
		fmt.Println("case 2")
	default:
		fmt.Println("Unknow number")
	}
}
