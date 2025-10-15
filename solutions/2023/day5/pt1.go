package day5

import (
	"advent-of-go/utils"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type rangeMap struct {
	destination, source, length int
}

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 5,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	if len(lines) == 0 {
		return "", errors.New("Empty input!")
	}
	seeds := utils.FindPosInts(lines[0])
	lines = lines[3:]
	rangeMaps, e := getRanges(lines)
	if e != nil {
		return "", fmt.Errorf("Error getting mapping ranges: %w", e)
	}
	minLocation := math.MaxInt
	for _, seed := range seeds {
		location := getLocation(seed, rangeMaps)
		if location < minLocation {
			minLocation = location
		}
	}
	return strconv.Itoa(minLocation), nil
}

func getLocation(seed int, maps [][]rangeMap) int {
	for _, mapping := range maps {
		seed = mapValue(seed, mapping)
	}
	return seed
}

func mapValue(initial int, rangeMap []rangeMap) int {
	for _, mapping := range rangeMap {
		if mapping.source <= initial && initial < mapping.source + mapping.length {
			return mapping.destination + (initial - mapping.source)
		}
	}
	return initial
}

func getSeeds(line string) ([]int, error) {
	seeds := utils.FindPosInts(line)
	return seeds, nil
}

func getRanges(lines []string) ([][]rangeMap, error) {
	ranges := [][]rangeMap{}
	curRange := []rangeMap{}
	for _, line := range lines {
		if line == "" {
			ranges = append(ranges, curRange)
			curRange = []rangeMap{}
		} else if !strings.Contains(line, ":") {
			newRange, e := GetRange(line)
			if e != nil {
				return nil, fmt.Errorf("Error getting range from line %s: %w", line, e)
			}
			curRange = append(curRange, newRange)
		} 
	}
	ranges = append(ranges, curRange)
	return ranges, nil
}

func GetRange(line string) (rangeMap, error) {
	nums := utils.FindPosInts(line)
	if len(nums) != 3 {
		return rangeMap{}, errors.New("Malformed line: Range line should have 3 numbers!")
	}
	return rangeMap{nums[0], nums[1], nums[2]}, nil
}