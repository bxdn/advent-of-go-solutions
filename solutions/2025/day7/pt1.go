package day7

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        7,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	row := make([]int, len(lines[0]))
	row[strings.Index(lines[0], "S")] = 1
	total := 0
	for i := 2; i < len(lines); i += 2 {
		total += processLine(lines[i], row)
	}
	return strconv.Itoa(total), nil
}

func processLine(line string, row []int) int {
	total := 0
	for i, char := range line {
		if char != '^' {
			continue
		}
		row[i-1] = 1
		row[i+1] = 1
		total += row[i]
		row[i] = 0
	}
	return total
}
