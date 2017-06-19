package main

import "testing"

func TestDefault(t *testing.T){
	checkResult(Get(1), "1", t)
	checkResult(Get(2), "2", t)
	checkResult(Get(4), "4", t)
	checkResult(Get(7), "7", t)
	checkResult(Get(8), "8", t)
}

func TestFizz(t *testing.T){
	checkResult(Get(3), "Fizz", t)
	checkResult(Get(6), "Fizz", t)
	checkResult(Get(9), "Fizz", t)
}

func TestBuzz(t *testing.T){
	checkResult(Get(5), "Buzz", t)
	checkResult(Get(10), "Buzz", t)
}

func TestFizzBuzz(t *testing.T){
	checkResult(Get(15), "FizzBuzz", t)
}

func checkResult(a string, e string, t *testing.T){
	if a != e {
		t.Fatalf("Expected %s but got %s", e, a)
	}
}

