package ex02

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var dataTests = []struct{
		op1 int
		op2 int
		expectedResult int
	}{
		{1, 2, 3},
		{5, 10, 15},
		{10, -5, 5},
	}

	for _, test := range dataTests{
		result := add(test.op1, test.op2)
		if result != test.expectedResult {
			t.Fatalf("Expected %d but got %d", test.expectedResult, result)
		}
	}
}
