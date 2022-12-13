package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const single = 1
const multiple = 2

const incorrectOrder = 0
const correctOrder = 1
const unknown = 2

func main() {
	lines := getData("./input.txt")

	fmt.Printf("Part1: %d\n", part1(lines))
	fmt.Printf("Part2: %d\n", part2(lines))
}
func part1(in []string) int {
	result := 0

	for i := 0; i < len(in)/2; i++ {
		if inRightOrder(in[i*2+0], in[i*2+1]) == correctOrder {
			result += i + 1
		}
	}

	return result
}
func part2(in []string) int {
	packetList := append(in, "[[2]]", "[[6]]")

	sort.Slice(packetList, func(i1, i2 int) bool {
		return inRightOrder(packetList[i1], packetList[i2]) == correctOrder
	})

	l1 := getIndex(packetList, "[[2]]") + 1
	l2 := getIndex(packetList, "[[6]]") + 1

	return l1 * l2
}
func getIndex(packets []string, target string) int {
	for i, p := range packets {
		if p == target {
			return i
		}
	}
	return 0
}

func inRightOrder(s1, s2 string) int {
	if isSingleInteger(s1) && isSingleInteger(s2) {
		i1, _ := strconv.Atoi(s1)
		i2, _ := strconv.Atoi(s2)

		if i1 < i2 {
			return correctOrder
		} else if i1 > i2 {
			return incorrectOrder
		} else {
			return unknown
		}
	} else if isSingleInteger(s1) || isSingleInteger(s2) {
		if isSingleInteger(s1) {
			s1 = "[" + s1 + "]"
		} else if isSingleInteger(s2) {
			s2 = "[" + s2 + "]"
		}

		return inRightOrder(s1, s2)
	}

	parts1 := getParts(s1)
	parts2 := getParts(s2)
	limit := min(len(parts1), len(parts2))
	for i := 0; i < limit; i++ {
		result := inRightOrder(parts1[i], parts2[i])
		if result != unknown {
			return result
		}
	}

	if len(parts1) < len(parts2) {
		return correctOrder
	} else if len(parts1) > len(parts2) {
		return incorrectOrder
	}

	return unknown
}
func isSingleInteger(in string) bool {
	return in[0] != '['
}
func min(v1, v2 int) int {
	if v1 <= v2 {
		return v1
	}
	return v2
}
func getParts(in string) []string {

	if in[0] != '[' || in[len(in)-1] != ']' {
		return []string{}
	}

	contents := in[1 : len(in)-1]
	result := []string{}

	bracketCount := 0
	current := ""
	for _, r := range contents {
		if r == '[' {
			bracketCount++
		} else if r == ']' {
			bracketCount--
		}

		if bracketCount != 0 {
			current += string(r)
			continue
		}

		if r == ',' {
			result = append(result, current)
			current = ""
		} else {
			current += string(r)
		}
	}

	if len(current) > 0 {
		result = append(result, current)
	}

	return result
}

func getData(filename string) []string {
	result := []string{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line1 := scanner.Text()

		if !scanner.Scan() {
			log.Fatal("Failed to read second line!")
		}

		line2 := scanner.Text()

		result = append(result, line1, line2)

		scanner.Scan()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
