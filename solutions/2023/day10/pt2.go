package day10

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 10,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	grid := utils.GridFromString(input)
	startingPoint, e := grid.Find('S').OrErr("Input does not have starting point!")
	if e != nil {
		return "", e
	}
	totalSteps, vertices, e := FindLoopCountSteps(grid, startingPoint)
	if e != nil {
		return "", fmt.Errorf("Error getting total steps in loop: %w", e)
	}
	total := shoelace(vertices) - totalSteps / 2 + 1
	return strconv.Itoa(total), nil
}

func shoelace(pts []utils.Point) int {
	n := len(pts)
	if n < 3 {
		return 0
	}
	twiceArea := 0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		x1, y1 := pts[i].X, pts[i].Y
		x2, y2 := pts[j].X, pts[j].Y
		twiceArea += x1*y2 - x2*y1
	}
	if twiceArea < 0 {
		twiceArea = -twiceArea
	}
	A := twiceArea / 2
	return A
}