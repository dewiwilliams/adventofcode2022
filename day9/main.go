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
	fmt.Printf("Part2: %d\n", part2(data))
}
func part2(data []int) int {
	coverage := make(map[string]int)

	p := []position{}
	for i := 0; i < 10; i++ {
		p = append(p, position{})
	}

	for i := 0; i < len(data)/2; i++ {
		p = processLine(p, data[i*2+0], data[i*2+1], coverage)
	}

	return len(coverage)
}
func part1(data []int) int {
	coverage := make(map[string]int)

	p := []position{}
	p = append(p, position{}, position{})

	for i := 0; i < len(data)/2; i++ {
		p = processLine(p, data[i*2+0], data[i*2+1], coverage)
	}

	return len(coverage)
}
func processLine(nodes []position, direction, amount int, coverage map[string]int) []position {
	for i := 0; i < amount; i++ {
		nodes[0] = processMovement(nodes[0], direction)

		for j := 1; j < len(nodes); j++ {
			nodes[j] = processTail(nodes[j-1], nodes[j])
		}

		coverage[makeKey(nodes[len(nodes)-1])] = 1
	}

	return nodes
}
func processMovement(p position, direction int) position {
	if direction == up {
		p.y--
	} else if direction == down {
		p.y++
	} else if direction == left {
		p.x--
	} else if direction == right {
		p.x++
	}

	return p
}
func processTail(h, t position) position {
	xdiff := h.x - t.x
	ydiff := h.y - t.y

	if abs(xdiff) > 1 || abs(ydiff) > 1 {
		t.x += toUnit(xdiff)
		t.y += toUnit(ydiff)
	}

	return t
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
