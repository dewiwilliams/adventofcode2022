package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type cube struct {
	id int
	x  int
	y  int
	z  int
}

func (c cube) adjacentTo(c1 cube) bool {

	if c.x == c1.x && c.y == c1.y && abs(c.z-c1.z) == 1 {
		return true
	}
	if c.z == c1.z && c.y == c1.y && abs(c.x-c1.x) == 1 {
		return true
	}
	if c.x == c1.x && c.z == c1.z && abs(c.y-c1.y) == 1 {
		return true
	}

	return false
}

func main() {
	data := getData("./input.txt")

	fmt.Printf("Got data: %v\n", data)
	fmt.Printf("Part 1: %d\n", part1(data))
}
func part1(cubes []cube) int {
	result := 0

	for _, c := range cubes {
		result += getExposedFacesForCube(c, cubes)
	}

	return result
}
func getExposedFacesForCube(target cube, cubes []cube) int {
	result := 6

	for _, c := range cubes {
		if c.id == target.id {
			continue
		}

		if target.adjacentTo(c) {
			result--
		}
	}

	return result
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func getData(filename string) []cube {
	result := []cube{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineIndex := 0
	for scanner.Scan() {

		line := scanner.Text()

		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		c := cube{
			id: lineIndex,
			x:  x,
			y:  y,
			z:  z,
		}
		result = append(result, c)

		lineIndex++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
