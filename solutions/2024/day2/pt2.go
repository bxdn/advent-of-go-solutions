package day2

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2024,
		Day:        2,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	total := 0
	for _, line := range lines {
		nums, e := utils.StringsToInts(strings.Split(line, " "))
		if e != nil {
			return "", fmt.Errorf("error parsing malformed line %s, %w", line, e)
		}
		if isLineValidWithFaltTolerance(nums) {
			total++
		}
	}
	return strconv.Itoa(total), nil
}

func isLineValidWithFaltTolerance(nums []int) bool {
	for i := -1; i < len(nums); i++ {
		if isLineValid(nums, i) {
			return true
		}
	}
	return false
}
