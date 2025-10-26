package day7

import (
	"advent-of-go/solutions/2019/intcode"
	"advent-of-go/utils"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 7,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	nums := []int{0, 1, 2, 3, 4}
	maxSignal := 0
	var e error
	permutations(nums, func(p []int) {
		maxSignal = max(maxSignal, runPermutation(input, p))
	})
	return strconv.Itoa(maxSignal), e
}

func runPermutation(program string, phases []int) int {
	outputSignal := 0
	for _, phase := range phases {
		sendPhase := true
		input := func() int {
			if sendPhase {
				sendPhase = false
				return phase
			}
			return outputSignal
		}
		output := func(n int) {outputSignal = n}
		intcode.RunString(program, input, output)
	}
	return outputSignal
}

// permutations generates all permutations of nums and calls fn for each one.
func permutations(nums []int, fn func([]int)) {
	var generate func(int)
	generate = func(n int) {
		if n == 1 {
			cp := make([]int, len(nums))
			copy(cp, nums)
			fn(cp)
			return
		}
		for i := 0; i < n; i++ {
			generate(n - 1)
			if n%2 == 1 {
				nums[0], nums[n-1] = nums[n-1], nums[0]
			} else {
				nums[i], nums[n-1] = nums[n-1], nums[i]
			}
		}
	}
	generate(len(nums))
}