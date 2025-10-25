package day3

import (
	"advent-of-go/utils"
	"fmt"
	"math"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 3,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	totalHRanges, totalVRanges := []rang{}, []rang{}
	if len(lines) != 2 {
		return "", fmt.Errorf("Malformed input! Expected 2 lines but got %d", len(lines))
	}
	walker1 := newWireWalker()
	if e := walker1.calcRanges(lines[0]); e != nil {
		return "", fmt.Errorf("Error parsing line: %w", e)
	}
	totalHRanges = append(totalHRanges, walker1.hRanges...)
	totalVRanges = append(totalVRanges, walker1.vRanges...)
	walker2 := newWireWalker()
	if e := walker2.calcRanges(lines[1]); e != nil {
		return "", fmt.Errorf("Error parsing line: %w", e)
	}
	totalHRanges = append(totalHRanges, walker2.hRanges...)
	totalVRanges = append(totalVRanges, walker2.vRanges...)
	intersections := findIntersections(totalHRanges, totalVRanges)
	minDistance := math.MaxInt
	for _, intersection := range intersections {
		wire1Distance, wire2Distance := getDistance(walker1, intersection), getDistance(walker2, intersection)
		if wire1Distance > 0 && wire2Distance > 0 {
			totalDistance := wire1Distance + wire2Distance
			if totalDistance < minDistance {
				minDistance = totalDistance
			}
		}
	}
	return strconv.Itoa(minDistance), nil
}

func getDistance(walker wireWalker, intersection utils.Point) int {
	curX, curY := 0, 0
	totalTravelled := 0
	for _, rng := range walker.allRanges {
		if rng.isVertical && 
				intersection.X == rng.start.X && 
				intersection.Y > rng.start.Y && 
				intersection.Y < rng.start.Y + rng.length {
			return totalTravelled + utils.Abs(curY - intersection.Y)
		}
		if !rng.isVertical && intersection.Y == rng.start.Y && intersection.X > rng.start.X && intersection.X < rng.start.X + rng.length {
			return totalTravelled + utils.Abs(curX - intersection.X)
		}
		if rng.isVertical {
			if rng.start.Y == curY {
				curY += rng.length
			} else {
				curY = rng.start.Y
			}
		} else {
			if rng.start.X == curX {
				curX += rng.length
			} else {
				curX = rng.start.X
			}
		}
		totalTravelled += rng.length
	}
	return -1
}