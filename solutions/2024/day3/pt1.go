package day3

import (
	"advent-of-go/utils"
	"fmt"
	"regexp"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2024,
		Day:        3,
		Part:       1,
		Calculator: pt1,
	}
}

var mulRe = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func pt1(input string) (string, error) {
	total := 0
	for _, matches := range mulRe.FindAllStringSubmatch(input, -1) {
		if len(matches) != 3 {
			return "", fmt.Errorf("Match did not have the right subgroups?")
		}
		prod, e := getProduct(matches[1], matches[2])
		if e != nil {
			return "", fmt.Errorf("Error Parsing mul() command: %w", e)
		}
		total += prod
	}
	return strconv.Itoa(total), nil
}

func getProduct(lStr, rStr string) (int, error) {
	l, e := strconv.Atoi(lStr)
	if e != nil {
		return 0, fmt.Errorf("Error Parsing left opperand: %w", e)
	}
	r, e := strconv.Atoi(rStr)
	if e != nil {
		return 0, fmt.Errorf("Error Parsing right opperand: %w", e)
	}
	return l * r, nil
}
