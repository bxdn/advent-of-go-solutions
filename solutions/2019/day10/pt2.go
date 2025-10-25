package day10

import (
	"advent-of-go/utils"
	"math"
	"slices"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 10,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	g := utils.GridFromString(input)
	asteroids := g.FindAll('#')
	seen := map[slopePart]utils.Point{}
	for _, asteroid := range asteroids {
		newSeen := calcSeen(asteroid, asteroids)
		if len(newSeen) > len(seen) {
			seen = newSeen
		}
	}
	keys := make([]slopePart, 0, len(seen))
	for sp := range seen {
		keys = append(keys, sp)
	}
	slices.SortFunc(keys, sortSlopeParts)
	target := seen[keys[199]]
	return strconv.Itoa(target.X * 100 + target.Y), nil
}

func sortSlopeParts(p1, p2 slopePart) int {
	if p1.right && !p2.right {
		return -1
	}
	if p2.right && !p1.right {
		return 1
	}
	if math.IsInf(p1.slope, 0) {
		return -1
	}
	if math.IsInf(p2.slope, 0) {
		return 1
	}
	if p1.slope > p2.slope {
		return 1
	}
	return -1
}