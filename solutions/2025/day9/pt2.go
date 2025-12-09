package day9

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        9,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	points := make([]utils.Point, len(lines))
	for i, line := range lines {
		coords, e := utils.StringsToInts(strings.Split(line, ","))
		if e != nil {
			return "", fmt.Errorf("error parsing line: %w", e)
		}
		points[i] = utils.Point{X: coords[0], Y: coords[1]}
	}
	return strconv.Itoa(calcTotal(points)), nil
}

func calcTotal(points []utils.Point) int {
	maxArea := 0
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			p1, p2 := normalize(p1, p2)
			if isValid(p1, p2, points) {
				maxArea = max(maxArea, (utils.Abs(p1.X-p2.X)+1)*(utils.Abs(p1.Y-p2.Y)+1))
			}
		}
	}
	return maxArea
}

func isValid(p1, p2 utils.Point, points []utils.Point) bool {
	points = append(points, points[0])
	for i, pb := range points[1:] {
		pa := points[i]
		pa, pb = normalize(pa, pb)
		if pa.Y == pb.Y {
			if p1.Y < pa.Y && pa.Y < p2.Y && p1.X < pb.X && pa.X < p2.X {
				return false
			}
		} else {
			if p1.X < pa.X && pa.X < p2.X && p1.Y < pb.Y && pa.Y < p2.Y {
				return false
			}
		}
	}
	return true
}

func normalize(p1, p2 utils.Point) (utils.Point, utils.Point) {
	return utils.Point{X: min(p1.X, p2.X), Y: min(p1.Y, p2.Y)}, utils.Point{X: max(p1.X, p2.X), Y: max(p1.Y, p2.Y)}
}
