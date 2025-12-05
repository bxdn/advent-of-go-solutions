package day5

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        5,
		Part:       1,
		Calculator: pt1,
	}
}

type Range struct {
	start, end int
}

func pt1(input string) (string, error) {
	toks := strings.Split(input, "\n\n")
	if len(toks) != 2 {
		return "", fmt.Errorf("expected 2 sections in input, got %d", len(toks))
	}
	ranges, e := getRanges(toks[0])
	if e != nil {
		return "", fmt.Errorf("error getting ranges: %w", e)
	}
	nums, e := utils.StringsToInts(utils.GetLines(toks[1]))
	if e != nil {
		return "", fmt.Errorf("error getting nums: %w", e)
	}
	total := 0
	for _, n := range nums {
		if isFresh(n, ranges) {
			total++
		}
	}
	return strconv.Itoa(total), nil
}

func isFresh(n int, ranges []Range) bool {
	for _, rng := range ranges {
		if n <= rng.end && n >= rng.start {
			return true
		}
	}
	return false
}

func getRanges(rangeList string) ([]Range, error) {
	rangeLines := utils.GetLines(rangeList)
	ranges := make([]Range, len(rangeLines))
	for i, line := range rangeLines {
		rng, e := parseRange(line)
		if e != nil {
			return nil, fmt.Errorf("error parsing range: %w", e)
		}
		ranges[i] = rng
	}
	return ranges, nil
}

func parseRange(line string) (Range, error) {
	var start, end int
	_, e := fmt.Sscanf(line, "%d-%d", &start, &end)
	if e != nil {
		return Range{}, fmt.Errorf("error parsing range: %w", e)
	}
	return Range{start, end}, nil
}
