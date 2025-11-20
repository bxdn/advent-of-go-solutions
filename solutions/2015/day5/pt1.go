package day5

import (
	"advent-of-go/utils"
	"slices"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2015,
		Day:        5,
		Part:       1,
		Calculator: pt1,
	}
}

var vowels = []rune{'a', 'e', 'i', 'o', 'u'}

func pt1(input string) (string, error) {
	total := 0
	for _, line := range utils.GetLines(input) {
		if isNice(line) {
			total++
		}
	}
	return strconv.Itoa(total), nil
}

func isNice(line string) bool {
	doubled := false
	vowelCount := 0
	prev := '\000'
	for _, char := range line {
		if slices.Contains(vowels, char) {
			vowelCount++
		}
		if prev == char {
			doubled = true
		}
		if prev == 'a' && char == 'b' || prev == 'c' && char == 'd' ||
			prev == 'p' && char == 'q' || prev == 'x' && char == 'y' {
			return false
		}
		prev = char
	}
	return doubled && vowelCount >= 3
}
