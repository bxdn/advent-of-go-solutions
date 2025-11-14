package day1

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2024,
		Day:        1,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	numStrs := strings.Fields(input)
	left, right := map[int]int{}, map[int]int{}
	for i, str := range numStrs {
		num, e := strconv.Atoi(str)
		if e != nil {
			return "", fmt.Errorf("error Parsing number: %w", e)
		}
		if i&1 == 0 {
			left[num] += 1
		} else {
			right[num] += 1
		}
	}
	total := 0
	for k, v := range left {
		total += k * v * right[k]
	}
	return strconv.Itoa(total), nil
}
