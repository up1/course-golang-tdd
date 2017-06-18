package main

import "fmt"

func first() {
	fmt.Println("Call first")
}

func second() {
	fmt.Println("Call second")
}

func main() {
	defer second()
	first()
}