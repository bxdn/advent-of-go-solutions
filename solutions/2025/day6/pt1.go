package day6

import (
	"advent-of-go/utils"
	"fmt"
	"regexp"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        6,
		Part:       1,
		Calculator: pt1,
	}
}

var regex = regexp.MustCompile(`\d+`)
var regex2 = regexp.MustCompile(`\S+`)

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	numLines, e := getNumLines(lines[:len(lines)-1])
	if e != nil {
		return "", fmt.Errorf("error getting number lines: %w", e)
	}
	ops := regex2.FindAllString(lines[len(lines)-1], -1)
	return strconv.Itoa(getTotal(numLines, ops)), nil

}

func getNumLines(lines []string) ([][]int, error) {
	numLines := make([][]int, len(lines))
	for i, line := range lines {
		nums, e := utils.StringsToInts(regex.FindAllString(line, -1))
		if e != nil {
			return nil, fmt.Errorf("error converting row to numbers: %w", e)
		}
		numLines[i] = nums
	}
	return numLines, nil
}

func getTotal(nums [][]int, ops []string) int {
	total := 0
	for i := range len(nums[0]) {
		switch ops[i] {
		case "+":
			total += getColSum(nums, i)
		case "*":
			total += getColProduct(nums, i)
		}
	}
	return total
}

func getColSum(nums [][]int, i int) int {
	sum := 0
	for _, row := range nums {
		sum += row[i]
	}
	return sum
}

func getColProduct(nums [][]int, i int) int {
	prod := 1
	for _, row := range nums {
		prod *= row[i]
	}
	return prod
}
