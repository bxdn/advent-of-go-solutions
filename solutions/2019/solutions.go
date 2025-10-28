package y2023

import (
	"advent-of-go/solutions/2019/day1"
	"advent-of-go/solutions/2019/day10"
	"advent-of-go/solutions/2019/day11"
	"advent-of-go/solutions/2019/day12"
	"advent-of-go/solutions/2019/day13"
	"advent-of-go/solutions/2019/day14"
	"advent-of-go/solutions/2019/day15"
	"advent-of-go/solutions/2019/day2"
	"advent-of-go/solutions/2019/day3"
	"advent-of-go/solutions/2019/day4"
	"advent-of-go/solutions/2019/day5"
	"advent-of-go/solutions/2019/day6"
	"advent-of-go/solutions/2019/day7"
	"advent-of-go/solutions/2019/day8"
	"advent-of-go/solutions/2019/day9"
	"advent-of-go/utils"
)

func Solutions() []utils.Solution {
	return []utils.Solution {
		day1.Pt1(), day1.Pt2(), day2.Pt1(), day2.Pt2(),
		day3.Pt1(), day3.Pt2(), day4.Pt1(), day4.Pt2(),
		day5.Pt1(), day5.Pt2(), day6.Pt1(), day6.Pt2(),
		day7.Pt1(), day7.Pt2(), day8.Pt1(), day8.Pt2(),
		day9.Pt1(), day9.Pt2(), day10.Pt1(), day10.Pt2(),
		day11.Pt1(), day11.Pt2(), day12.Pt1(), day12.Pt2(),
		day13.Pt1(), day13.Pt2(), day14.Pt1(), day14.Pt2(),
		day15.Pt1(), day15.Pt2(),
	}
}