package day2

import (
	"advent-of-go/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2015,
		Day:        2,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	total := 0
	for _, line := range lines {
		wrap, e := processWrappingPaper(line)
		if e != nil {
			return "", fmt.Errorf("error processing line: %w", e)
		}
		total += wrap
	}
	return strconv.Itoa(total), nil
}

func processWrappingPaper(line string) (int, error) {
	w, l, h, e := getDims(line)
	if e != nil {
		return 0, fmt.Errorf("error parsing line: %w", e)
	}
	sides := []int{w * l, l * h, h * w}
	sort.Ints(sides)
	total := sides[0]
	for _, side := range sides {
		total += 2 * side
	}
	return total, nil
}

func getDims(line string) (int, int, int, error) {
	toks := strings.Split(line, "x")
	if len(toks) != 3 {
		return 0, 0, 0, fmt.Errorf("error: should be 3 x-split numbers")
	}
	w, e := strconv.Atoi(toks[0])
	if e != nil {
		return 0, 0, 0, fmt.Errorf("error: first number %s was not a number: %w", toks[0], e)
	}
	l, e := strconv.Atoi(toks[1])
	if e != nil {
		return 0, 0, 0, fmt.Errorf("error: second number %s was not a number: %w", toks[1], e)
	}
	h, e := strconv.Atoi(toks[2])
	if e != nil {
		return 0, 0, 0, fmt.Errorf("error: third number %s was not a number: %w", toks[2], e)
	}
	return w, l, h, nil
}
