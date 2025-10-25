package day7

import (
	"advent-of-go/solutions/2019/intcode"
	"advent-of-go/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 7,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	nums := []int{5, 6, 7, 8, 9}
	baseProgram, e := utils.StringsToInts(strings.Split(input, ","))
	if e != nil {
		return "", fmt.Errorf("Error parsing program: %w", e)
	}
	maxSignal := 0
	run := func(phases []int) {
		maxSignal = max(maxSignal, runPermutationParallel(baseProgram, phases))
	}
	permutations(nums, run)
	return strconv.Itoa(maxSignal), nil
}

func runPermutationParallel(baseProgram []int, phases []int) int {
	chans := make([]chan int, 5)
	tg := utils.TaskGroup{}
	for i := range 5 {
		chans[i] = make(chan int, 1)
		fn := func() {
			runAmp(slices.Clone(baseProgram), chans[i], chans[(i + 1) % 5], phases[i])
		}
		tg.Add(fn)
	}
	tg.Start()
	startChan := chans[0]
	startChan <- 0
	tg.Wait()
	return <- startChan
}

func runAmp(program []int, inChan, outChan chan int, phase int) {
	sendPhase := true
	input := func() int {
		if sendPhase {
			sendPhase = false
			return phase
		}
		return <- inChan
	}
	output := func(n int) {
		outChan <- n
	}
	intcode.Run(program, input, output)
}