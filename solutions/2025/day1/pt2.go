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
		prev := pos
		pos += offset
		if pos <= 0 {
			total += -pos / 100
			if prev != 0 {
				total++
			}
		} else {
			total += pos / 100
		}
		pos = ((pos % 100) + 100) % 100
	}
	return strconv.Itoa(total), nil
}
