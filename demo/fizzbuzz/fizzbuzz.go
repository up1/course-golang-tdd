package main

import (
	"strconv"
)

func Get(d int) string {
	if divideBy3(d) {
		return "Fizz"
	}
	if divideBy5(d) {
		return "Buzz"
	}
	return strconv.Itoa(d)
}
func divideBy5(d int) bool {
	return d%5 == 0
}
func divideBy3(d int) bool {
	return d%3 == 0
}
