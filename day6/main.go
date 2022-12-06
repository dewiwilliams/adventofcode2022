package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	data := getData()

	fmt.Printf("Part1: %d\n", part1(data))
	fmt.Printf("Part2: %d\n", part2(data))
}
func part1(data string) int {
	for i := 3; i < len(data); i++ {
		if isStartOfMarker(data[i-3 : i+1]) {
			return i + 1
		}
	}

	return -1
}
func part2(data string) int {
	for i := 13; i < len(data); i++ {
		if isStartOfMarker(data[i-13 : i+1]) {
			return i + 1
		}
	}

	return -1
}
func isStartOfMarker(data string) bool {
	runesSeen := make(map[rune]int)

	for _, r := range data {
		if runesSeen[r] != 0 {
			return false
		}

		runesSeen[r] = 1
	}

	return true
}
func getData() string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return line
}
