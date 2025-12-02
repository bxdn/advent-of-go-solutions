package day2

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        2,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	total := 0
	for rang := range strings.SplitSeq(input, ",") {
		if invalid, e := processRange(rang, pt1Val); e != nil {
			return "", fmt.Errorf("error processing range: %w", e)
		} else {
			total += invalid
		}
	}
	return strconv.Itoa(total), nil
}

func processRange(rang string, validator func(int) bool) (int, error) {
	minN, maxN := 0, 0
	_, e := fmt.Sscanf(rang, "%d-%d", &minN, &maxN)
	if e != nil {
		return 0, fmt.Errorf("error parsing range: %w", e)
	}
	invalid := 0
	for i := minN; i <= maxN; i++ {
		if validator(i) {
			invalid += i
		}
	}
	return invalid, nil
}

func pt1Val(n int) bool {
	s := strconv.Itoa(n)
	subLength := len(s) / 2
	return s[:subLength] == s[subLength:]
}
