package solutions

import (
	y2019 "advent-of-go/solutions/2019"
	y2023 "advent-of-go/solutions/2023"
	y2024 "advent-of-go/solutions/2024"
	"advent-of-go/utils"
	"slices"
)

func Solutions() []utils.Solution {
	return slices.Concat[[]utils.Solution](y2019.Solutions(), y2023.Solutions(), y2024.Solutions())
}
