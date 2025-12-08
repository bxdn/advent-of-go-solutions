package day8

import (
	"advent-of-go/utils"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        8,
		Part:       1,
		Calculator: pt1,
	}
}

type point3d struct {
	x, y, z int
}

type distPair struct {
	a, b   point3d
	distSq int
}

func pt1(input string) (string, error) {
	points, e := getPoints(input)
	if e != nil {
		return "", fmt.Errorf("error getting points: %w", e)
	}
	pairs := getDistPairs(points)
	circuits := utils.NewDisjointSet[point3d]()
	for _, point := range points {
		circuits.Add(point)
	}
	for _, pair := range pairs[:1000] {
		rep1, rep2 := circuits.Find(pair.a), circuits.Find(pair.b)
		if rep1 != rep2 {
			circuits.Union(rep1, rep2)
		}
	}
	circuitLengths := make([]int, circuits.Count())
	for _, group := range circuits.Sets() {
		circuitLengths = append(circuitLengths, len(group))
	}
	sort.Ints(circuitLengths)
	total := circuitLengths[len(circuitLengths)-1] * circuitLengths[len(circuitLengths)-2] * circuitLengths[len(circuitLengths)-3]
	return strconv.Itoa(total), nil
}

func getDistPairs(points []point3d) []distPair {
	pairs := []distPair{}
	for i, a := range points {
		for _, b := range points[i+1:] {
			pairs = append(pairs, getDistPair(a, b))
		}
	}
	slices.SortFunc(pairs, func(a, b distPair) int {
		return a.distSq - b.distSq
	})
	return pairs
}

func getDistPair(a, b point3d) distPair {
	dx, dy, dz := a.x-b.x, a.y-b.y, a.z-b.z
	dist := dx*dx + dy*dy + dz*dz
	return distPair{a, b, dist}
}

func getPoints(input string) ([]point3d, error) {
	lines := utils.GetLines(input)
	points := make([]point3d, len(lines))
	for i, line := range lines {
		p, e := parseLine(line)
		if e != nil {
			return nil, fmt.Errorf("error parsing line: %w", e)
		}
		points[i] = p
	}
	return points, nil
}

func parseLine(line string) (point3d, error) {
	nums, e := utils.StringsToInts(strings.Split(line, ","))
	if e != nil {
		return point3d{}, fmt.Errorf("error converting number strings to ints: %w", e)
	}
	return point3d{nums[0], nums[1], nums[2]}, nil
}
