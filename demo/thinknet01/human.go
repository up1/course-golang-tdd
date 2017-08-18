package main

import "fmt"

type Human struct {
	firstname string
	lastname string
	age int
}

func (h Human) eat() {
	fmt.Println("Human can eat!! ", h)
}

func NewFullHuman(f string, l string, a int) Human {
	return Human{f, l, a}
}

func main() {
	h1 := Human{firstname:"somkiat"}
	h1.eat()
}
