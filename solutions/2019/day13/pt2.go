package day13

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 13,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	ops, e := utils.StringsToInts(strings.Split(input, ","))
	if e != nil {
		return "", fmt.Errorf("Error parsing input: %w", e)
	}
	ops[0] = 2
	grid := utils.NewInfGrid[int]()
	score, e := runGame(ops, grid)
	return strconv.Itoa(score), e
}