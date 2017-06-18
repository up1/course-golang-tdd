package ex01

import(
	"testing"
)

func TestHello(t *testing.T) {
	expectedResult := "Hello my first testing"
	result := hello()
	if result != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, result)
	}
}
