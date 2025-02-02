package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Sum(2, 3) = %d; want %d", result, expected)
	}
}

func Sum(a int, b int) interface{} {
	return a + b
}
