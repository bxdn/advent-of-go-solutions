package day6

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 6,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	orbits, e := getOrbits(input)
	if e != nil {
		return "", fmt.Errorf("Error getting orbits: %w", e)
	}
	memo := map[string]int{}
	total := 0
	for k := range orbits {
		total += getNumOrbits(k, orbits, memo)
	}
	return strconv.Itoa(total), nil
}

func getOrbits(input string) (map[string]string, error) {
	lines := utils.GetLines(input)
	orbits :=map[string]string{}
	for i, line := range lines {
		parent, child, ok := strings.Cut(line, ")")
		if !ok {
			return nil, fmt.Errorf("Malformed line: %s at line %d", line, i)
		}
		orbits[child] = parent
	}
	return orbits, nil
}

func getNumOrbits(key string, orbits map[string]string, memo map[string]int) int {
	if v, ok := memo[key]; ok {
		return v
	}
	parent, ok := orbits[key]
	if !ok {
		return 0
	}
	toRet := 1 + getNumOrbits(parent, orbits, memo)
	memo[key] = toRet
	return toRet
}