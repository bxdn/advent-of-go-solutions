package day1

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2015,
		Day:        1,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	total := 0
	for _, char := range input {
		if char == '(' {
			total++
		} else {
			total--
		}
	}
	return strconv.Itoa(total), nil
}
