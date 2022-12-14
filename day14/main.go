package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const empty = 0
const wall = 1
const sand = 2

func main() {
	lines := getData("./input.txt")

	fmt.Printf("Part 1: %d\n", part1(lines))
}

func part1(dataLines [][]int) int {

	data, width, _, _ := buildMap(dataLines)

	for i := 0; i < 99999; i++ {
		if dropSand(data, width) {
			return i
		}
	}

	return 0
}

func dropSand(data []int, width int) bool {

	height := len(data) / width

	x := 500
	y := 0

	for {
		down := x + (y+1)*width
		downLeft := (x - 1) + (y+1)*width
		downRight := (x + 1) + (y+1)*width

		if y == height-1 {
			//Reached bottom
			return true
		}

		if data[down] == empty {
			y++
		} else if data[downLeft] == empty {
			y++
			x--
		} else if data[downRight] == empty {
			y++
			x++
		} else {
			data[x+y*width] = sand
			return false
		}
	}
}

func buildMap(data [][]int) ([]int, int, int, int) {
	yMax := getLowestPoint(data)
	xMin, xMax := getXRange(data)
	width := xMax + 2

	result := make([]int, (yMax+2)*width)

	for _, wall := range data {
		addWall(result, wall, width)
	}

	return result, width, xMin, xMax
}
func traceWall(target []int, width, startX, endX int) {
	height := len(target) / width

	for y := 0; y < height; y++ {
		for x := startX; x <= endX; x++ {
			value := target[x+y*width]
			if value == empty {
				fmt.Print(".")
			} else if value == wall {
				fmt.Print("#")
			} else if value == sand {
				fmt.Print("o")
			}
		}
		fmt.Println("")
	}
}
func addWall(target, wall []int, width int) {
	for i := 1; i < len(wall)/2; i++ {
		addWallSection(target, width, wall[i*2-2], wall[i*2-1], wall[i*2+0], wall[i*2+1])
	}
}
func addWallSection(target []int, width, x1, y1, x2, y2 int) {
	if x1 == x2 {
		startY := y1
		endY := y2

		if y1 > y2 {
			startY = y2
			endY = y1
		}

		for y := startY; y <= endY; y++ {
			target[x1+y*width] = wall
		}
	} else if y1 == y2 {
		startX := x1
		endX := x2

		if x1 > x2 {
			startX = x2
			endX = x1
		}

		for x := startX; x <= endX; x++ {
			target[x+y1*width] = wall
		}
	}
}
func getXRange(data [][]int) (int, int) {
	min := 500
	max := 500

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i])/2; j++ {
			value := data[i][j*2+0]

			if value < min {
				min = value
			}
			if value > max {
				max = value
			}
		}
	}

	return min, max
}
func getLowestPoint(data [][]int) int {

	result := 0

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i])/2; j++ {
			value := data[i][j*2+1]

			if value > result {
				result = value
			}
		}
	}

	return result
}
func getData(filename string) [][]int {
	result := [][]int{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		result = append(result, parseLine(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
func parseLine(line string) []int {
	result := []int{}

	pairs := strings.Split(line, " ")
	for i, pair := range pairs {
		if i%2 == 1 {
			continue
		}

		parts := strings.Split(pair, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		result = append(result, x, y)
	}

	return result
}
