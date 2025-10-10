package day2

import (
	"advent-of-go/utils"
	"regexp"
	"strconv"
)

var redRegex = regexp.MustCompile(`(\d+) red`)
var greenRegex = regexp.MustCompile(`(\d+) green`)
var blueRegex = regexp.MustCompile(`(\d+) blue`)

type regexMax struct {
	regex *regexp.Regexp
	max int
}
var regexMaxs = []regexMax{
	regexMax{redRegex, 12}, 
	regexMax{greenRegex, 13}, 
	regexMax{blueRegex, 14},
}

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2023, 
		Day: 2,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	total := 0
	for i, line := range lines {
		isValid, e := isLineValid(line)
		if e != nil {
			return "", e
		}
		if isValid {
			total += i + 1
		}
	}
	return strconv.Itoa(total), nil
}

func isLineValid(line string) (bool, error) {
	for _, regexMax := range regexMaxs {
		over, e := isOverMax(regexMax.regex, line, regexMax.max)
		if e != nil {
			return false, e
		}
		if over {
			return false, nil
		}
	}
	return true, nil
}

func isOverMax(regex *regexp.Regexp, line string, maxAllowed int) (bool, error) {
	matches := regex.FindAllStringSubmatch(line, -1)
	max, e := maxOfMatches(matches)
	if e != nil {
		return false, e
	}
	if max > maxAllowed {
		return true, nil
	}

	return false, nil
}

func maxOfMatches(matches [][]string) (int, error) {
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