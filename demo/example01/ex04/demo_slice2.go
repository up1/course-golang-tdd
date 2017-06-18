package main

import "fmt"

func main() {
	arr := [5]int {1, 2, 3, 4, 5}

	fmt.Printf("%v\n", arr[0:5])
	fmt.Printf("%v\n", arr[0:len(arr)])
	fmt.Printf("%v\n", arr[0:])
	fmt.Printf("%v\n", arr[:5])
	fmt.Printf("%v\n", arr[:])

	fmt.Printf("%v\n", arr[1:5])
	fmt.Printf("%v\n", arr[1:4])
}