package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type monkey struct {
	value     int
	monkey1   string
	operation string
	monkey2   string
}

func main() {
	data := getData("./input.txt")

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}
func part2(data map[string]monkey) int {
	low, high := findRange(data)

	roughValue := binarySearch(data, low, high)
	return findExactHumnValue(data, roughValue)
}
func findExactHumnValue(data map[string]monkey, value int) int {
	for i := 0; ; i++ {
		currentValue := value + i/2
		if i%2 == 0 {
			currentValue = value - i/2
		}

		result, valid := getDiff(data, currentValue)
		if result == 0 && valid {
			return currentValue
		}
	}
}
func binarySearch(data map[string]monkey, low, high int) int {

	for {
		midpoint := (low + high) / 2
		lowValue, _ := getDiff(data, low)
		midpointValue, _ := getDiff(data, midpoint)
		highValue, _ := getDiff(data, high)

		if midpointValue == 0 {
			return midpoint
		}

		if differentSigns(lowValue, midpointValue) {
			high = midpoint
		} else if differentSigns(midpointValue, highValue) {
			low = midpoint
		}
	}
}
func getDiff(data map[string]monkey, value int) (int, bool) {
	humnValue := data["humn"]
	humnValue.value = value
	data["humn"] = humnValue

	value1, valid1 := getValue(data, data["root"].monkey1)
	value2, valid2 := getValue(data, data["root"].monkey2)
	valid := valid1 && valid2

	return value1 - value2, valid
}
func findRange(data map[string]monkey) (int, int) {
	currentValue := 1
	for {
		diff1, _ := getDiff(data, currentValue)
		diff2, _ := getDiff(data, currentValue*2)

		if differentSigns(diff1, diff2) {
			return currentValue, currentValue * 2
		}

		currentValue *= 2
	}
}
func differentSigns(diff1, diff2 int) bool {
	if diff1 > 0 && diff2 < 0 {
		return true
	}
	if diff1 < 0 && diff2 > 0 {
		return true
	}

	return false
}
func part1(data map[string]monkey) int {
	result, _ := getValue(data, "root")
	return result
}
func getValue(data map[string]monkey, key string) (int, bool) {
	monkey := data[key]
	if len(monkey.operation) == 0 {
		return data[key].value, true
	}

	value1, valid1 := getValue(data, monkey.monkey1)
	value2, valid2 := getValue(data, monkey.monkey2)
	valid := true

	if !valid1 || !valid2 {
		valid = false
	}

	if monkey.operation == "+" {
		return value1 + value2, valid
	}
	if monkey.operation == "-" {
		return value1 - value2, valid
	}
	if monkey.operation == "*" {
		return value1 * value2, valid
	}
	if monkey.operation == "/" {
		if value1%value2 != 0 {
			return value1 / value2, false
		}
		return value1 / value2, valid
	}

	return 0, false
}
func getData(filename string) map[string]monkey {
	result := make(map[string]monkey)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		parts := strings.Split(line, " ")
		monkeyName := parts[0][:4]

		if len(parts) == 2 {
			value, _ := strconv.Atoi(parts[1])
			result[monkeyName] = monkey{
				value: value,
			}
		} else {
			result[monkeyName] = monkey{
				monkey1:   parts[1],
				operation: parts[2],
				monkey2:   parts[3],
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
