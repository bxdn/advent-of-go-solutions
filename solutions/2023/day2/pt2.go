package day2

import (
	"advent-of-go/utils"
	"regexp"
	"strconv"
)

var regexs = []*regexp.Regexp{
	regexp.MustCompile(`(\d+) red`),
	regexp.MustCompile(`(\d+) green`),
	regexp.MustCompile(`(\d+) blue`),
}

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2023, 
		Day: 2,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	total := 0
	for _, line := range lines {
		power, e := getLinePower(line)
		if e != nil {
			return "", e
		}
		total += power
	}
	return strconv.Itoa(total), nil
}

func getLinePower(line string) (int, error) {
	prod := 1
	for _, regex := range regexs {
		max, e := maxOfMatchess(regex.FindAllStringSubmatch(line, -1))
		if e != nil {
			return 0, e
		}
		prod *= max
	}
	return prod, nil
}
func maxOfMatchess(matches [][]string) (int, error) {
	max := 0
	if matches == nil {
		return max, nil
	}
	for _, match := range matches {
		numStr := match[1]
		num, e := strconv.Atoi(numStr)
		if e != nil {
			return max, e
		}
		if num > max {
			max = num
		}
	}
	return max, nil
}