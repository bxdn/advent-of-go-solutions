package solutions

import (
	y2019 "advent-of-go/solutions/2019"
	y2023 "advent-of-go/solutions/2023"
	"advent-of-go/utils"
	"slices"
)

func Solutions() []utils.Solution { 
	return slices.Concat(y2019.Solutions(), y2023.Solutions())
}