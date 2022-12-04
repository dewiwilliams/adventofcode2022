package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	fmt.Printf("Part1: %d\n", calculatePart1(data))
	fmt.Printf("Part2: %d\n", calculatePart2(data))
}
func calculatePart1(data []string) int {

	result := 0

	for _, line := range data {
		result += getStringScore(line)
	}

	return result
}
func calculatePart2(data []string) int {
	result := 0

	for i := 0; i < len(data)/3; i++ {
		result += getGroupScore(data[i*3 : i*3+3])
	}

	return result
}
func getGroupScore(data []string) int {

	for _, r := range data[0] {
		if strings.IndexRune(data[1], r) == -1 {
			continue
		}
		if strings.IndexRune(data[2], r) == -1 {
			continue
		}

		return getRuneScore(r)
	}

	fmt.Printf("FAILED!")
	return 0
}
func getStringScore(item string) int {
	compartment1 := item[:len(item)/2]
	compartment2 := item[len(item)/2:]

	runesSeen := make(map[rune]int)

	result := 0

	for _, r := range compartment1 {
		if strings.IndexRune(compartment2, r) == -1 {
			continue
		}
		if runesSeen[r] != 0 {
			continue
		}

		result += getRuneScore(r)
		runesSeen[r] = 1
	}

	return result
}
func getRuneScore(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r) - int('a') + 1
	}
	if r >= 'A' && r <= 'Z' {
		return int(r) - int('A') + 27
	}

	return 0
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
