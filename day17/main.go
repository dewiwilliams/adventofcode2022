package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	data := getData("./input.txt")

	fmt.Printf("Part 1: %d\n", part1(data))
}
func getPiece(index int) []bool {
	if index%5 == 0 {
		return []bool{
			true, true, true, true,
		}
	}
	if index%5 == 1 {
		return []bool{
			false, true, false, false,
			true, true, true, false,
			false, true, false, false,
		}
	}
	if index%5 == 2 {
		return []bool{
			true, true, true, false,
			false, false, true, false,
			false, false, true, false,
		}
	}
	if index%5 == 3 {
		return []bool{
			true, false, false, false,
			true, false, false, false,
			true, false, false, false,
			true, false, false, false,
		}
	}
	if index%5 == 4 {
		return []bool{
			true, true, false, false,
			true, true, false, false,
		}
	}

	return []bool{}
}

func part1(data []int) int {
	grid := extendGrid([]bool{})
	windIndex := 0

	for i := 0; i < 2022; i++ {
		grid = extendGrid(grid)
		windIndex = processNextPiece(grid, i, data, windIndex)
	}

	//renderGrid(grid)

	return getFilledRows(grid)
}
func processNextPiece(grid []bool, pieceIndex int, wind []int, windIndex int) int {
	filledRows := getFilledRows(grid)
	row := filledRows + 3
	column := 2
	piece := getPiece(pieceIndex)

	for {
		if canInsertPiece(grid, piece, column+wind[windIndex], row) {
			column += wind[windIndex]
		}
		windIndex++
		windIndex %= len(wind)

		if !canInsertPiece(grid, piece, column, row-1) {
			insertPieceInGrid(grid, piece, column, row)
			break
		}
		row--
	}

	return windIndex
}
func canInsertPiece(grid []bool, piece []bool, x, y int) bool {
	if y < 0 || x < 0 {
		return false
	}

	rows := len(piece) / 4
	for row := 0; row < rows; row++ {
		for i := 0; i < 4; i++ {
			if !piece[i+row*4] {
				continue
			}
			if x+i >= 7 {
				return false
			}

			index := (x + i) + (y+row)*7

			if grid[index] {
				return false
			}
		}
	}

	return true
}
func insertPieceInGrid(grid []bool, piece []bool, x, y int) {
	rows := len(piece) / 4
	for row := 0; row < rows; row++ {
		for i := 0; i < 4; i++ {

			if !piece[i+row*4] {
				continue
			}

			index := (x + i) + (y+row)*7
			grid[index] = true
		}
	}
}
func extendGrid(grid []bool) []bool {
	filledRows := getFilledRows(grid)
	totalRows := len(grid) / 7
	emptyRows := totalRows - filledRows

	for i := 0; i < 8-emptyRows; i++ {
		grid = append(grid, false, false, false, false, false, false, false)
	}

	return grid
}
func getFilledRows(grid []bool) int {
	rows := len(grid) / 7
	for i := 0; i < rows; i++ {
		if isRowEmpty(grid, i) {
			return i
		}
	}

	return rows
}
func isRowEmpty(grid []bool, row int) bool {
	offset := row * 7
	for i := 0; i < 7; i++ {
		if grid[offset+i] {
			return false
		}
	}
	return true
}
func renderGrid(grid []bool) {
	rows := len(grid) / 7
	for i := rows - 1; i >= 0; i-- {
		fmt.Print("|")
		for j := 0; j < 7; j++ {
			index := i*7 + j
			if grid[index] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("|\n")
	}
	fmt.Println("+-------+")
}
func getData(filename string) []int {
	result := []int{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	for _, r := range line {
		if r == '<' {
			result = append(result, -1)
		} else {
			result = append(result, 1)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
