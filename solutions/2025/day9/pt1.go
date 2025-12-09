package day9

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        9,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	points := make([]utils.Point, len(lines))
	for i, line := range lines {
		coords, e := utils.StringsToInts(strings.Split(line, ","))
		if e != nil {
			return "", fmt.Errorf("error parsing line: %w", e)
		}
		points[i] = utils.Point{X: coords[0], Y: coords[1]}
	}
	maxArea := 0
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			maxArea = max(maxArea, (utils.Abs(p1.X-p2.X)+1)*(utils.Abs(p1.Y-p2.Y)+1))
		}
	}
	return strconv.Itoa(maxArea), nil
}
