package main

import (
	"advent-of-go/utils"
	"encoding/json"
	"fmt"
	"os"
)

type testResult struct {
	solution utils.Solution
	err      error
}

func testSolutions(solutions []utils.Solution) []testResult {
	answerContents := utils.Unpack(os.ReadFile("private/answers.json"))
	answers := map[string]map[string][]string{}
	utils.Must(json.Unmarshal(answerContents, &answers))
	toRet := make([]testResult, len(solutions))
	for i, s := range solutions {
		e := testSolution(s, answers)
		toRet[i] = testResult{s, e}
	}
	return toRet
}

func testSolution(solution utils.Solution, allAnswers map[string]map[string][]string) error {
	yearAnswers, found := allAnswers[fmt.Sprintf("%d", solution.Year)]
	if !found {
		return fmt.Errorf("error: no answers found for year %d", solution.Year)
	}
	dayAnswers, found := yearAnswers[fmt.Sprintf("%d", solution.Day)]
	if !found || len(dayAnswers) < solution.Part {
		return fmt.Errorf("error: no answers found for year %d day %d part %d", solution.Year, solution.Day, solution.Part)
	}
	result, e := solution.Calculate()
	if e != nil {
		return fmt.Errorf("error in %s: %v", solution.Name(), e)
	}
	expected := dayAnswers[solution.Part-1]
	if result != expected {
		return fmt.Errorf("error in %s: expected %s but got %s", solution.Name(), expected, result)
	}
	return nil
}
