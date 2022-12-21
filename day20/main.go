package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	data := getData("./input.txt")

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}
func part2(data []int) int {
	workspace := make([]int, len(data))
	copy(workspace, data)

	ordering := []int{}
	for i := range data {
		ordering = append(ordering, i)
		workspace[i] *= 811589153
	}

	for j := 0; j < 10; j++ {
		for i := range data {
			index := indexOf(ordering, i)
			value := workspace[index]
			if value == 0 {
				continue
			}

			workspace = removeInt(workspace, index)
			ordering = removeInt(ordering, index)
			index += value

			if index < 0 {
				index += (-index / len(workspace)) * len(workspace)
				index += len(workspace)
			}
			index %= len(workspace)

			workspace = insertInt(workspace, value, index)
			ordering = insertInt(ordering, i, index)
		}
	}

	zeroPosition := indexOf(workspace, 0)

	return getItemAtIndex(workspace, zeroPosition+1000) +
		getItemAtIndex(workspace, zeroPosition+2000) +
		getItemAtIndex(workspace, zeroPosition+3000)
}
func part1(data []int) int {
	workspace := make([]int, len(data))
	copy(workspace, data)

	ordering := []int{}
	for i := range data {
		ordering = append(ordering, i)
	}

	for i, n := range data {
		if n == 0 {
			continue
		}

		index := indexOf(ordering, i)

		workspace = removeInt(workspace, index)
		ordering = removeInt(ordering, index)
		index += n
		for index < 0 {
			index += len(workspace)
		}
		index %= len(workspace)

		workspace = insertInt(workspace, n, index)
		ordering = insertInt(ordering, i, index)
	}

	zeroPosition := indexOf(workspace, 0)

	return getItemAtIndex(workspace, zeroPosition+1000) +
		getItemAtIndex(workspace, zeroPosition+2000) +
		getItemAtIndex(workspace, zeroPosition+3000)
}
func getItemAtIndex(data []int, index int) int {
	index %= len(data)
	return data[index]
}
func insertInt(array []int, value int, index int) []int {
	return append(array[:index], append([]int{value}, array[index:]...)...)
}

func removeInt(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func moveInt(array []int, srcIndex int, dstIndex int) []int {
	value := array[srcIndex]
	return insertInt(removeInt(array, srcIndex), value, dstIndex)
}
func indexOf(data []int, item int) int {
	for i, v := range data {
		if v == item {
			return i
		}
	}

	return 0
}
func getData(filename string) []int {
	result := []int{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		value, _ := strconv.Atoi(line)

		result = append(result, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
func max(v1, v2 int) int {
	if v1 >= v2 {
		return v1
	}
	return v2
}
