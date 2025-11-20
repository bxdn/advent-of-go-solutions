package day5

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2015,
		Day:        5,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	total := 0
	for _, line := range utils.GetLines(input) {
		if improvedIsNice(line) {
			total++
		}
	}
	return strconv.Itoa(total), nil
}

func improvedIsNice(line string) bool {
	firstEncounteredPair := map[[2]rune]int{}
	doubledPair, distancedRepeat := false, false
	prevPrev, prev := '\000', '\000'
	for i, char := range line {
		if prevPrev == char {
			distancedRepeat = true
		}
		pair := [2]rune{prev, char}
		if previdx, ok := firstEncounteredPair[pair]; ok && previdx < i-1 {
			doubledPair = true
		} else if !ok {
			firstEncounteredPair[pair] = i
		}
		prevPrev, prev = prev, char
	}
	return doubledPair && distancedRepeat
}
