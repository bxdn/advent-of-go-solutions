package solutions

import (
	y2023 "advent-of-go/solutions/2023"
	"advent-of-go/utils"
)

type solutionSet struct {
	Year int
	Solutions []utils.Solution
}

func Solutions() []solutionSet { 
	return []solutionSet{
		solutionSet{2023, y2023.Solutions()},
	}
}