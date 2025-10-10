package day4

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:     2023,
		Day:      4,
		Part:     2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	cardSet := getInitialCardSet(len(lines))
	totalCards := 0
	for i, line := range lines {
		lineTotal, e := propagateCards(cardSet[i:], line)
		if e != nil {
			return "", fmt.Errorf("Error getting line total: %w", e)
		}
		totalCards += lineTotal
	}
	return strconv.Itoa(totalCards), nil
}

// Propagates copies of current card, then returns the amount of the current card for summing
func propagateCards(nextCards []int, line string) (int, error) {
	numCurrentCard := nextCards[0]
	totalWinners, e := getNumWinners(line)
	if e != nil {
		return 0, fmt.Errorf("Error getting the number of winners: %w", e)
	}
	for i := 1; i <= totalWinners; i++ {
		nextCards[i] += numCurrentCard
	}
	return numCurrentCard, nil
}

// Start with 1 card each
func getInitialCardSet(numCards int) []int {
	toRet := make([]int, numCards)
	for i := 0; i < numCards; i++ {
		toRet[i] = 1
	}
	return toRet
}