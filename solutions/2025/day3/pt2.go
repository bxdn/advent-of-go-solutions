package day3

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        3,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	total := 0
	lines := utils.GetLines(input)
	for _, line := range lines {
		total += solve(line, 12)
	}
	return strconv.Itoa(total), nil
}
