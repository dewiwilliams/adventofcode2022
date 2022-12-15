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

	fmt.Printf("Part 1: %d\n", part1(data, 2000000))
	fmt.Printf("Part 2: %d\n", part2(0, 4000000, data))
}
func part1(data []int, row int) int {
	minX, maxX := getBeaconXRange(data)

	return getLitCells(minX, maxX, row, data)
}
func part2(min, max int, data []int) int {

	/*
	 I've defined the sensor 'range' to be the distance from a sensor to the beacon it detected.
	 The sensor remaining range for a given point is the sensor range - distance from the point to the sensor.
	 I split the search area up in to stepSize*stepSize blocks, and look for the largest remaining sensor range
	 for the midpoint of that block. If the remaining range is bigger than stepSize, I know the empty cell isn't
	 in this block and can skip it.

	 There is an additional optimisation to be done here where each block is broken down again in to smaller
	 sub-blocks, but this solution gave me the answer in a <30s, so didn't worry about it.
	*/

	stepSize := 500
	halfStepSize := stepSize/2 + 1

	for y := min; y <= max; y += stepSize {
		for x := min; x <= max; x += stepSize {
			midPointX := x + halfStepSize
			midPointY := y + halfStepSize

			if getBiggestSensorRemainingDistance(midPointX, midPointY, data) > stepSize {
				continue
			}

			value := searchArea(x, x+stepSize, y, y+stepSize, data)
			if value != 0 {
				return value
			}
		}
	}

	return 0
}
func searchArea(minX, maxX, minY, maxY int, data []int) int {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if isCellOnBeaconOrSensor(x, y, data) || isCellLit(x, y, data) {
				continue
			}
			return x*4000000 + y
		}
	}

	return 0
}
func getBiggestSensorRemainingDistance(x, y int, data []int) int {
	result := 0

	for i := 0; i < len(data)/4; i++ {
		beaconRange := distance(data[i*4+0], data[i*4+1], data[i*4+2], data[i*4+3])
		distanceToBeacon := distance(x, y, data[i*4+0], data[i*4+1])

		remainingDistance := beaconRange - distanceToBeacon
		if remainingDistance > result {
			result = remainingDistance
		}
	}

	return result
}
func printGrid(minX, maxX, minY, maxY int, data []int) {
	fmt.Println()

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if isCellOnBeacon(x, y, data) {
				fmt.Print("B")
			} else if isCellOnSensor(x, y, data) {
				fmt.Print("S")
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
func isCellOnBeacon(x, y int, data []int) bool {
	for i := 0; i < len(data)/4; i++ {
		if x == data[i*4+0] && y == data[i*4+1] {
			return true
		}
	}

	return false
}
func isCellOnSensor(x, y int, data []int) bool {
	for i := 0; i < len(data)/4; i++ {
		if x == data[i*4+2] && y == data[i*4+3] {
			return true
		}
	}

	return false
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
