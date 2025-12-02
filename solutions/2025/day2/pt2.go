package day2

import (
	"advent-of-go/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        2,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	total := 0
	for rang := range strings.SplitSeq(input, ",") {
		if invalid, e := processRange(rang, pt2Val); e != nil {
			return "", fmt.Errorf("error processing range: %w", e)
		} else {
			total += invalid
		}
	}
	return strconv.Itoa(total), nil
}

func pt2Val(n int) bool {
	digits := utils.Digits(n)
	doubled := slices.Concat(digits, digits)
	return utils.ContainsSubslice(doubled[1:len(doubled)-1], digits)
}
