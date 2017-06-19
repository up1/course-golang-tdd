package main

import "testing"

func TestInsertAShouldRemoveA(t *testing.T){
	NewBuffer()
	e := "A"
	Insert("A")
	a := Remove()
	if a != e {
		t.Fatalf("Expect %s but got %s", e, a)
	}
}

func TestInsertBShouldRemoveB(t *testing.T){
	NewBuffer()
	e := "B"
	Insert("B")
	a := Remove()
	if a != e {
		t.Fatalf("Expect %s but got %s", e, a)
	}
}

func TestInsertABShouldRemoveAB(t *testing.T){
	NewBuffer()
	Insert("A")
	Insert("B")
	a1 := Remove()
	a2 := Remove()
	if a1 != "A" || a2 != "B" {
		t.Fatalf("Error")
	}
}

func TestInsertABAndRemoveABShouldEmpty(t *testing.T){
	NewBuffer()
	Insert("A")
	Insert("B")
	Remove()
	Remove()
	a := IsEmpty()
	if a != true {
		t.Fatalf("Error")
	}
}