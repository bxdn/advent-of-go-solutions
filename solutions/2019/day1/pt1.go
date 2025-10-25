package day1

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 1,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	total := 0
	for _, line := range lines {
		num, e := getLineNum(line)
		if e != nil {
			return "", e
		}
		total += num
	}
	return strconv.Itoa(total), nil
}

func getLineNum(line string) (int, error) {
	total, e := strconv.Atoi(line)
	if e != nil {
		return 0, fmt.Errorf("Malformed input: line %s was not able to be parsed: %w", line, e)
	}
	return total / 3 - 2, nil
}