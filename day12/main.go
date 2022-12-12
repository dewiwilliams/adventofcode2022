package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type heightMap struct {
	values []int
	width  int
	height int
	start  int
	finish int

	iterations int
}

func (h *heightMap) getNeighbours(x, y int) []int {
	result := []int{}

	if x > 0 {
		result = append(result, (x-1)+y*h.width)
	}
	if x < h.width-1 {
		result = append(result, (x+1)+y*h.width)
	}
	if y > 0 {
		result = append(result, x+(y-1)*h.width)
	}
	if y < h.height-1 {
		result = append(result, x+(y+1)*h.width)
	}

	return result
}
func (h *heightMap) getAccessibleNeighbours(x, y int) []int {
	result := []int{}
	point := x + y*h.width
	neighbours := h.getNeighbours(x, y)

	for _, n := range neighbours {
		if h.values[n]-h.values[point] <= 1 {
			result = append(result, n)
		}
	}

	return result
}
func (h *heightMap) traverse() int {
	visited := make(map[int]int)
	costMap := make(map[int]int)

	for i := 0; i < h.width*h.height; i++ {
		costMap[i] = 999999999
	}

	result := h.traverseFromPoint(h.start, 0, costMap, visited)
	//h.outputCostMap(costMap)
	return result
}
func (h *heightMap) outputCostMap(costMap map[int]int) {
	for y := 0; y < h.height; y++ {
		for x := 0; x < h.width; x++ {
			point := x + y*h.width
			if costMap[point] > 1000 {
				fmt.Printf("  X,")
			} else if costMap[point] < 10 {
				fmt.Printf("  %d,", costMap[point])
			} else if costMap[point] < 100 {
				fmt.Printf(" %d,", costMap[point])
			} else {
				fmt.Printf("%d,", costMap[point])
			}
		}
		fmt.Print("\n")
	}

}
func (h *heightMap) traverseFromPoint(p int, count int, costMap, visited map[int]int) int {
	result := 999999999
	x := p % h.width
	y := p / h.width
	visited[p] = 1

	h.iterations++

	accessibleNeighbours := h.getAccessibleNeighbours(x, y)

	for i, n := range accessibleNeighbours {
		if costMap[n] <= count+1 {
			accessibleNeighbours[i] = h.width * h.height
			continue
		}
		costMap[n] = count + 1
	}
	for _, n := range accessibleNeighbours {
		if n == h.width*h.height {
			continue
		}
		if visited[n] != 0 {
			continue
		}

		if n == h.finish {
			visited[p] = 0
			return 1
		}

		pointCount := h.traverseFromPoint(n, count+1, costMap, visited) + 1
		if pointCount < result {
			result = pointCount
		}
	}

	visited[p] = 0

	return result
}

func main() {
	values, width, startFinish := getData("./input.txt")

	h := heightMap{
		values: values,
		width:  width,
		height: len(values) / width,
		start:  startFinish[0],
		finish: startFinish[1],
	}

	fmt.Printf("Part 1: %d\n", h.traverse())
	fmt.Printf("Iterations: %d\n", h.iterations)
}
func getRuneScore(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r) - int('a') + 1
	}
	if r == 'S' {
		return getRuneScore('a')
	} else if r == 'E' {
		return getRuneScore('z')
	}

	return 0
}
func getData(filename string) ([]int, int, [2]int) {
	heightMap := []int{}
	width := 0
	startFinish := [2]int{-1, -1}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		width = len(line)

		//for i := 0; i < len(line); i++ {
		for i, r := range line {
			if r == 'S' {
				startFinish[0] = i + lineNumber*width
			} else if r == 'E' {
				startFinish[1] = i + lineNumber*width
			}

			heightMap = append(heightMap, getRuneScore(r))
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return heightMap, width, startFinish
}
