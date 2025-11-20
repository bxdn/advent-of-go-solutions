package day3

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2015,
		Day:        3,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	curSanta := utils.Point{X: 0, Y: 0}
	curRobo := utils.Point{X: 0, Y: 0}
	visited := map[utils.Point]bool{}
	for i, char := range input {
		if i&1 == 0 {
			visited[curSanta] = true
			curSanta = curSanta.Add(offsets[char])
		} else {
			visited[curRobo] = true
			curRobo = curRobo.Add(offsets[char])
		}
	}
	visited[curSanta] = true
	visited[curRobo] = true
	return strconv.Itoa(len(visited)), nil
}
