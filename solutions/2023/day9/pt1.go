package day9

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 9,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	total := 0
	for _, line := range lines {
		total += extrapolateLine(line, false)
	}
	return strconv.Itoa(total), nil
}

func extrapolateLine(line string, rev bool) int {
	nums := utils.FindInts(line)
	if rev {
		nums = utils.Rev(nums)
	}
	numStack := [][]int{nums}
	numStack = propagateForward(numStack)
	return propagateBackward(numStack)
}

func propagateForward(numStack [][]int) [][]int {
	var done bool
	for {
		numStack, done = forwardOnce(numStack)
		if done {
			return numStack
		}
	}
}

func propagateBackward(numStack [][]int) int {
	extrapolated := 0
	for i := len(numStack) - 1; i > 0; i-- {
		extrapolated += utils.Last(numStack[i - 1]) 
	}
	return extrapolated
}

func forwardOnce(numStack [][]int) ([][]int, bool) {
	currentNums := utils.Last(numStack)
	done := true
	nextNums := make([]int, len(currentNums) - 1)
	for i := range nextNums {
		toSet := currentNums[i + 1] - currentNums[i]
		if toSet != 0 {
			done = false
			nextNums[i] = toSet
		}
	}
	numStack = append(numStack, nextNums)
	return numStack, done
}