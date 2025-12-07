package day7

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        7,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	counts := make([]int, len(lines[0]))
	counts[strings.Index(lines[0], "S")] = 1
	for i := 2; i < len(lines); i += 2 {
		processLineEverett(lines[i], counts)
	}
	return strconv.Itoa(sum(counts)), nil
}

func processLineEverett(line string, counts []int) {
	for i, char := range line {
		if char != '^' {
			continue
		}
		counts[i-1] += counts[i]
		counts[i+1] += counts[i]
		counts[i] = 0
	}
}

func sum(timelines []int) int {
	total := 0
	for _, n := range timelines {
		total += n
	}
	return total
}
