package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const empty = 0
const elf = 1

func main() {
	filename := "./input.txt"
	fmt.Printf("Part 1: %d\n", part1(filename))
	fmt.Printf("Part 2: %d\n", part2(filename))
}
func part1(filename string) int {
	data, width := getData(filename)
	newWidth := 100

	expandedBoard := expandGrid(data, width, newWidth)
	for i := 0; i < 10; i++ {
		step(expandedBoard, newWidth, i)
	}

	minX, maxX, minY, maxY := getExtent(expandedBoard, newWidth)
	return getEmptyTiles(expandedBoard, newWidth, minX, maxX, minY, maxY)
}
func part2(filename string) int {
	data, width := getData(filename)
	newWidth := 1000

	expandedBoard := expandGrid(data, width, newWidth)
	for i := 0; ; i++ {
		if !step(expandedBoard, newWidth, i) {
			return i + 1
		}
	}
}
func getEmptyTiles(data []int, width, minX, maxX, minY, maxY int) int {
	result := 0

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			cell := x + y*width
			if data[cell] == empty {
				result++
			}
		}
	}

	return result
}
func expandGrid(data []int, width, sideLength int) []int {
	height := len(data) / width
	xOffset := (sideLength - width) / 2
	yOffset := (sideLength - height) / 2

	result := make([]int, sideLength*sideLength)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			newCell := (xOffset + x) + (yOffset+y)*sideLength
			oldCell := x + y*width

			result[newCell] = data[oldCell]
		}
	}

	return result
}
func getExtent(data []int, width int) (int, int, int, int) {
	height := len(data) / width

	minX := width - 1
	maxX := 0
	minY := height - 1
	maxY := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := x + y*width

			if data[cell] == empty {
				continue
			}

			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	return minX, maxX, minY, maxY
}
func getTarget(data []int, width, cell, startDirection int) int {
	/*
	   If there is no Elf in the N, NE, or NW adjacent positions, the Elf proposes moving north one step.
	   If there is no Elf in the S, SE, or SW adjacent positions, the Elf proposes moving south one step.
	   If there is no Elf in the W, NW, or SW adjacent positions, the Elf proposes moving west one step.
	   If there is no Elf in the E, NE, or SE adjacent positions, the Elf proposes moving east one step.
	*/

	height := len(data) / width
	x := cell % width
	y := cell / width

	if x == 0 || x == width-1 || y == 0 || y == height-1 {
		fmt.Printf("Cell: %d, %d\n", x, y)
		log.Fatal("Out of bounds!")
	}

	n := cell - width
	ne := n + 1
	nw := n - 1
	e := cell + 1
	w := cell - 1
	s := cell + width
	se := s + 1
	sw := s - 1

	if data[n] == empty &&
		data[ne] == empty &&
		data[nw] == empty &&
		data[e] == empty &&
		data[w] == empty &&
		data[s] == empty &&
		data[se] == empty &&
		data[sw] == empty {
		return cell
	}

	for i := 0; i < 4; i++ {
		direction := (startDirection + i) % 4

		if direction == 0 {
			if data[nw] == empty && data[n] == empty && data[ne] == empty {
				return n
			}
		} else if direction == 1 {
			if data[sw] == empty && data[s] == empty && data[se] == empty {
				return s
			}
		} else if direction == 2 {
			if data[nw] == empty && data[w] == empty && data[sw] == empty {
				return w
			}
		} else if direction == 3 {
			if data[ne] == empty && data[e] == empty && data[se] == empty {
				return e
			}
		}
	}

	return cell
}
func step(data []int, width, iteration int) bool {
	height := len(data) / width

	targets := make(map[int]int)
	targetCounts := make(map[int]int)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := x + y*width

			if data[cell] == empty {
				continue
			}

			target := getTarget(data, width, cell, iteration)

			targets[cell] = target
			targetCounts[target] = targetCounts[target] + 1
		}
	}

	moved := false

	for k, v := range targets {
		if targetCounts[v] > 1 {
			continue
		}
		if k == v {
			continue
		}

		data[k] = empty
		data[v] = elf
		moved = true
	}

	return moved
}
func contains(haystack []int, needle int) bool {

	for _, item := range haystack {
		if item == needle {
			return true
		}
	}

	return false
}
func printBoard(data []int, width int) {
	minX, maxX, minY, maxY := getExtent(data, width)

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			cell := x + y*width
			if data[cell] == empty {
				fmt.Print(".")
			} else if data[cell] == elf {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
func getData(filename string) ([]int, int) {
	result := []int{}
	width := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, r := range line {
			if r == '.' {
				result = append(result, empty)
			} else if r == '#' {
				result = append(result, elf)
			}
		}

		if len(line) > width {
			width = len(line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result, width
}
