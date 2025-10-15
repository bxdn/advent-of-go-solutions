package day6

import (
	"advent-of-go/utils"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile(`\d+`)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 6,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	if len(lines) != 2 {
		return "", errors.New("Malformed input: did not have the expected number of lines!")
	}
	pair, e := getData(lines[0], lines[1])
	if e != nil {
		return "", fmt.Errorf("Error getting data pair: %w", e)
	}
	return strconv.Itoa(getPairTotalFast(pair)), nil
}

func getPairTotalFast(pair timeDistancePair) int {
	timef, distancef := float64(pair.time), float64(pair.distance)
	rootf := (timef - math.Sqrt(math.Pow(timef, 2) - 4.0 * distancef)) / 2.0
	root := int(rootf)
	return pair.time - root * 2 - 1
}

func getData(timeLine, distanceLine string) (timeDistancePair, error) {
	timeNumStrs := regex.FindAllString(timeLine, -1)
	distanceStrs := regex.FindAllString(distanceLine, -1)
	timeNum, e := strconv.Atoi(strings.Join(timeNumStrs, ""))
	if e != nil {
		return timeDistancePair{}, fmt.Errorf("Error converting time number to int: %w", e)
	}
	distanceNum, e := strconv.Atoi(strings.Join(distanceStrs, ""))
	if e != nil {
		return timeDistancePair{}, fmt.Errorf("Error converting distance number to int: %w", e)
	}
	return timeDistancePair{timeNum, distanceNum}, nil
}