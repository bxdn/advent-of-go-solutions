package day5

import (
	"advent-of-go/utils"
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 5,
		Part: 2,
		Calculator: pt2,
	}
}

type seedRange struct {
	start, length int
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	if len(lines) == 0 {
		return "", errors.New("Empty input!")
	}
	seedRanges, e := getSeedRanges(lines[0])
	if e != nil {
		return "", fmt.Errorf("Error getting seeds: %w", e)
	}
	lines = lines[3:]
	rangeMaps, e := getRanges(lines)
	if e != nil {
		return "", fmt.Errorf("Error getting mapping ranges: %w", e)
	}
	minLocation := math.MaxInt
	for _, seedRange := range seedRanges {
		location := getMinLocationFromRange(seedRange, rangeMaps)
		if location < minLocation {
			minLocation = location
		}
	}
	return strconv.Itoa(minLocation), nil
}

func getSeedRanges(line string) ([]seedRange, error) {
	seeds := utils.FindPosInts(line)
	if len(seeds) % 2 != 0 {
		return nil, errors.New("Malformed data: partial seed range!")
	}
	ranges := make([]seedRange, len(seeds) / 2)
	for i := 0; i < len(seeds); i += 2 {
		ranges[i / 2] = seedRange{seeds[i], seeds[i + 1]}
	}
	return ranges, nil
}

func getMinLocationFromRange(initialRange seedRange, rangeMaps [][]rangeMap) int {
	seedRanges := []seedRange{initialRange}
	for _, rangeMap := range rangeMaps {
		mappedSeedRanges := []seedRange{}
		for _, sRange := range seedRanges {
			mappedSeedRanges = append(mappedSeedRanges, getRangeMapping(sRange, rangeMap)...)
		}
		seedRanges = mappedSeedRanges
	}
	minLocation := math.MaxInt
	for _, sRange := range seedRanges {
		if sRange.start < minLocation {
			minLocation = sRange.start
		}
	}
	return minLocation
}

func getRangeMapping(seedRng seedRange, rangeMaps []rangeMap) []seedRange {
	result := []seedRange{}
	mappedSourceRanges := []seedRange{}
	for _, mapping := range rangeMaps {
		intMin, intMax := findIntersection(seedRng, mapping)
		intLength := intMax - intMin
		if intLength > 0 {
			result = append(result, seedRange{intMin + (mapping.destination - mapping.source), intLength})
			mappedSourceRanges = append(mappedSourceRanges, seedRange{intMin, intLength})
		}
	}
	result = append(result, getUnmappable(seedRng, mappedSourceRanges)...)
	return result
}

func getUnmappable(originalRange seedRange, mappedSourceRanges []seedRange) []seedRange {
	mapped := []seedRange{}
	slices.SortFunc(mappedSourceRanges, func(a, b seedRange) int { return a.start - b.start })
	seedNum := originalRange.start
	for _, fromMap := range mappedSourceRanges {
		if fromMap.start > seedNum {
			mapped = append(mapped, seedRange{seedNum, fromMap.start - seedNum})
		}
		seedNum = fromMap.start + fromMap.length + 1
	}
	if seedNum < originalRange.start + originalRange.length {
		mapped = append(mapped, seedRange{seedNum, originalRange.start + originalRange.length - seedNum})
	}
	return mapped
}

func findIntersection(sdRange seedRange, mapping rangeMap) (int, int) {
	seedMin, seedMax := sdRange.start, sdRange.start + sdRange.length
	mappingMin, mappingMax := mapping.source, mapping.source + mapping.length
	return max(seedMin, mappingMin), min(seedMax, mappingMax)
}