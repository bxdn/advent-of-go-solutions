package day15

import (
	"advent-of-go/solutions/2019/intcode"
	"advent-of-go/utils"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 15,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	e := intcode.RunBasicString(input)
	return "Program run!", e
}