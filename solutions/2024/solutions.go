package y2024

import (
	"advent-of-go/solutions/2024/day1"
	"advent-of-go/solutions/2024/day2"
	"advent-of-go/solutions/2024/day3"
	"advent-of-go/utils"
)

func Solutions() []utils.Solution {
	return []utils.Solution{
		day1.Pt1(), day1.Pt2(), day2.Pt1(), day2.Pt2(),
		day3.Pt1(), day3.Pt2(),
	}
}
