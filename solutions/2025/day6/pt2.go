package day6

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
	"unicode"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        6,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	operands := []int{}
	total := 0
	for i := len(lines[0]) - 1; i >= 0; i-- {
		var op rune
		operands, op = collectColumn(lines, operands, i)
		if op == '*' || op == '+' {
			total += processProblem(op, operands)
			operands = []int{}
		}
	}
	return strconv.Itoa(total), nil
}

func processProblem(operator rune, operands []int) int {
	switch operator {
	case '*':
		return getProduct(operands)
	case '+':
		return getSum(operands)
	}
	panic("Op is illegal")
}

func collectColumn(lines []string, cols []int, i int) ([]int, rune) {
	op := rune(lines[len(lines)-1][i])
	sb := strings.Builder{}
	for _, row := range lines[:len(lines)-1] {
		char := rune(row[i])
		if unicode.IsDigit(char) {
			sb.WriteRune(char)
		}
	}
	str := sb.String()
	if str == "" {
		return cols, rune(op)
	}
	num, e := strconv.Atoi(str)
	if e != nil {
		panic(e) // all digits, should be impossible
	}
	return append(cols, num), rune(op)
}

func getSum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getProduct(nums []int) int {
	prod := 1
	for _, num := range nums {
		prod *= num
	}
	return prod
}
