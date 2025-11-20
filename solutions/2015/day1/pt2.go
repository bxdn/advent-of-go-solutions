package day1

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2015,
		Day:        1,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	total := 0
	for i, char := range input {
		if char == '(' {
			total++
		} else {
			total--
		}
		if total == -1 {
			return strconv.Itoa(i + 1), nil
		}
	}
	return "", fmt.Errorf("basement never reached")
}
