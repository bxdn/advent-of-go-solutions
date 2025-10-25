package day5

import (
	"advent-of-go/solutions/2019/intcode"
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
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
	ops, e := utils.StringsToInts(strings.Split(input, ","))
	if e != nil {
		return "", fmt.Errorf("Error parsing input: %w", e)
	}
	lastOutputVal := -1
	diagInput := func() int {return 1}
	diagOutput := func(toOutput int) {lastOutputVal = toOutput}
	intcode.Run(ops, diagInput, diagOutput)
	return strconv.Itoa(lastOutputVal), nil
}