package day1

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        1,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	pos, total := 50, 0
	for _, line := range lines {
		offset, e := getLineOffset(line)
		if e != nil {
			return "", fmt.Errorf("error parsing line: %w", e)
		}
		if pos+offset <= 0 && pos != 0 {
			total++
		}
		pos += offset
		total += utils.Abs(pos) / 100
		pos = ((pos % 100) + 100) % 100
	}
	return strconv.Itoa(total), nil
}
