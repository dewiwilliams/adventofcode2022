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
	data := getData("./input.txt")

	//fmt.Printf("Part 1: %d\n", part1(data, 10))
	fmt.Printf("Part 1: %d\n", part1(data, 2000000))
}
func part1(data []int, row int) int {
	//printGrid(-4, 26, 9, 11, data)

	minX, maxX := getBeaconXRange(data)

	return getLitCells(minX, maxX, row, data)
}
func printGrid(minX, maxX, minY, maxY int, data []int) {
	fmt.Println()

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if isCellOnBeaconOrSensor(x, y, data) {
				fmt.Print("@")
			} else if isCellLit(x, y, data) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}
func getBeaconXRange(data []int) (int, int) {
	minX := 0
	maxX := 0

	for i := 0; i < len(data)/4; i++ {
		beaconDistance := distance(data[i*4+0], data[i*4+1], data[i*4+2], data[i*4+3])

		sensorMinX := data[i*4+0] - beaconDistance
		sensorMaxX := data[i*4+0] + beaconDistance

		if sensorMinX < minX {
			minX = sensorMinX
		}
		if sensorMaxX > maxX {
			maxX = sensorMaxX
		}
	}

	return minX, maxX
}
func getBiggestBeaconSensorDistance(data []int) int {
	result := 0

	for i := 0; i < len(data)/4; i++ {
		beaconDistance := distance(data[i*4+0], data[i*4+1], data[i*4+2], data[i*4+3])

		if beaconDistance > result {
			result = beaconDistance
		}
	}

	return result
}
func getLitCells(minX, maxX, y int, data []int) int {
	result := 0

	for x := minX; x <= maxX; x++ {
		if !isCellOnBeaconOrSensor(x, y, data) && isCellLit(x, y, data) {
			result++
		}
	}

	return result
}
func isCellOnBeaconOrSensor(x, y int, data []int) bool {
	for i := 0; i < len(data)/4; i++ {
		if x == data[i*4+0] && y == data[i*4+1] {
			return true
		}
		if x == data[i*4+2] && y == data[i*4+3] {
			return true
		}
	}

	return false
}
func isCellLit(x, y int, data []int) bool {
	for i := 0; i < len(data)/4; i++ {
		cellDistance := distance(x, y, data[i*4+0], data[i*4+1])
		beaconDistance := distance(data[i*4+0], data[i*4+1], data[i*4+2], data[i*4+3])

		if cellDistance <= beaconDistance {
			return true
		}
	}

	return false
}
func distance(x1, y1, x2, y2 int) int {
	xdiff := x2 - x1
	ydiff := y2 - y1

	return abs(xdiff) + abs(ydiff)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
		result = append(result, parseLine(line)...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
func parseLine(line string) []int {
	result := []int{}

	parts := strings.Split(line, " ")
	x1, _ := strconv.Atoi(parts[2][2 : len(parts[2])-1])
	y1, _ := strconv.Atoi(parts[3][2 : len(parts[3])-1])
	x2, _ := strconv.Atoi(parts[8][2 : len(parts[8])-1])
	y2, _ := strconv.Atoi(parts[9][2:])
	result = append(result, x1, y1, x2, y2)

	return result
}
