package day5

import (
	"advent-of-go/solutions/2019/intcode"
	"advent-of-go/utils"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 5,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lastOutputVal := -1
	diagInput := func() int {return 1}
	diagOutput := func(toOutput int) {lastOutputVal = toOutput}
	intcode.RunString(input, diagInput, diagOutput)
	return strconv.Itoa(lastOutputVal), nil
}