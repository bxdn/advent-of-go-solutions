package day6

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 6,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	orbits, e := getOrbits(input)
	if e != nil {
		return "", fmt.Errorf("Error getting orbits: %w", e)
	}
	youMemo := map[string]int{}
	sanMemo := map[string]int{}
	youTotal := getNumOrbits("YOU", orbits, youMemo)
	sanTotal := getNumOrbits("SAN", orbits, sanMemo)
	maxCommon := 0
	for k, v1 := range youMemo {
		if _, ok := sanMemo[k]; ok {
			maxCommon = max(maxCommon, v1)
		}
	}
	minRequired := youTotal + sanTotal - 2 * maxCommon - 2
	return strconv.Itoa(minRequired), nil
}