package day12

import (
	"advent-of-go/utils"
	"fmt"
	"slices"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 12,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
		lines := utils.GetLines(input)
	total := 0
	for _, line := range lines {
		states, layout, e := processLine(line)
		if e != nil {
			return "", fmt.Errorf("Error processing line: %w", e)
		}
		states, layout = multiplyLine(states, layout)
		states = append(states, Off)
		numWays := calculateNumWays(states, layout)
		total += numWays
	}
	return strconv.Itoa(total), nil
}

func multiplyLine(states []State, layout []int) ([]State, []int) {
	newStates := []State{}
	newLayout := []int{}
	for i := 0; i < 5; i++ {
		newStates = slices.Concat(newStates, states)
		newStates = append(newStates, Unknown)
		newLayout = slices.Concat(newLayout, layout)
	}
	return newStates[:len(newStates) - 1], newLayout
}