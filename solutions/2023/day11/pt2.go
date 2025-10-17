package day11

import (
	"advent-of-go/utils"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 11,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	return getAnswer(input, 1_000_000), nil
}