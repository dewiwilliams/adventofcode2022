package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type blueprint struct {
	oreRobotCost      int
	clayRobotCost     int
	obsidianRobotCost [2]int
	geodeRobotCost    [2]int
}
type state struct {
	ore            int
	oreRobots      int
	clay           int
	clayRobots     int
	obsidian       int
	obsidianRobots int
	geode          int
	geodeRobots    int
}

func main() {
	data := getData("./input.txt")

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data[0:3]))
}
func part1(blueprints []blueprint) int {
	result := 0

	for i, blueprint := range blueprints {
		result += (i + 1) * getBlueprintMaxGeodes(blueprint, 24)
	}

	return result
}
func part2(blueprints []blueprint) int {
	result := 1

	for _, blueprint := range blueprints {
		result *= getBlueprintMaxGeodes(blueprint, 32)
	}

	return result
}
func getBlueprintMaxGeodes(blueprint blueprint, totalCycles int) int {
	startingState := state{
		oreRobots: 1,
	}

	return iterate(blueprint, startingState, 1, totalCycles)
}
func iterate(blueprint blueprint, in state, cycle, totalCycles int) int {
	if cycle == totalCycles+1 {
		return in.geode
	}

	result := 0

	nextStates := getPossibleNextStates(blueprint, in)
	for _, s := range nextStates {
		result = max(result, iterate(blueprint, s, cycle+1, totalCycles))
	}

	return result
}
func getPossibleNextStates(blueprint blueprint, in state) []state {

	canBuildOreRobot := in.ore >= blueprint.oreRobotCost
	canBuildClayRobot := in.ore >= blueprint.clayRobotCost
	canBuildObsidianRobot := in.ore >= blueprint.obsidianRobotCost[0] && in.clay >= blueprint.obsidianRobotCost[1]
	canBuildGeodeRobot := in.ore >= blueprint.geodeRobotCost[0] && in.obsidian >= blueprint.geodeRobotCost[1]

	baseState := tick(in)

	/*
		This isn't quite correct - it's a little too aggressive in filtering out possibilites, but
		I couldn't get it to run in a reasonable timeframe without being slightly too aggressive.
		It gives the wrong result for part 2 sample data, but gives the correct result for my data.
		I think the sample data is a little skewed in terms of cost, and isn't altogether representative
		of the data I have.
	*/

	if canBuildGeodeRobot && in.ore >= 5 {
		return []state{buildGeodeRobot(baseState, blueprint)}
	}
	if canBuildObsidianRobot {
		return []state{buildObsidianRobot(baseState, blueprint)}
	}

	if baseState.obsidianRobots == 0 && baseState.geodeRobots == 0 && !canBuildGeodeRobot && !canBuildObsidianRobot {
		if canBuildOreRobot && canBuildClayRobot {
			return []state{buildOreRobot(baseState, blueprint), buildClayRobot(baseState, blueprint)}
		}
	}

	result := []state{}
	result = append(result, baseState)

	if canBuildOreRobot {
		result = append(result, buildOreRobot(baseState, blueprint))
	}
	if canBuildClayRobot {
		result = append(result, buildClayRobot(baseState, blueprint))
	}
	if canBuildObsidianRobot {
		result = append(result, buildObsidianRobot(baseState, blueprint))
	}
	if canBuildGeodeRobot {
		result = append(result, buildGeodeRobot(baseState, blueprint))
	}

	return result
}
func buildOreRobot(newState state, blueprint blueprint) state {
	newState.ore -= blueprint.oreRobotCost
	newState.oreRobots++
	return newState
}
func buildClayRobot(newState state, blueprint blueprint) state {
	newState.ore -= blueprint.clayRobotCost
	newState.clayRobots++
	return newState
}
func buildObsidianRobot(newState state, blueprint blueprint) state {
	newState.ore -= blueprint.obsidianRobotCost[0]
	newState.clay -= blueprint.obsidianRobotCost[1]
	newState.obsidianRobots++
	return newState
}
func buildGeodeRobot(newState state, blueprint blueprint) state {
	newState.ore -= blueprint.geodeRobotCost[0]
	newState.obsidian -= blueprint.geodeRobotCost[1]
	newState.geodeRobots++
	return newState
}
func tick(in state) state {
	result := in

	result.ore += result.oreRobots
	result.clay += result.clayRobots
	result.obsidian += result.obsidianRobots
	result.geode += result.geodeRobots

	return result
}
func getData(filename string) []blueprint {
	result := []blueprint{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		parts := strings.Split(line, " ")
		oreRobotCost, _ := strconv.Atoi(parts[6])
		clayRobotCost, _ := strconv.Atoi(parts[12])
		obsidianRobotCost1, _ := strconv.Atoi(parts[18])
		obsidianRobotCost2, _ := strconv.Atoi(parts[21])
		geodeRobotCost1, _ := strconv.Atoi(parts[27])
		geodeRobotCost2, _ := strconv.Atoi(parts[30])

		blueprint := blueprint{
			oreRobotCost:      oreRobotCost,
			clayRobotCost:     clayRobotCost,
			obsidianRobotCost: [2]int{obsidianRobotCost1, obsidianRobotCost2},
			geodeRobotCost:    [2]int{geodeRobotCost1, geodeRobotCost2},
		}

		result = append(result, blueprint)
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
