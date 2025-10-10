package day1

import (
	"advent-of-go/utils"
	"strconv"
	"unicode"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2023, 
		Day: 1,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	total := 0
	for _, line := range lines {
		num, e := getLineNum(line)
		if e != nil {
			return "", e
		}
		total += num
	}
	return strconv.Itoa(total), nil
}

func getLineNum(line string) (int, error) {
	var first, last rune
	for _, r := range line {
		if unicode.IsDigit(r) {
			if first == 0 {
				first = r
			}
			last = r
		}
	}
	return strconv.Atoi(string(first) + string(last))
}