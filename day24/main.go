package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const north = 10
const east = 11
const south = 12
const west = 13

const empty = 100
const wall = 101

type blizzard struct {
	x         int
	y         int
	direction int
}

func main() {
	filename := "./input.txt"
	fmt.Printf("Part 1: %d\n", part1(filename))
}
func part1(filename string) int {

	grid, width, blizzards := getData(filename)
	height := len(grid) / width
	fmt.Printf("Blizzards area: %d X %d\n", width-2, height-2)

	bestResult := 1000
	statesSeen := []int{}
	search(grid, width, [][]blizzard{blizzards}, 0, 1, &statesSeen, &bestResult)

	return bestResult
}
func blizzardsEqual(blizzards1, blizzards2 []blizzard) bool {
	for i := range blizzards1 {
		if blizzards1[i].x != blizzards2[i].x || blizzards1[i].y != blizzards2[i].y || blizzards1[i].direction != blizzards2[i].direction {
			return false
		}
	}
	return true
}
func search(grid []int, width int, blizzards [][]blizzard, stepNumber, position int, statesSeen *[]int, bestResult *int) {
	if stepNumber > *bestResult {
		return
	}

	height := len(grid) / width
	stateIndex := position + (stepNumber * 10000)
	if contains(*statesSeen, stateIndex) {
		return
	}
	*statesSeen = append(*statesSeen, stateIndex)

	/*if len(*statesSeen)%1000 == 0 {
		fmt.Printf("Seen %d states\n", len(*statesSeen))
	}*/

	if stepNumber >= len(blizzards) {
		blizzards = append(blizzards, stepBlizzards(width, height, blizzards[len(blizzards)-1]))
	}

	validMoves := getValidMoves(grid, width, blizzards[stepNumber], position)
	sort.Sort(sort.Reverse(sort.IntSlice(validMoves)))

	for _, nextPosition := range validMoves {
		if nextPosition == width*height-2 {
			if stepNumber < *bestResult {
				fmt.Printf("Lowering best result: %d -> %d\n", *bestResult, stepNumber)
				*bestResult = stepNumber
			}
			return
		}

		search(grid, width, blizzards, stepNumber+1, nextPosition, statesSeen, bestResult)
	}
}
func getValidMoves(grid []int, width int, blizzards []blizzard, position int) []int {
	result := []int{}

	x := position % width
	y := position / width

	if grid[position] != wall && !isInBlizzard(x, y, blizzards) {
		result = append(result, position)
	}
	if grid[position-1] != wall && !isInBlizzard(x-1, y, blizzards) {
		result = append(result, position-1)
	}
	if grid[position+1] != wall && !isInBlizzard(x+1, y, blizzards) {
		result = append(result, position+1)
	}
	if y > 0 && grid[position-width] != wall && !isInBlizzard(x, y-1, blizzards) {
		result = append(result, position-width)
	}
	if grid[position+width] != wall && !isInBlizzard(x, y+1, blizzards) {
		result = append(result, position+width)
	}

	return result
}
func isInBlizzard(x, y int, blizzards []blizzard) bool {

	for _, blizzard := range blizzards {
		if x == blizzard.x && y == blizzard.y {
			return true
		}
	}

	return false
}
func contains(haystack []int, needle int) bool {

	for _, item := range haystack {
		if item == needle {
			return true
		}
	}

	return false
}
func drawGrid(grid []int, width int, blizzards []blizzard) {
	composedGrid := composeGrid(grid, width, blizzards)
	height := len(grid) / width

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := x + y*width
			if composedGrid[cell] == empty {
				fmt.Print(".")
			} else if composedGrid[cell] == wall {
				fmt.Print("#")
			} else if composedGrid[cell] == north {
				fmt.Print("^")
			} else if composedGrid[cell] == south {
				fmt.Print("v")
			} else if composedGrid[cell] == west {
				fmt.Print("<")
			} else if composedGrid[cell] == east {
				fmt.Print(">")
			} else {
				fmt.Printf("%d", composedGrid[cell])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
func composeGrid(grid []int, width int, blizzards []blizzard) []int {
	result := make([]int, len(grid))
	copy(result, grid)

	for _, b := range blizzards {
		cell := b.x + b.y*width
		if result[cell] == empty {
			result[cell] = b.direction
		} else if result[cell] == north || result[cell] == south || result[cell] == east || result[cell] == west {
			result[cell] = 2
		} else {
			result[cell]++
		}
	}

	return result
}
func stepBlizzards(width, height int, blizzards []blizzard) []blizzard {
	result := make([]blizzard, len(blizzards))

	emptySpaceWidth := width - 2
	emptySpaceHeight := height - 2

	for i := range blizzards {
		result[i] = blizzards[i]

		if result[i].direction == north {
			result[i].y += (emptySpaceHeight - 1)
		} else if result[i].direction == south {
			result[i].y++
		} else if result[i].direction == west {
			result[i].x += (emptySpaceWidth - 1)
		} else if result[i].direction == east {
			result[i].x++
		}

		result[i].x--
		result[i].x %= emptySpaceWidth
		result[i].x++

		result[i].y--
		result[i].y %= emptySpaceHeight
		result[i].y++
	}

	return result
}
func getData(filename string) ([]int, int, []blizzard) {
	grid := []int{}
	width := 0
	blizzards := []blizzard{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, r := range line {
			if r == '.' {
				grid = append(grid, empty)
			} else if r == '#' {
				grid = append(grid, wall)
			} else if r == '>' {
				grid = append(grid, empty)
				blizzards = append(blizzards, blizzard{x: x, y: y, direction: east})
			} else if r == '<' {
				grid = append(grid, empty)
				blizzards = append(blizzards, blizzard{x: x, y: y, direction: west})
			} else if r == '^' {
				grid = append(grid, empty)
				blizzards = append(blizzards, blizzard{x: x, y: y, direction: north})
			} else if r == 'v' {
				grid = append(grid, empty)
				blizzards = append(blizzards, blizzard{x: x, y: y, direction: south})
			}
		}

		if len(line) > width {
			width = len(line)
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid, width, blizzards
}
