package y2023

import (
	"advent-of-go/solutions/2023/day1"
	"advent-of-go/solutions/2023/day2"
	"advent-of-go/solutions/2023/day3"
	"advent-of-go/solutions/2023/day4"
	"advent-of-go/solutions/2023/day5"
	"advent-of-go/utils"
)

func Solutions() []utils.Solution {
	return []utils.Solution {
		day1.Pt1(),
		day1.Pt2(),
		day2.Pt1(),
		day2.Pt2(),
		day3.Pt1(),
		day3.Pt2(),
		day4.Pt1(),
		day4.Pt2(),
		day5.Pt1(),
		day5.Pt2(),
	}
}