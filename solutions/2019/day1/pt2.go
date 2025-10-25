package day1

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 1,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	total := 0
	for _, line := range lines {
		num, e := getLineNumPt2(line)
		if e != nil {
			return "", e
		}
		total += num
	}
	return strconv.Itoa(total), nil
}

func getLineNumPt2(line string) (int, error) {
	remaining, e := strconv.Atoi(line)
	if e != nil {
		return 0, fmt.Errorf("Malformed input: line %s was not able to be parsed: %w", line, e)
	}
	needed := 0
	for remaining > 8 {
		remaining = remaining / 3 - 2
		needed += remaining
	}
	return needed, nil
}