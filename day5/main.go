package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}
func part2() {
	state, instructions := getData()

	for i := 0; i < len(instructions)/3; i++ {
		applyInstructionsPart2(state, instructions[i*3+0], instructions[i*3+1], instructions[i*3+2])
	}

	printStateSolution(state)
}
func applyInstructionsPart2(state [][]rune, quantity, start, end int) {
	for i := 0; i < quantity; i++ {
		mover := state[start][len(state[start])-quantity+i]
		state[end] = append(state[end], mover)
	}

	state[start] = state[start][:len(state[start])-quantity]
}
func part1() {
	state, instructions := getData()

	for i := 0; i < len(instructions)/3; i++ {
		applyInstructionsPart1(state, instructions[i*3+0], instructions[i*3+1], instructions[i*3+2])
	}

	printStateSolution(state)
}
func printStateSolution(state [][]rune) {
	for i := 0; i < 9; i++ {
		fmt.Print(string(state[i][len(state[i])-1]))
	}
	fmt.Print("\n")
}
func applyInstructionsPart1(state [][]rune, quantity, start, end int) {
	for i := 0; i < quantity; i++ {
		mover := state[start][len(state[start])-1]
		state[start] = state[start][:len(state[start])-1]
		state[end] = append(state[end], mover)
	}
}
func getData() ([][]rune, []int) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	startingState := getStartingState(scanner)
	instructions := getInstructions(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return startingState, instructions
}
func getStartingState(scanner *bufio.Scanner) [][]rune {
	result := [][]rune{{}}
	for i := 0; i < 9; i++ {
		result = append(result, []rune{})
	}

	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		line = PadRight(line, 35, ' ')
		lines = append(lines, line)
	}

	lines = lines[:len(lines)-1]

	for i := len(lines) - 1; i >= 0; i-- {

		line := []rune(lines[i])
		handleParsingCrate(result, 0, line[1])
		handleParsingCrate(result, 1, line[5])
		handleParsingCrate(result, 2, line[9])
		handleParsingCrate(result, 3, line[13])
		handleParsingCrate(result, 4, line[17])
		handleParsingCrate(result, 5, line[21])
		handleParsingCrate(result, 6, line[25])
		handleParsingCrate(result, 7, line[29])
		handleParsingCrate(result, 8, line[33])
	}

	return result
}
func PadRight(str string, length int, pad byte) string {
	if len(str) >= length {
		return str
	}
	buf := bytes.NewBufferString(str)
	for i := 0; i < length-len(str); i++ {
		buf.WriteByte(pad)
	}
	return buf.String()
}
func handleParsingCrate(target [][]rune, index int, r rune) {
	if r == ' ' {
		return
	}
	target[index] = append(target[index], r)
}
func getInstructions(scanner *bufio.Scanner) []int {
	result := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		quantity, _ := strconv.Atoi(parts[1])
		start, _ := strconv.Atoi(parts[3])
		end, _ := strconv.Atoi(parts[5])

		result = append(result, quantity, start-1, end-1)
	}

	return result
}
