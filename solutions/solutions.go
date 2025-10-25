package solutions

import (
	y2019 "advent-of-go/solutions/2019"
	"advent-of-go/utils"
)

type solutionSet struct {
	Year int
	Solutions []utils.Solution
}

func Solutions() []solutionSet { 
	return []solutionSet{
		solutionSet{2019, y2019.Solutions()},
		// solutionSet{2023, y2023.Solutions()},
	}
}