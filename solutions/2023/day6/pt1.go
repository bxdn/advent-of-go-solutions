package day6

import (
	"advent-of-go/utils"
	"errors"
	"fmt"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 6,
		Part: 1,
		Calculator: pt1,
	}
}

type timeDistancePair struct {
	time, distance int
}

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	if len(lines) != 2 {
		return "", errors.New("Malformed input: did not have the expected number of lines!")
	}
	pairs, e := getPairs(lines[0], lines[1])
	if e != nil {
		return "", fmt.Errorf("Error getting data pairs: %w", e)
	}
	result := 1
	for _, pair := range pairs {
		result *= getPairTotal(pair)
	}
	return strconv.Itoa(result), nil
}

func getPairTotal(pair timeDistancePair) int {
	numWays := 0
	for pressTime := range pair.time {
		remainder := pair.time - pressTime
		distance := pressTime * remainder
		if distance > pair.distance {
			numWays++
		}
	}
	return numWays
}

func getPairs(timeLine, distanceLine string) ([]timeDistancePair, error) {
	timeNums := utils.FindPosInts(timeLine)
	distanceNums := utils.FindPosInts(distanceLine)
	toRet := make([]timeDistancePair, len(timeNums))
	for i := range timeNums {
		toRet[i] = timeDistancePair{timeNums[i], distanceNums[i]}
	}
	return toRet, nil
}