package day8

import (
	"advent-of-go/utils"
	"math"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 8,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	scores := [][3]int{}
	for i := 150; i <= len(input); i += 150 {
		scores = append(scores, calcScores(input[i - 150:i]))
	}
	best :=findBestLayerScore(scores)
	return strconv.Itoa(best), nil
}

func findBestLayerScore(allScores [][3]int) int {
	minRow := [3]int{math.MaxInt}
	for _, scoreRow := range allScores {
		if scoreRow[0] < minRow[0] {
			minRow = scoreRow
		}
	}
	return minRow[1] * minRow[2]
}

func calcScores(layer string) [3]int {
	scores := [3]int{}
	for _, char := range layer {
		if char == '0' {
			scores[0]++
		} else if char == '1' {
			scores[1]++
		} else if char == '2' {
			scores[2]++
		}
	}
	return scores
}