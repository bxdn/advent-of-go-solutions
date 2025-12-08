package day8

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        8,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	points, e := getPoints(input)
	if e != nil {
		return "", fmt.Errorf("error getting points: %w", e)
	}
	pairs := getDistPairs(points)
	circuits := utils.NewDisjointSet[point3d]()
	for _, point := range points {
		circuits.Add(point)
	}
	var lastPair distPair
	for _, pair := range pairs {
		rep1, rep2 := circuits.Find(pair.a), circuits.Find(pair.b)
		if rep1 != rep2 {
			circuits.Union(rep1, rep2)
		}
		if circuits.Count() == 1 {
			lastPair = pair
			break
		}
	}
	total := lastPair.a.x * lastPair.b.x
	return strconv.Itoa(total), nil
}
