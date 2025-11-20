package day2

import (
	"advent-of-go/utils"
	"fmt"
	"sort"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2015,
		Day:        2,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	total := 0
	for _, line := range lines {
		rib, e := processRibbon(line)
		if e != nil {
			return "", fmt.Errorf("error processing line: %w", e)
		}
		total += rib
	}
	return strconv.Itoa(total), nil
}

func processRibbon(line string) (int, error) {
	w, l, h, e := getDims(line)
	if e != nil {
		return 0, fmt.Errorf("error parsing line: %w", e)
	}
	dims := []int{w, l, h}
	sort.Ints(dims)
	return w*l*h + 2*(dims[0]+dims[1]), nil
}
