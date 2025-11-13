package day2

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2024,
		Day:        2,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	total := 0
	for _, line := range lines {
		nums, e := utils.StringsToInts(strings.Split(line, " "))
		if e != nil {
			return "", fmt.Errorf("Error: Malformed line %s, %w", line, e)
		}
		if isLineValid(nums, -1) {
			total++
		}
	}
	return strconv.Itoa(total), nil
}

func isLineValid(line []int, skip int) bool {
	increasing, decreasing := true, true
	prev := -1
	for i, num := range line {
		if skip != i {
			increasing, decreasing = processState(increasing, decreasing, prev, num)
			if !increasing && !decreasing {
				return false
			}
			prev = num
		}
	}
	return true
}

func processState(increasing, decreasing bool, prev, current int) (bool, bool) {
	if prev < 0 {
		return increasing, decreasing
	}
	dif := current - prev
	if dif >= 1 && dif <= 3 {
		return increasing, false
	} else if dif >= -3 && dif <= -1 {
		return false, decreasing
	} else {
		return false, false
	}
}
