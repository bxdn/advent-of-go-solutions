package day1

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        1,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	pos, total := 50, 0
	for _, line := range lines {
		offset, e := getLineOffset(line)
		if e != nil {
			return "", fmt.Errorf("error parsing line: %w", e)
		}
		pos = (offset + pos) % 100
		if pos < 0 {
			pos += 100
		} else if pos == 0 {
			total += 1
		}
	}
	return strconv.Itoa(total), nil
}

func getLineOffset(line string) (int, error) {
	if len(line) == 0 {
		return 0, fmt.Errorf("error: empty line encountered")
	}
	switch line[0] {
	case 'R':
		return strconv.Atoi(line[1:])
	case 'L':
		if n, e := strconv.Atoi(line[1:]); e != nil {
			return 0, fmt.Errorf("error parsing number: %w", e)
		} else {
			return -n, nil
		}
	default:
		return 0, fmt.Errorf("error: malformed line: unknown direction found in line: %s", line)
	}
}
