package main

import (
	"testing"
	"reflect"
)

func TestCase1(t *testing.T) {
	expected := []int{1}
	result := sort(1)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}

func TestCase2(t *testing.T) {
	expected := []int{1, 2}
	result := sort(1, 2)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}

func TestCase3(t *testing.T) {
	expected := []int{1, 2}
	result := sort(2, 1)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}

func TestCase4(t *testing.T) {
	expected := []int{1, 2, 3}
	result := sort(1, 2, 3)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}

func TestCase5(t *testing.T) {
	expected := []int{1, 2, 3}
	result := sort(2, 1, 3)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}

func TestCase6(t *testing.T) {
	expected := []int{1, 2, 3}
	result := sort(1, 3, 2)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}

func TestCase7(t *testing.T) {
	expected := []int{1, 2, 3}
	result := sort(3, 2, 1)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}







func checkResult(expected, result []int) bool {
	if result == nil {
		return false
	}

	if len(result) != len(expected) {
		return false
	}

	for i, _ := range expected {
		if expected[i] != result[i] {
			return false
		}
	}

	return true

	//return reflect.DeepEqual(expected, result)
}

