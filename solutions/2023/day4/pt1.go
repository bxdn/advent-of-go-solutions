package day4

import (
	"advent-of-go/utils"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile(`\d+`)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 4,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	total := 0
	for _, line := range lines {
		lineTotal, e := getTotal(line)
		if e != nil {
			return "", fmt.Errorf("Error getting line total: %w", e)
		}
		total += lineTotal
	}
	return strconv.Itoa(total), nil
}

// Gets the total score of the card
func getTotal(line string) (int, error) {
	totalWinners, e := getNumWinners(line)
	if e != nil {
		return 0, fmt.Errorf("Error getting the number of winners: %w", e)
	}
	if totalWinners > 0 {
		return 1 << (totalWinners - 1), nil
	}
	return 0, nil
}

// Gets the number of winning draws on a given line
func getNumWinners(line string) (int, error) {
	winners, drawn, e := getNums(line)
	if e != nil {
		return 0, fmt.Errorf("Error getting the numbers: %w", e)
	}
	totalWinners := 0
	for _, num := range drawn {
		if slices.Contains(winners, num) {
			totalWinners++
		}
	}
	return totalWinners, nil
}

// Gets the []string representation of the winning and drawn numbers
func getNums(line string) ([]string, []string, error) {
	toks := strings.Split(line, ":")
	if len(toks) != 2 {
		return nil, nil, errors.New("Misformatted data: expected : delimited row")
	}
	toks = strings.Split(toks[1], "|")
	if len(toks) != 2 {
		return nil, nil, errors.New("Misformatted data: expected | delimited row")
	}
	winnersStr, drawnStr := toks[0], toks[1]
	winners := regex.FindAllString(winnersStr, -1)
	drawn := regex.FindAllString(drawnStr, -1)
	return winners, drawn, nil
}