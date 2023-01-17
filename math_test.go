package main

import "testing"

func testSum(t *testing.T) {
	input1 := 20
	input2 := 30
	totalExpected := 50

	total := Sum(input1, input2)

	if total != totalExpected {
		t.Errorf("Sum result invalid. Was: %d, expected: %d", total, totalExpected)
	}
}
