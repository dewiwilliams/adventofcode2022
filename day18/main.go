package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 * I've taken the approach here of cramming {x, y, z} in to int by some bit-shifting: x | (y<<8) | (z<<16)
 * Not sure I like that approach - it made some things easier, but meant my values weren't readable, and meant
 * that I had to move my data set a little ot make sure I was always dealing with positive integers.
 */

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

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}
func part1(cubes []cube) int {
	return getTotalExposedFaces(cubes)
}
func getTotalExposedFaces(cubes []cube) int {
	result := 0

	for _, c := range cubes {
		result += getExposedFacesForCube(c, cubes)
	}

	return result
}

func part2(cubes []cube) int {
	bounds := getBounds(cubes)

	bounds[0]--
	bounds[1]++
	bounds[2]--
	bounds[3]++
	bounds[4]--
	bounds[5]++

	toProcess := []int{toInt(bounds[0], bounds[2], bounds[4])}
	water := []int{}
	intCubes := convertCubesToIntCube(cubes)

	for len(toProcess) > 0 {
		currentCell := toProcess[len(toProcess)-1]
		toProcess = toProcess[:len(toProcess)-1]

		water = append(water, currentCell)

		neighbours := getNeighbours(currentCell)
		for _, n := range neighbours {
			if !isInBounds(n, bounds) {
				continue
			}
			if contains(intCubes, n) || contains(water, n) || contains(toProcess, n) {
				continue
			}
			toProcess = append(toProcess, n)
		}
	}

	filledGrid := fillEmptyHoles(bounds, water, cubes, intCubes)

	return getTotalExposedFaces(filledGrid)
}
func fillEmptyHoles(bounds []int, water []int, cubes []cube, intcubes []int) []cube {
	result := make([]cube, len(cubes))
	copy(result, cubes)

	for x := bounds[0]; x <= bounds[1]; x++ {
		for y := bounds[2]; y <= bounds[3]; y++ {
			for z := bounds[4]; z <= bounds[5]; z++ {

				value := toInt(x, y, z)
				if contains(water, value) || contains(intcubes, value) {
					continue
				}

				result = append(result, cube{id: len(result), x: x, y: y, z: z})
			}
		}
	}

	return result
}
func toIntCube(c cube) int {
	return toInt(c.x, c.y, c.z)
}
func toInt(x, y, z int) int {
	return x | y<<8 | z<<16
}
func getNeighbours(cell int) []int {
	return []int{
		cell + 1,
		cell - 1,
		cell + (1 << 8),
		cell - (1 << 8),
		cell + (1 << 16),
		cell - (1 << 16),
	}
}
func isInBounds(cell int, bounds []int) bool {

	x := cell & 0xFF
	y := (cell >> 8) & 0xFF
	z := (cell >> 16) & 0xFF

	return x >= bounds[0] && x <= bounds[1] && y >= bounds[2] && y <= bounds[3] && z >= bounds[4] && z <= bounds[5]
}
func convertCubesToIntCube(cubes []cube) []int {
	result := []int{}

	for _, c := range cubes {
		result = append(result, toIntCube(c))
	}

	return result
}
func contains(haystack []int, needle int) bool {

	for _, item := range haystack {
		if item == needle {
			return true
		}
	}

	return false
}
func getBounds(cubes []cube) []int {

	base := 1<<8 - 1

	minX := base - 1
	maxX := 0
	minY := base - 1
	maxY := 0
	minZ := base - 1
	maxZ := 0

	for _, c := range cubes {
		if c.x < 0 || c.x >= base || c.y < 0 || c.y >= base || c.z < 0 || c.z >= base {
			fmt.Println("Out of bounds!")
			os.Exit(1)
		}

		if c.x > maxX {
			maxX = c.x
		}
		if c.x < minX {
			minX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
		if c.y < minY {
			minY = c.y
		}
		if c.z > maxZ {
			maxZ = c.z
		}
		if c.z < minZ {
			minZ = c.z
		}
	}

	return []int{minX, maxX, minY, maxY, minZ, maxZ}
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
			x:  x + 16,
			y:  y + 16,
			z:  z + 16,
		}
		result = append(result, c)

		lineIndex++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
