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

	printBoard(data, width)
	//fmt.Printf("Instructions: %v\n", instructions)

	fmt.Printf("Part 1: %d\n", part1(data, width, instructions))
}
func part1(data []int, width int, instructions []int) int {
	cell := findTopLeft(data, width)
	fmt.Printf("Got topleft: %d\n", cell)
	facing := right

	for i, instruction := range instructions {
		if i%2 == 0 {
			cell = moveForward(data, width, cell, facing, instruction)
		} else {
			facing = getNextDirection(facing, instruction)
		}
	}

	x := cell % width
	y := cell / width
	fmt.Printf("End: %d, %d, %d\n", x, y, facing)

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
		fmt.Printf("%d, ", newCell)
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

	fmt.Printf("Got instructions: %s\n", instructions)

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
