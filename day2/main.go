package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
func calculatePart1(data []int) int {
	result := 0

	for i := 0; i < len(data)/2; i++ {
		result += getScoreForPlayer2(data[i*2+0], data[i*2+1])
	}

	return result
}
func calculatePart2(data []int) int {
	result := 0

	for i := 0; i < len(data)/2; i++ {
		p1 := data[i*2+0]
		directive := data[i*2+1]

		result += getScoreForPlayer2(p1, getDesiredPlay(p1, directive))
	}

	return result
}
func getDesiredPlay(p1, directive int) int {
	if directive == lose {
		return getLosingPlay(p1)
	} else if directive == win {
		return getWinningPlay(p1)
	}

	return p1
}
func getWinningPlay(p1 int) int {
	if p1 == rock {
		return paper
	} else if p1 == paper {
		return scissors
	} else if p1 == scissors {
		return rock
	}

	return 0
}
func getLosingPlay(p1 int) int {
	if p1 == rock {
		return scissors
	} else if p1 == paper {
		return rock
	} else if p1 == scissors {
		return paper
	}
	return 0
}
func getScoreForPlayer2(h1, h2 int) int {
	if h1 == h2 {
		return 3 + h2
	}

	if h1 == rock && h2 == paper {
		return 6 + h2
	}
	if h1 == paper && h2 == scissors {
		return 6 + h2
	}
	if h1 == scissors && h2 == rock {
		return 6 + h2
	}

	return h2
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

		if line[0] == 'A' {
			result = append(result, rock)
		} else if line[0] == 'B' {
			result = append(result, paper)
		} else if line[0] == 'C' {
			result = append(result, scissors)
		}

		if line[2] == 'X' {
			result = append(result, rock)
		} else if line[2] == 'Y' {
			result = append(result, paper)
		} else if line[2] == 'Z' {
			result = append(result, scissors)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
