package day3

import (
	"advent-of-go/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 3,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	hRanges, vRanges := []rang{}, []rang{}
	for _, line := range lines {
		walker := newWireWalker()
		if e := walker.calcRanges(line); e != nil {
			return "", fmt.Errorf("Error parsing line: %w", e)
		}
		hRanges = append(hRanges, walker.hRanges...)
		vRanges = append(vRanges, walker.vRanges...)
	}
	intersections := findIntersections(hRanges, vRanges)
	minDistance := math.MaxInt
	for _, intersection := range intersections {
		distance := utils.Abs(intersection.X) + utils.Abs(intersection.Y)
		if distance < minDistance && distance > 0 {
			minDistance = distance
		}
	}
	return strconv.Itoa(minDistance), nil
}

func findIntersections(hRanges, vRanges []rang) []utils.Point {
	intersections := []utils.Point{}
	for _, hRange := range hRanges {
		for _, vRange := range vRanges {
			if hRange.start.Y > vRange.start.Y && 
					hRange.start.Y < vRange.start.Y + vRange.length && 
					vRange.start.X > hRange.start.X && 
					vRange.start.X < hRange.start.X + hRange.length {
				intersections = append(intersections, utils.Point{X: vRange.start.X, Y: hRange.start.Y})
			}
		}
	}
	return intersections
}

type rang struct {
	start utils.Point
	length int
	isVertical bool
}

type wireWalker struct {
	hRanges, vRanges, allRanges []rang
	curX, curY int
}

func newWireWalker() wireWalker {
	walker := wireWalker{}
	walker.hRanges = []rang{}
	walker.vRanges = []rang{}
	return walker
}

func (w *wireWalker) calcRanges(line string) error {
	toks := strings.Split(line, ",")
	for _, r := range toks {
		if e := w.walkRange(r); e != nil {
			return fmt.Errorf("Error parsing num %s: %w", r, e)
		}
	}
	return nil
}

func (w *wireWalker) walkRange(r string) error {
	dir := rune(r[0])
	num, e := strconv.Atoi(r[1:])
	if e != nil {
		return fmt.Errorf("Error parsing num %s: %w", r, e)
	}
	switch dir {
		case 'R': {
			newRange := rang{utils.Point{X: w.curX, Y: w.curY}, num, false}
			w.hRanges = append(w.hRanges, newRange)
			w.allRanges = append(w.allRanges, newRange)
			w.curX += num
		}
		case 'L': {
			w.curX -= num
			newRange := rang{utils.Point{X: w.curX, Y: w.curY}, num, false}
			w.hRanges = append(w.hRanges, newRange)
			w.allRanges = append(w.allRanges, newRange)
		}
		case 'U': {
			newRange := rang{utils.Point{X: w.curX, Y: w.curY}, num, true}
			w.vRanges = append(w.vRanges, newRange)
			w.allRanges = append(w.allRanges, newRange)
			w.curY += num
		}
		case 'D': {
			w.curY -= num
			newRange := rang{utils.Point{X: w.curX, Y: w.curY}, num, true}
			w.vRanges = append(w.vRanges, newRange)
			w.allRanges = append(w.allRanges, newRange)
		}
		default: return fmt.Errorf("Unexpected direction %c", dir)
	}
	return nil
}