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

	fmt.Printf("Part1: %d\n", part1(data))
}
func part1(data map[string]monkey) int {
	return getValue(data, "root")
}
func getValue(data map[string]monkey, key string) int {
	monkey := data[key]
	if len(monkey.operation) == 0 {
		return data[key].value
	}

	value1 := getValue(data, monkey.monkey1)
	value2 := getValue(data, monkey.monkey2)

	if monkey.operation == "+" {
		return value1 + value2
	}
	if monkey.operation == "-" {
		return value1 - value2
	}
	if monkey.operation == "*" {
		return value1 * value2
	}
	if monkey.operation == "/" {
		return value1 / value2
	}

	return 0
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
