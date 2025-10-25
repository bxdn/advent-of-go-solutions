package day5

import (
	"advent-of-go/solutions/2019/intcode"
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 5,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	ops, e := utils.StringsToInts(strings.Split(input, ","))
	if e != nil {
		return "", fmt.Errorf("Error parsing input: %w", e)
	}
	lastOutputVal := -1
	diagInput := func() int {return 5}
	diagOutput := func(toOutput int) {lastOutputVal = toOutput}
	intcode.Run(ops, diagInput, diagOutput)
	return strconv.Itoa(lastOutputVal), nil
}