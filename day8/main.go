package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := getData()
	grid, width := parseData(data)

	fmt.Printf("Part 1: %d\n", part1(grid, width))
	fmt.Printf("Part 2: %d\n", part2(grid, width))
}
func part2(grid []int, width int) int {
	height := len(grid) / width

	highestScore := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			score := getScenicScore(grid, width, x, y)
			if score > highestScore {
				highestScore = score
			}
		}
	}

	return highestScore
}
func getScenicScore(grid []int, width, x, y int) int {
	height := len(grid) / width
	cell := x + y*width
	startHeight := grid[cell]

	upScore := getTreesVisible(grid, startHeight, cell-width, -width, y)
	leftScore := getTreesVisible(grid, startHeight, cell-1, -1, x)
	downScore := getTreesVisible(grid, startHeight, cell+width, width, height-y-1)
	rightScore := getTreesVisible(grid, startHeight, cell+1, 1, width-x-1)

	return upScore * rightScore * downScore * leftScore
}
func getTreesVisible(grid []int, startHeight, start, step, count int) int {
	result := 0

	for i := 0; i < count; i++ {
		cell := start + i*step
		value := grid[cell]

		result++

		if value >= startHeight {
			return result
		}
	}

	return result
}
func part1(grid []int, width int) int {
	visibilityMap := map[int]int{}
	height := len(grid) / width

	for y := 0; y < height; y++ {
		buildTreeVisibiltyMap(grid, y*width, 1, width, visibilityMap)
		buildTreeVisibiltyMap(grid, ((y+1)*width)-1, -1, width, visibilityMap)
	}
	for x := 0; x < width; x++ {
		buildTreeVisibiltyMap(grid, x, width, height, visibilityMap)
		buildTreeVisibiltyMap(grid, (height-1)*width+x, -width, height, visibilityMap)
	}

	return len(visibilityMap)
}
func buildTreeVisibiltyMap(grid []int, start, step, count int, target map[int]int) {
	highestTree := -1

	for i := 0; i < count; i++ {
		cell := start + i*step
		value := grid[cell]

		if value > highestTree {
			target[cell] = 1
			highestTree = value
		}
	}
}
func parseData(lines []string) ([]int, int) {
	result := []int{}
	width := len(lines[0])

	for _, line := range lines {
		result = append(result, parseLine(line)...)
	}

	return result, width
}
func parseLine(line string) []int {
	result := []int{}

	numbers := strings.Split(line, "")
	for _, d := range numbers {
		value, _ := strconv.Atoi(d)

		result = append(result, value)
	}

	return result
}
func getData() []string {
	result := []string{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
func getTestData() []string {
	return []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}
}
