package day4

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        4,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	grid := utils.GridFromString(input)
	total := findAccessible(grid, false)
	return strconv.Itoa(total), nil
}

func findAccessible(grid utils.FinGrid[rune], remove bool) int {
	total := 0
	for p := range grid.Points() {
		if grid.AtP(p).Is('@') && isAccessible(grid, p) {
			total++
			if remove {
				grid.SetP(p, '.')
			}
		}
	}
	return total
}

func isAccessible(grid utils.FinGrid[rune], p utils.Point) bool {
	total := 0
	for char := range grid.AdjC(p) {
		if char.Is('@') {
			total++
		}
	}
	return total < 4
}
