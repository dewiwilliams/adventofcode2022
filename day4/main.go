package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const rock = 1
const paper = 2
const scissors = 3

const lose = 1
const draw = 2
const win = 3

func main() {
	data := getData()

	fmt.Printf("Part1: %d\n", getPart1(data))
	fmt.Printf("Part1: %d\n", getPart2(data))
}
func getPart1(data []int) int {
	result := 0

	for i := 0; i < len(data)/4; i++ {
		if completeOverlap(data[i*4+0], data[i*4+1], data[i*4+2], data[i*4+3]) {
			result++
		}
	}

	return result
}
func getPart2(data []int) int {
	result := 0

	for i := 0; i < len(data)/4; i++ {
		if partialOverlap(data[i*4+0], data[i*4+1], data[i*4+2], data[i*4+3]) {
			result++
		}
	}

	return result
}
func partialOverlap(e1start, e1end, e2start, e2end int) bool {
	if e1start >= e2start && e1start <= e2end {
		return true
	}
	if e1end >= e2start && e1end <= e2end {
		return true
	}
	if e2start >= e1start && e2start <= e1end {
		return true
	}
	if e2end >= e1start && e2end <= e1end {
		return true
	}

	return false
}
func completeOverlap(e1start, e1end, e2start, e2end int) bool {
	if e1start >= e2start && e1end <= e2end {
		return true
	}
	if e2start >= e1start && e2end <= e1end {
		return true
	}

	return false
}
func getData() []int {
	result := []int{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		elves := strings.Split(line, ",")
		elf1Parts := strings.Split(elves[0], "-")
		elf2Parts := strings.Split(elves[1], "-")

		result = append(result, parseString(elf1Parts[0]))
		result = append(result, parseString(elf1Parts[1]))
		result = append(result, parseString(elf2Parts[0]))
		result = append(result, parseString(elf2Parts[1]))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
func parseString(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return v
}
