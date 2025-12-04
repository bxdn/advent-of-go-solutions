package day4

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        4,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	total := 0
	grid := utils.GridFromString(input)
	for {
		removed := findAccessible(grid, true)
		if removed == 0 {
			return strconv.Itoa(total), nil
		}
		total += removed
	}
}
