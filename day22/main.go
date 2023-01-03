package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const invalid = 0
const empty = 1
const wall = 2

const right = 0
const down = 1
const left = 2
const up = 3

func main() {
	data, width, instructions := getData("./input.txt")

	//printBoard(data, width)
	//fmt.Printf("Instructions: %v\n", instructions)

	fmt.Printf("Part 1: %d\n", part1(data, width, instructions))
	fmt.Printf("Part 2: %d\n", part2(data, width, instructions))
}
func part2(data []int, width int, instructions []int) int {
	cell := findTopLeft(data, width)
	facing := right

	fmt.Printf("Top left: %d, %d\n", cell%width, cell/width)

	for i, instruction := range instructions {
		if i%2 == 0 {
			cell, facing = moveForwardCube(data, width, cell, facing, instruction)
		} else {
			facing = getNextDirection(facing, instruction)
		}

		/*fmt.Printf("Moved to (%d, %d), %d\n", cell%width, cell/width, facing)

		if i > 10 {
			break
		}*/
	}

	// 93210  too low
	// 133229 too low
	// 884470 too high

	return getPassword(cell, width, facing)
}
func getPassword(cell, width, facing int) int {
	x := cell % width
	y := cell / width

	return (y+1)*1000 + (x+1)*4 + facing
}
func getNextDirection(current, turn int) int {
	if turn == left {
		current += 3
	} else if turn == right {
		current++
	}
	current %= 4
	return current
}
func moveForwardCube(data []int, width, cell, facing, count int) (int, int) {
	for i := 0; i < count; i++ {
		newCell, newFacing := getNextEmptyCellCube(data, width, cell, facing)
		if newCell == cell {
			fmt.Printf("Hit wall moveforward\n")
			return cell, facing
		}

		cell = newCell
		facing = newFacing
	}

	return cell, facing
}

func moveForward(data []int, width, cell, facing, count int) int {
	for i := 0; i < count; i++ {
		newCell := getNextEmptyCell(data, width, cell, facing)
		if newCell == cell {
			return newCell
		}
		cell = newCell
	}

	return cell
}
func getNextEmptyCellCube(data []int, width, cell, facing int) (int, int) {
	nextCell, nextFacing := getNextCellCube(data, width, cell, facing)
	if data[nextCell] == wall {
		fmt.Printf("Hit wall\n")
		return cell, facing
	}
	return nextCell, nextFacing
}
func getNextCellCube(data []int, width, cell, facing int) (int, int) {
	height := len(data) / width
	x := cell % width
	y := cell / width

	if facing == right {
		x++
	} else if facing == left {
		x--
	} else if facing == down {
		y++
	} else if facing == up {
		y--
	}

	//inBounds := x >= 0 && x < width && y >= 0 && y < height

	//fmt.Printf("Got: %d, %d, %v\n", x, y, inBounds)

	/*if inBounds {
		nextCell := x + y*width
		if data[nextCell] == wall {
			return cell, facing
		} else if data[nextCell] == empty {
			return nextCell, facing
		}
	}*/

	/*
		    |---|---|
			| 5 | 6 |
			|---|---|
			| 4 |
		|---|---|
		| 2 | 3 |
		|---|---|
		| 1 |
		|---|
	*/

	sideLength := 50

	if x == -1 && y >= 2*sideLength && y < 3*sideLength && facing == left {
		// left from face 2 in to left side face 5
		relativeY := y - 2*sideLength
		x = 1 * sideLength
		y = 50 - relativeY - 1
		facing = right
	} else if x == sideLength-1 && y >= 0 && y < sideLength && facing == left {
		// left from face 5 in to left side face 2
		x = 0
		y = 2*sideLength + (sideLength - y) - 1
		facing = right
	} else if x == -1 && y >= 3*sideLength && y < 4*sideLength && facing == left {
		// left from face 1 in to top of face 5
		relativeY := y - 3*sideLength
		y = 0
		x = 1*sideLength + relativeY
		facing = down
	} else if y == -1 && x >= sideLength && x < 2*sideLength && facing == up {
		// up from face 5 in to left of face 1
		y = 3*sideLength + (x - sideLength)
		x = 0
		facing = right
	} else if x >= 0 && x < sideLength && y == 4*sideLength && facing == down {
		// down from face 1 in to top of face 6
		x += 2 * sideLength
		y = 0
	} else if y == -1 && x >= 2*sideLength && x < 3*sideLength && facing == up {
		// up from face 6 in to bottom of face 1
		x -= 2 * sideLength
		y = 4*sideLength - 1
	} else if x == 2*sideLength && y >= 2*sideLength && y < 3*sideLength && facing == right {
		// right from face 3 in to right side of face 6
		y -= 2 * sideLength
		y = sideLength - 1 - y
		x = 3*sideLength - 1
		facing = left
	} else if x == 3*sideLength && y >= 0 && y < sideLength && facing == right {
		// right from face 6 in to right side of face 3
		y = sideLength - 1 - y
		y += 2 * sideLength
		x = 2*sideLength - 1
		facing = left
	} else if x >= 0 && x < sideLength && y == 2*sideLength-1 && facing == up {
		// up from face 2 in to left side of face 4
		y = sideLength + x
		x = sideLength
		facing = right
	} else if x == sideLength-1 && y >= sideLength && y < 2*sideLength && facing == left {
		// left from face 4 in to the top of face 2
		x = y - sideLength
		y = 2 * sideLength
		facing = down
	} else if x == sideLength && y >= 3*sideLength && y < 4*sideLength && facing == right {
		// right from face 1 in to the bottom of face 3
		x = 50 + (y - 3*sideLength)
		y = 3*sideLength - 1
		facing = up
	} else if x >= sideLength && x < 2*sideLength && y == 3*sideLength && facing == down {
		// down from face 3 in to right side of face 1
		y = 3*sideLength + (x - sideLength)
		x = sideLength - 1
		facing = left
	} else if x == 2*sideLength && y >= sideLength && y < 2*sideLength && facing == right {
		// right from face 4 up in to face 6
		x = 2*sideLength + (y - sideLength)
		y = sideLength - 1
		facing = up
	} else if x >= 2*sideLength && x < 3*sideLength && y == sideLength && facing == down {
		// down from face 6 in to the left of side 4
		y = sideLength + (x - 2*sideLength)
		x = 2*sideLength - 1
		facing = left
	}

	inBounds := x >= 0 && x < width && y >= 0 && y < height

	if !inBounds {
		log.Fatalf("Out of bounds: (%d, %d), %d", x, y, facing)
	}
	if data[x+y*width] == invalid {
		log.Fatalf("not on valid cell: (%d, %d)\n", x, y)
	}

	return x + y*width, facing
}

func part1(data []int, width int, instructions []int) int {
	cell := findTopLeft(data, width)
	facing := right

	for i, instruction := range instructions {
		if i%2 == 0 {
			cell = moveForward(data, width, cell, facing, instruction)
		} else {
			facing = getNextDirection(facing, instruction)
		}
	}

	return getPassword(cell, width, facing)
}
func getNextEmptyCell(data []int, width, cell, facing int) int {
	nextCell := getNextCell(data, width, cell, facing)
	if data[nextCell] == wall {
		return cell
	}
	return nextCell
}
func getNextCell(data []int, width, cell, facing int) int {
	height := len(data) / width
	x := cell % width
	y := cell / width

	for {
		if facing == right {
			x++
			if x == width {
				x = 0
			}
		} else if facing == left {
			x--
			if x == -1 {
				x = width - 1
			}
		} else if facing == down {
			y++
			if y == height {
				y = 0
			}
		} else if facing == up {
			y--
			if y == -1 {
				y = height - 1
			}
		}

		newCell := x + y*width
		if data[newCell] != invalid {
			return newCell
		}
	}
}
func printBoard(data []int, width int) {
	height := len(data) / width
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := x + y*width
			if data[cell] == invalid {
				fmt.Print(" ")
			} else if data[cell] == empty {
				fmt.Print(".")
			} else if data[cell] == wall {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
func findTopLeft(data []int, width int) int {
	height := len(data) / width

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			cell := x + y*width
			if data[cell] == empty {
				return cell
			}
		}
	}

	return -1
}
func getOppositeDirection(direction int) int {
	if direction == left {
		return right
	} else if direction == right {
		return left
	} else if direction == up {
		return down
	} else if direction == down {
		return up
	}
	return 0
}
func getData(filename string) ([]int, int, []int) {
	result := []int{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	width := 0
	instructions := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			scanner.Scan()
			instructions = scanner.Text()
			break
		}

		lines = append(lines, line)
		if len(line) > width {
			width = len(line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		for _, r := range line {
			if r == ' ' {
				result = append(result, invalid)
			} else if r == '.' {
				result = append(result, empty)
			} else if r == '#' {
				result = append(result, wall)
			}
		}
		for i := len(line); i < width; i++ {
			result = append(result, invalid)
		}
	}

	instructionsArray := []int{}
	current := ""
	for _, r := range instructions {
		if r == 'L' {
			value, _ := strconv.Atoi(current)
			instructionsArray = append(instructionsArray, value, left)
			current = ""
		} else if r == 'R' {
			value, _ := strconv.Atoi(current)
			instructionsArray = append(instructionsArray, value, right)
			current = ""
		} else {
			current += string(r)
		}
	}
	value, _ := strconv.Atoi(current)
	instructionsArray = append(instructionsArray, value)

	return result, width, instructionsArray
}
