package day16

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 16,
		Part: 1,
		Calculator: pt1,
	}
}

var pat = [4]int{0, 1, 0, -1}

func pt1(input string) (string, error) {
	current := stringToInts(input)
	for range 100 {
		current = runPhase(current)
	}
	return strconv.Itoa(concat(current[:8])), nil
}

func concat(digits []int) int {
	total := 0
	for _, d := range digits {
		total *= 10
		total += d
	}
	return total
}

func runPhase(current []int) []int {
	numDigits := len(current)
	next := make([]int, numDigits)
	for i := range numDigits {
		next[i] = processDigit(current, i)
	}
	return next
}

func processDigit(current []int, i int) int {
	total := 0
	for j, digit := range current {
		patIdx := ((1 + j) / (i + 1)) % 4
		total += digit * pat[patIdx]
	}
	if total < 0 {
		total = -total
	}
	return total % 10
}

func stringToInts(input string) []int {
	toRet := []int{}
	for _, char := range input {
		toRet = append(toRet, int(char - '0'))
	}
	return toRet
}