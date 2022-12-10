package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const noop = 1
const addx = 2

func main() {
	data := getData("./input.txt")

	xValues := buildXValues(data)

	fmt.Printf("Part1: %d\n", part1(xValues))

	fmt.Printf("Part2:\n%v\n", part2(xValues))
}
func part1(xValues []int) int {
	result := 0
	for i := 20; i < len(xValues); i += 40 {
		result += i * xValues[i]
	}

	return result
}
func part2(xValues []int) string {
	width := 40
	height := 6

	result := ""

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cycle := x + y*width

			if abs((cycle%width)-xValues[cycle+1]) < 2 {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "\n"
	}

	return result
}
func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
func buildXValues(data []int) []int {
	xValues := []int{}

	currentXValue := 1
	xValues = append(xValues, currentXValue)

	for i := 0; i < len(data)/2; i++ {
		if data[i*2+0] == noop {
			xValues = append(xValues, currentXValue)
		} else if data[i*2+0] == addx {
			xValues = append(xValues, currentXValue)
			xValues = append(xValues, currentXValue)
			currentXValue += data[i*2+1]
		}
	}

	return xValues
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

		if parts[0] == "noop" {
			result = append(result, noop, 0)
		} else if parts[0] == "addx" {
			result = append(result, addx)

			value, _ := strconv.Atoi(parts[1])
			result = append(result, value)
		} else {
			log.Fatal("Failed to understand instruction")
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
