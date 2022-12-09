package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const up = 1
const down = 2
const left = 3
const right = 4

type position struct {
	x int
	y int
}

func (p position) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func main() {
	data := getData("./input.txt")

	fmt.Printf("Part1: %d\n", part1(data))
}
func part1(data []int) int {
	headPosition := position{}
	tailPosition := position{}

	coverage := make(map[string]int)

	for i := 0; i < len(data)/2; i++ {
		headPosition, tailPosition = processLine(headPosition, tailPosition, data[i*2+0], data[i*2+1], coverage)
	}

	return len(coverage)
}
func processLine(startingHeadPosition, startingTailPosition position, direction, amount int, coverage map[string]int) (position, position) {
	headPosition := startingHeadPosition
	tailPosition := startingTailPosition

	for i := 0; i < amount; i++ {
		if direction == up {
			headPosition.y--
		} else if direction == down {
			headPosition.y++
		} else if direction == left {
			headPosition.x--
		} else if direction == right {
			headPosition.x++
		}

		headPosition, tailPosition = processTail(headPosition, tailPosition)

		coverage[makeKey(tailPosition)] = 1
	}

	return headPosition, tailPosition
}
func processTail(h, t position) (position, position) {
	xdiff := h.x - t.x
	ydiff := h.y - t.y

	if abs(xdiff) > 1 || abs(ydiff) > 1 {
		t.x += toUnit(xdiff)
		t.y += toUnit(ydiff)
	}

	return h, t
}
func toUnit(v int) int {
	if v > 0 {
		return 1
	} else if v < 0 {
		return -1
	} else {
		return 0
	}
}
func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
func makeKey(p position) string {
	return fmt.Sprintf("%d_%d", p.x, p.y)
}
func getData(filename string) []int {
	result := []int{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		if parts[0] == "U" {
			result = append(result, up)
		} else if parts[0] == "D" {
			result = append(result, down)
		} else if parts[0] == "L" {
			result = append(result, left)
		} else if parts[0] == "R" {
			result = append(result, right)
		}

		length, _ := strconv.Atoi(parts[1])

		result = append(result, length)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
