package day10

import (
	"advent-of-go/utils"
	"math"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 10,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	g := utils.GridFromString(input)
	asteroids := g.FindAll('#')
	maxSeen := 0
	for _, asteroid := range asteroids {
		maxSeen = max(maxSeen, len(calcSeen(asteroid, asteroids)))
	}
	return strconv.Itoa(maxSeen), nil
}

type slopePart struct {
	slope float64
	right bool
}

func calcSeen(asteroid utils.Point, asteroids []utils.Point) map[slopePart]utils.Point {
	slopeSet := map[slopePart]utils.Point{}
	for _, other := range asteroids {
		dy := other.Y - asteroid.Y
		dx := other.X - asteroid.X
		if dx == 0 && dy == 0 {
			continue
		}
		slope := float64(dy) / float64(dx)
		sp := slopePart{slope, dx > 0 || math.IsInf(slope, -1) }
		if existing, found := slopeSet[sp]; found {
			existingDy := existing.Y - asteroid.Y
			existingDx := existing.X - asteroid.X
			if abs(dy) + abs(dx) < abs(existingDx) + abs(existingDy) {
				slopeSet[sp] = other
			}
		} else {
			slopeSet[sp] = other
		}
	}
	return slopeSet
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}