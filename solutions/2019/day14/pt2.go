package day14

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 14,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	costs, e := getCosts(input)
	if e != nil {
		return "", fmt.Errorf("Error getting costs: %w", e)
	}
	maxPerFuel := get1FuelTotal(costs)
	remaining := map[string]int{}
	oreRemaining := 1_000_000_000_000
	totalFuel := 0
	for {
		fuelAttempt := max(1, oreRemaining / maxPerFuel)
		needed := []resourcePair{{fuelAttempt, "FUEL"}}
		roundTotal := 0
		for len(needed) > 0 {
			var numOre int
			numOre, needed = processItem(needed, remaining, costs)
			roundTotal += numOre
		}
		if oreRemaining < roundTotal {
			return strconv.Itoa(totalFuel), nil
		}
		totalFuel += fuelAttempt
		oreRemaining -= roundTotal
	}
}