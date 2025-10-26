package day12

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 12,
		Part: 2,
		Calculator: pt2,
	}
}

type axisState struct {
	p, v int
}

func pt2(input string) (string, error) {
	moons, e := parseMoons(input)
	if e != nil {
		return "", fmt.Errorf("Error parsing moon: %w", e)
	}
	xStart, yStart, zStart := getCurrentAxisStates(moons)
	xCycleLength, yCycleLength, zCycleLength := -1, -1, -1
	for i := 1; xCycleLength == -1 || yCycleLength == -1 || zCycleLength == -1; i++ {
		simOnce(moons)
		xStates, yStates, zStates := getCurrentAxisStates(moons)
		xCycleLength, yCycleLength, zCycleLength = check(xStart, xStates, i, xCycleLength), check(yStart, yStates, i, yCycleLength), check(zStart, zStates, i, zCycleLength)
	}
	total := utils.Lcm([]int{xCycleLength, yCycleLength, zCycleLength})
	return strconv.Itoa(total), nil
}

func check(start [4]axisState, newStates [4]axisState, idx, cycleLength int) int {
	if cycleLength != -1 {
		return cycleLength
	}
	if start == newStates {
		return idx
	}
	return -1
}

func getCurrentAxisStates(moons []moon) ([4]axisState, [4]axisState, [4]axisState) {
	xStates, yStates, zStates := [4]axisState{}, [4]axisState{}, [4]axisState{}
	for i, moon := range moons {
		xStates[i] = axisState{moon.px, moon.vx}
		yStates[i] = axisState{moon.py, moon.vy}
		zStates[i] = axisState{moon.pz, moon.vz}
	}
	return xStates, yStates, zStates
}