package day4

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 4,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	rng, e := utils.StringsToInts(strings.Split(input, "-"))
	if e != nil {
		return "", fmt.Errorf("Error parsing tokens: %w", e)
	}
	if len(rng) != 2 {
		return "", fmt.Errorf("Input %s is malformed: should have 2 - delimited sections", input)
	}
	min, max := rng[0], rng[1]
	total := 0
	for i := min; i <= max; i++ {
		if isValidPassword(i) {
			total++
		}
	}
	return strconv.Itoa(total), nil
}

func isValidPassword(n int) bool {
	prevRune := '.'
	nStr := strconv.Itoa(n)
	repeat := false
	for _, char := range nStr {
		if char == prevRune {
			repeat = true
		}
		if char < prevRune {
			return false
		}
		prevRune = char
	}
	return repeat
}