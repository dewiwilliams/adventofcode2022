package main

import (
	"fmt"
	"testing"
)

func TestTick(t *testing.T) {
	// 1 ore-collecting robot collects 1 ore; you now have 2 ore.
	// 2 clay-collecting robots collect 2 clay; you now have 4 clay.

	state := state{
		ore:            1,
		oreRobots:      1,
		clay:           2,
		clayRobots:     2,
		obsidian:       3,
		obsidianRobots: 4,
		geode:          4,
		geodeRobots:    5,
	}

	newState := tick(state)
	if newState.ore != 2 || newState.clay != 4 || newState.obsidian != 7 || newState.geode != 9 {
		t.Errorf("Wrong quantities!")
	}
}
func TestGetPossibleNextStates(t *testing.T) {
	// Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.

	// 1 ore-collecting robot collects 1 ore; you now have 4 ore.
	// 3 clay-collecting robots collect 3 clay; you now have 15 clay.

	blueprint := blueprint{
		oreRobotCost:      4,
		clayRobotCost:     2,
		obsidianRobotCost: [2]int{3, 14},
		geodeRobotCost:    [2]int{2, 7},
	}
	state := state{
		ore:            4,
		oreRobots:      1,
		clay:           15,
		clayRobots:     3,
		obsidian:       0,
		obsidianRobots: 0,
		geode:          0,
		geodeRobots:    0,
	}

	possibleStates := getPossibleNextStates(blueprint, state)

	fmt.Printf("Possible states: %d\n", possibleStates)

	//if len(possibleStates) !=
}
