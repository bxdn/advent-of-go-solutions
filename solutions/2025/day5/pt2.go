package day5

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
		Day:        5,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	toks := strings.Split(input, "\n\n")
	if len(toks) != 2 {
		return "", fmt.Errorf("expected 2 sections in input, got %d", len(toks))
	}
	ranges, e := getRanges(toks[0])
	if e != nil {
		return "", fmt.Errorf("error getting ranges: %w", e)
	}
	return strconv.Itoa(sumRanges(ranges)), nil
}

func sumRanges(ranges []Range) int {
	slices.SortFunc(ranges, func(r1, r2 Range) int {
		return r1.start - r2.start
	})
	total, curEnd := 0, 0
	for _, rng := range ranges {
		start := max(rng.start, curEnd+1)
		total += max(0, rng.end-start+1)
		curEnd = max(rng.end, curEnd)
	}
	return total
}
