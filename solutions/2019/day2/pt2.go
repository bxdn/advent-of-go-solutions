package day2

import (
	"advent-of-go/solutions/2019/intcode"
	"advent-of-go/utils"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 2,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	ops, e := utils.StringsToInts(strings.Split(input, ","))
	if e != nil {
		return "", fmt.Errorf("Error parsing input: %w", e)
	}
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			tmpOps := slices.Clone(ops)
			tmpOps[1] = i
			tmpOps[2] = j
			if e := intcode.RunBasic(tmpOps); e != nil {
				return "", fmt.Errorf("Error running code with input %d, %d: %w", i, j, e)
			}
			if tmpOps[0] == 19690720 {
				return strconv.Itoa(100 * i + j), nil
			}
		}
	}
	return "", errors.New("Did not find a valid pair to produce the desired output!")
}