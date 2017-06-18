package main

import "fmt"

func main() {
	i := 1
	f := 1.0
	s := "Hello"
	b := true
	fmt.Printf("Type %T has value %d\n", i, i)
	fmt.Printf("Type %T has value %.2f\n", f, f)
	fmt.Printf("Type %T has value %s\n", s, s)
	fmt.Printf("Type %T has value %v\n", b, b)
}
