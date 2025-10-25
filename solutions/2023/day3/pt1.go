package day3

import (
	"advent-of-go/utils"
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

var regex = regexp.MustCompile(`\d+`)

type numWithPos struct {
	numString string
	x int
	y int
}

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2023, 
		Day: 3,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	grid := utils.GridFromLines(lines)
	nums := []numWithPos{}
	for i, line := range lines {
		nums = append(nums, getLineNums(i, line)...)
	}
	total := 0
	for _, strNum := range nums {
		num, e := getNumTotal(strNum, grid)
		if e != nil {
			return "", fmt.Errorf("Error getting the total for the string number: %w", e)
		}
		total += num
	}
	return strconv.Itoa(total), nil
}

func getLineNums(y int, line string) []numWithPos {
	toRet := []numWithPos{}
	for _, match := range regex.FindAllStringIndex(line, -1) {
		start, end := match[0], match[1]
		num := line[start:end]
		toRet = append(toRet, numWithPos{num, start, y})
	}
	return toRet
}

func getNumTotal(strNum numWithPos, grid utils.FinGrid[rune]) (int, error) {
	border := getBorder(strNum, grid)
	for _, char := range border {
		if !unicode.IsDigit(char) && '.' != char {
			return strconv.Atoi(strNum.numString)
		}
	}
	return 0, nil
}

func getBorder(strNum numWithPos, grid utils.FinGrid[rune]) []rune {
	length := len(strNum.numString)
	toRet := make([]rune, 0, length * 2 + 6)
	x, y := strNum.x, strNum.y
	for deltaX := -1; deltaX <= length; deltaX++ {
		toRet = append(toRet,
			grid.At(x + deltaX, y + 1).Or('.'),
		 	grid.At(x + deltaX, y - 1).Or('.'),
		) 
	}
	toRet = append(toRet, 
		grid.At(x - 1, y).Or('.'), 
		grid.At(x + length, y).Or('.'),
	) 
	return toRet
}