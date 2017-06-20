package main

import (
  "testing"
)

func TestReturnHelloWorldWhenSth(t *testing.T) {
  e := "Hello World"
  a := sayHi() //Action
  //Assert or Checking
  if(a != e) {
    t.Fatalf("Expected %s but got %s", e, a)
  }
}
