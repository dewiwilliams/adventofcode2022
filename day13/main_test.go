package main

import (
	"testing"
)

func TestFindValues(t *testing.T) {
	lines := getData("./test_input.txt")

	expectedResult := []int{1, 1, 0, 1, 0, 1, 0, 0}

	for i := range expectedResult {
		if inRightOrder(lines[i*2+0], lines[i*2+1]) != expectedResult[i] {
			t.Errorf("Ordering failed, %d: %s vs %s", i, lines[i*2+0], lines[i*2+1])
		}
	}
}
