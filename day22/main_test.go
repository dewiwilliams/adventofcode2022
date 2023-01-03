package main

import (
	"fmt"
	"testing"
)

func TestCubeMovement(t *testing.T) {
	type state struct {
		x      int
		y      int
		facing int
	}
	type testCase struct {
		start state
		end   state
	}

	data, width, _ := getData("./input.txt")

	testCases := []testCase{
		{start: state{x: 0, y: 100, facing: left}, end: state{x: 50, y: 49, facing: right}},   // 0
		{start: state{x: 0, y: 149, facing: left}, end: state{x: 50, y: 0, facing: right}},    // 1
		{start: state{x: 0, y: 150, facing: left}, end: state{x: 50, y: 0, facing: down}},     // 2
		{start: state{x: 0, y: 199, facing: left}, end: state{x: 99, y: 0, facing: down}},     // 3
		{start: state{x: 0, y: 199, facing: down}, end: state{x: 100, y: 0, facing: down}},    // 4
		{start: state{x: 49, y: 199, facing: down}, end: state{x: 149, y: 0, facing: down}},   // 5
		{start: state{x: 99, y: 100, facing: right}, end: state{x: 149, y: 49, facing: left}}, // 6
		{start: state{x: 99, y: 149, facing: right}, end: state{x: 149, y: 0, facing: left}},  // 7

		{start: state{x: 0, y: 100, facing: up}, end: state{x: 50, y: 50, facing: right}},   // 8
		{start: state{x: 49, y: 100, facing: up}, end: state{x: 50, y: 99, facing: right}},  // 9
		{start: state{x: 49, y: 150, facing: right}, end: state{x: 50, y: 149, facing: up}}, // 10
		{start: state{x: 49, y: 199, facing: right}, end: state{x: 99, y: 149, facing: up}}, // 11
		{start: state{x: 99, y: 50, facing: right}, end: state{x: 100, y: 49, facing: up}},  // 12
		{start: state{x: 99, y: 99, facing: right}, end: state{x: 149, y: 49, facing: up}},  // 13

		{start: state{x: 0, y: 150, facing: up}, end: state{x: 0, y: 149, facing: up}},         // 14
		{start: state{x: 49, y: 150, facing: up}, end: state{x: 49, y: 149, facing: up}},       // 15
		{start: state{x: 49, y: 100, facing: right}, end: state{x: 50, y: 100, facing: right}}, // 16
		{start: state{x: 49, y: 149, facing: right}, end: state{x: 50, y: 149, facing: right}}, // 17
		{start: state{x: 50, y: 49, facing: down}, end: state{x: 50, y: 50, facing: down}},     // 18
		{start: state{x: 99, y: 49, facing: down}, end: state{x: 99, y: 50, facing: down}},     // 19
		{start: state{x: 50, y: 99, facing: down}, end: state{x: 50, y: 100, facing: down}},    // 20
		{start: state{x: 99, y: 99, facing: down}, end: state{x: 99, y: 100, facing: down}},    // 21
		{start: state{x: 99, y: 0, facing: right}, end: state{x: 100, y: 0, facing: right}},    // 22
		{start: state{x: 99, y: 49, facing: right}, end: state{x: 100, y: 49, facing: right}},  // 23
	}

	//Forward case
	for i, testCase := range testCases {
		cell := testCase.start.x + testCase.start.y*width
		nextCell, nextDirection := getNextEmptyCellCube(data, width, cell, testCase.start.facing)
		expectedCell := testCase.end.x + testCase.end.y*width

		if nextCell != expectedCell || nextDirection != testCase.end.facing {
			t.Errorf("Cell/direction don't match for forward case %d: %s vs %s, %d vs %d", i,
				formatCell(nextCell, width), formatCell(expectedCell, width), nextDirection, testCase.end.facing)
		}
	}

	//Backwards case
	for i, testCase := range testCases {
		cell := testCase.end.x + testCase.end.y*width
		nextCell, nextDirection := getNextEmptyCellCube(data, width, cell, getOppositeDirection(testCase.end.facing))
		expectedCell := testCase.start.x + testCase.start.y*width

		if nextCell != expectedCell || nextDirection != getOppositeDirection(testCase.start.facing) {
			t.Errorf("Cell/direction don't match for reverse case %d: %s vs %s, %d vs %d", i,
				formatCell(nextCell, width), formatCell(expectedCell, width), nextDirection, getOppositeDirection(testCase.start.facing))
		}
	}
}
func formatCell(cell, width int) string {
	x := cell % width
	y := cell / width
	return fmt.Sprintf("(%d, %d)", x, y)
}
