package main

import (
	"advent-of-go/solutions"
	"advent-of-go/utils"
	"fmt"
	"testing"
)

func TestSolutions(t *testing.T) {
	for _, solutionSet := range solutions.Solutions() {
		testSolutions(t, solutionSet.Year, solutionSet.Solutions)
	}
}

func testSolutions(t *testing.T, year int, solutions []utils.Solution) {
	answers, e := utils.GetFileLines(fmt.Sprintf("private/answers/%d.txt", year))
	if e != nil {
		t.Errorf("Error opening answers file for year %d: %v", year, e)
		return
	}
	if len(answers) != len(solutions) {
		t.Errorf("Length of solutions and answers differs for year %d: solutions has length %d, answers has length %d", year, len(solutions), len(answers))
		return
	}
	for i, s := range solutions {
		testSolution(t, s, answers[i])
	}
}

func testSolution(t *testing.T, sol utils.Solution, expected string) {
	result, e := sol.Calculate()
	if e != nil {
		t.Errorf("Error in %s: %v", sol.Name(), e)
		return
	}
	if result != expected {
		t.Errorf("Error in %s: expected %s but got %s", sol.Name(), expected, result)
	}
}