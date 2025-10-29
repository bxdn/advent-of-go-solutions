package main

import (
	"advent-of-go/utils"
	"fmt"
)

type testResult struct {
	solution utils.Solution
	err error
}

func testSolutions(solutions []utils.Solution) []testResult {
	answerFileContents := map[int][]string{}
	toRet := make([]testResult, len(solutions))
	for i, s := range solutions {
		e := testSolution(s, answerFileContents)
		toRet[i] = testResult{s, e}
	}
	return toRet
}

func testSolution(solution utils.Solution, answerFileContents map[int][]string) error {
	answers, found := answerFileContents[solution.Year]
	if !found {
		contents, e := utils.GetFileLines(fmt.Sprintf("private/answers/%d.txt", solution.Year))
		if e != nil {
			return fmt.Errorf("Error opening answers file for year %d: %v\n", solution.Year, e)
		}
		answerFileContents[solution.Year] = contents
		answers = contents
	}
	answerIdx := (solution.Day - 1) * 2 + solution.Part - 1
	if answerIdx >= len(answers) {
		return fmt.Errorf("Error: no answer exists for %s", solution.Name())
	}
	if e := testSolutionAgainstExpected(solution, answers[answerIdx]); e != nil {
		return fmt.Errorf("Error testing solutions for %d: %w", solution.Year, e)
	}
	return nil
}

func testSolutionAgainstExpected(sol utils.Solution, expected string) error {
	result, e := sol.Calculate()
	if e != nil {
		return fmt.Errorf("Error in %s: %v", sol.Name(), e)
	}
	if result != expected {
		return fmt.Errorf("Error in %s: expected %s but got %s", sol.Name(), expected, result)
	}
	return nil
}