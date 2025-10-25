package day2

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
		Day: 2,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	ops, e := utils.StringsToInts(strings.Split(input, ","))
	if e != nil {
		return "", fmt.Errorf("Error parsing input: %w", e)
	}
	ops[1] = 12
	ops[2] = 2
	intcode.RunBasic(ops)
	return strconv.Itoa(ops[0]), nil
}

