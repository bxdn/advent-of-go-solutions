package day9

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 9,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	total := 0
	for _, line := range lines {
		total += extrapolateLine(line, true)
	}
	return strconv.Itoa(total), nil
}
