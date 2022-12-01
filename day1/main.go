package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := getData()
	loads := getLoads(data)
	sort.Ints(loads)

	fmt.Printf("Part 1: %d\n", loads[len(loads)-1])
	fmt.Printf("Part 2: %d\n", loads[len(loads)-1]+loads[len(loads)-2]+loads[len(loads)-3])
}
func getLoads(data []int) []int {
	result := []int{}

	currentValue := 0
	for _, v := range data {
		if v == -1 {
			result = append(result, currentValue)
			currentValue = 0
			continue
		}

		currentValue += v
	}

	return result
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

		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			result = append(result, -1)
			continue
		}

		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		result = append(result, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
