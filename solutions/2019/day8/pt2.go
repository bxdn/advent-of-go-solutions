package day8

import (
	"advent-of-go/utils"
	"fmt"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2019,
		Day:        8,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	scores := make([]int, 150)
	for i := range scores {
		scores[i] = 2
	}
	for i, char := range input {
		scoresIdx := i % 150
		if scores[scoresIdx] == 2 {
			if char == '0' {
				scores[scoresIdx] = 0
			} else if char == '1' {
				scores[scoresIdx] = 1
			}
		}
	}
	grid := utils.GridFromSlice(scores, 25)
	png, e := utils.GridToPng(grid)
	if e != nil {
		return "", fmt.Errorf("error creating image string: %w", e)
	}
	return utils.DetectText(png)
}
