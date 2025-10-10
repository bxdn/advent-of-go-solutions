package main

import (
	y2023 "advent-of-go/solutions/2023"
	"advent-of-go/utils"
	"fmt"
	"testing"
)

func TestSolutions(t *testing.T) {
	testSolutions(t, "2023", y2023.Solutions())
}

func testSolutions(t *testing.T, year string, solutions []utils.Solution) {
	answers, e := utils.GetFileLines(fmt.Sprintf("answers/%s/answers.txt", year))
	if e != nil {
		t.Errorf("Error opening answers file for year %s: %v", year, e)
		return
	}
	if len(answers) != len(solutions) {
		t.Errorf("Length of solutions and answers differs for year %s: solutions has length %d, answers has length %d", year, len(solutions), len(answers))
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