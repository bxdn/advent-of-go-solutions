package day9

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
		Day: 9,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	ops, e := utils.StringsToInts(strings.Split(input, ","))
	if e != nil {
		return "", fmt.Errorf("Error parsing input: %w", e)
	}
	outs := []int{}
	out := func (i int) {
		outs = append(outs, i)
	}
	in := func() int {
		return 1
	}
	if e = intcode.Run(ops, in, out); e != nil {
		return "", fmt.Errorf("Error while running program: %w", e)
	}
	return strconv.Itoa(outs[0]), nil
}