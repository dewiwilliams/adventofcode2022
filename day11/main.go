package main

import (
	"fmt"
	"sort"
)

type worryOperation func(int) int
type monkey struct {
	items       []int
	operation   worryOperation
	modulus     int
	targets     [2]int
	inspections int
}

func (m monkey) String() string {
	//return fmt.Sprintf("%v", m.items)
	return fmt.Sprintf("%v", m.inspections)
}

func main() {
	fmt.Printf("Part1: %v\n", part1(getData()))
	fmt.Printf("Part2: %v\n", part2(getData()))
}
func part1(data []monkey) int {
	return getScore(iteratePart1(data, 20))
}
func part2(data []monkey) int {
	return getScore(iteratePart2(data, 10000))
}
func getScore(m []monkey) int {
	items := []int{}
	for _, d := range m {
		items = append(items, d.inspections)
	}
	sort.Sort(sort.IntSlice(items))
	return items[len(items)-1] * items[len(items)-2]
}
func iteratePart1(m_in []monkey, rounds int) []monkey {

	m := m_in

	for i := 0; i < rounds; i++ {
		m = iterateSingleRoundPart1(m)
	}

	return m
}
func iterateSingleRoundPart1(m_in []monkey) []monkey {

	m := m_in

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i].items); j++ {
			newValue := m[i].operation(m[i].items[j])
			newValue /= 3

			if newValue%m[i].modulus == 0 {
				target := m[i].targets[0]
				m[target].items = append(m[target].items, newValue)
			} else {
				target := m[i].targets[1]
				m[target].items = append(m[target].items, newValue)
			}
		}

		m[i].inspections += len(m[i].items)
		m[i].items = []int{}
	}

	return m
}
func iteratePart2(m_in []monkey, rounds int) []monkey {

	m := m_in

	for i := 0; i < rounds; i++ {
		m = iterateSingleRoundPart2(m)
	}

	return m
}
func iterateSingleRoundPart2(m_in []monkey) []monkey {

	m := m_in

	/*
		I didn't figure this out myself, this comment on reddit game me this
			solution:
		https://www.reddit.com/r/adventofcode/comments/zihouc/comment/izrimjo

		I had noticed that the modulus numbers were all prime (for test data
			and real data), but I'm not sure that's relevant.

		I've forgotten too much number theory to have worked this out myself
			in a reasonable timeframe.
	*/

	modProduct := 1
	for i := 0; i < len(m); i++ {
		modProduct *= m[i].modulus
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i].items); j++ {
			newValue := m[i].operation(m[i].items[j])
			newValue %= modProduct

			if newValue%m[i].modulus == 0 {
				target := m[i].targets[0]
				m[target].items = append(m[target].items, newValue)
			} else {
				target := m[i].targets[1]
				m[target].items = append(m[target].items, newValue)
			}
		}

		m[i].inspections += len(m[i].items)
		m[i].items = []int{}
	}

	return m
}
func getTestData() []monkey {
	return []monkey{
		{
			items:     []int{79, 98},
			operation: func(v int) int { return v * 19 },
			modulus:   23,
			targets:   [2]int{2, 3},
		},
		{
			items:     []int{54, 65, 75, 74},
			operation: func(v int) int { return v + 6 },
			modulus:   19,
			targets:   [2]int{2, 0},
		},
		{
			items:     []int{79, 60, 97},
			operation: func(v int) int { return v * v },
			modulus:   13,
			targets:   [2]int{1, 3},
		},
		{
			items:     []int{74},
			operation: func(v int) int { return v + 3 },
			modulus:   17,
			targets:   [2]int{0, 1},
		},
	}
}
func getData() []monkey {
	return []monkey{
		{
			items:     []int{89, 74},
			operation: func(v int) int { return v * 5 },
			modulus:   17,
			targets:   [2]int{4, 7},
		},
		{
			items:     []int{75, 69, 87, 57, 84, 90, 66, 50},
			operation: func(v int) int { return v + 3 },
			modulus:   7,
			targets:   [2]int{3, 2},
		},
		{
			items:     []int{55},
			operation: func(v int) int { return v + 7 },
			modulus:   13,
			targets:   [2]int{0, 7},
		},
		{
			items:     []int{69, 82, 69, 56, 68},
			operation: func(v int) int { return v + 5 },
			modulus:   2,
			targets:   [2]int{0, 2},
		},
		{
			items:     []int{72, 97, 50},
			operation: func(v int) int { return v + 2 },
			modulus:   19,
			targets:   [2]int{6, 5},
		},
		{
			items:     []int{90, 84, 56, 92, 91, 91},
			operation: func(v int) int { return v * 19 },
			modulus:   3,
			targets:   [2]int{6, 1},
		},
		{
			items:     []int{63, 93, 55, 53},
			operation: func(v int) int { return v * v },
			modulus:   5,
			targets:   [2]int{3, 1},
		},
		{
			items:     []int{50, 61, 52, 58, 86, 68, 97},
			operation: func(v int) int { return v + 4 },
			modulus:   11,
			targets:   [2]int{5, 4},
		},
	}
}
