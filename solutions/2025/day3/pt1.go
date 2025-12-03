package day3

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        3,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	total := 0
	lines := utils.GetLines(input)
	for _, line := range lines {
		total += solve(line, 2)
	}
	return strconv.Itoa(total), nil
}

func solve(line string, nDigits int) int {
	length := len(line)
	total, startIdx := 0, 0
	for offset := nDigits - 1; offset >= 0; offset-- {
		maxDigit, maxOffset := maxWithIndex(line[startIdx : length-offset])
		startIdx += maxOffset + 1
		total += int(maxDigit-'0') * utils.Pow(10, offset)
	}
	return total
}

func maxWithIndex(nums string) (rune, int) {
	maxNum, maxIdx := rune(nums[0]), 0
	for i, n := range nums[1:] {
		if n > maxNum {
			maxNum, maxIdx = n, i+1
		}
	}
	return maxNum, maxIdx
}
