package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "./input.txt"
	data := getData(filename)

	fmt.Printf("Part 1: %s\n", part1(data))
}
func part1(data []string) string {
	result := 0

	for _, line := range data {
		result += toDecimal(line)
	}

	return toSnafu(result, 20, "")
}
func toSnafu(decimal int, column int, current string) string {
	if column == -1 {
		return current
	}

	columnValue := pow(5, column)
	maxPossibleValue := getMaxPossibleValue(column - 1)
	if decimal > maxPossibleValue {
		if decimal > columnValue+maxPossibleValue {
			return toSnafu(decimal-2*columnValue, column-1, current+"2")
		}

		return toSnafu(decimal-1*columnValue, column-1, current+"1")
	} else if decimal >= 0 {
		columnResult := decimal / columnValue
		if columnResult == 0 && len(current) == 0 {
			return toSnafu(decimal, column-1, "")
		}

		return toSnafu(decimal-columnResult*columnValue, column-1, fmt.Sprintf("%s%d", current, columnResult))
	} else if decimal < 0 {

		if decimal+maxPossibleValue >= 0 {
			return toSnafu(decimal, column-1, current+"0")
		} else if decimal+columnValue+maxPossibleValue >= 0 {
			return toSnafu(decimal+columnValue, column-1, current+"-")
		} else if decimal+2*columnValue+maxPossibleValue >= 0 {
			return toSnafu(decimal+2*columnValue, column-1, current+"=")
		}
	}

	log.Fatal("Unhandled state")
	return ""
}
func getMaxPossibleValue(column int) int {
	if column == 0 {
		return 2
	}
	if column < 0 {
		return 0
	}
	return 2*pow(5, column) + getMaxPossibleValue(column-1)
}
func toDecimal(snafu string) int {
	result := 0

	for i, r := range snafu {
		columnIndex := len(snafu) - i - 1
		columnValue := pow(5, columnIndex)

		if r == '2' {
			result += 2 * columnValue
		} else if r == '1' {
			result += columnValue
		} else if r == '0' {
		} else if r == '-' {
			result -= columnValue
		} else if r == '=' {
			result -= 2 * columnValue
		}
	}

	return result
}
func pow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
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
		line := scanner.Text()

		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
