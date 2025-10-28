package day15

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 15,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	grid, e := dfs(input)
	unoxygenated := grid.Find(2).OrPanic("Should be 1 oxygen square")
	steps := bfs(grid, unoxygenated)
	maxSteps := 0
	for _, step := range steps {
		maxSteps = max(maxSteps, step)
	}
	return strconv.Itoa(maxSteps), e
}