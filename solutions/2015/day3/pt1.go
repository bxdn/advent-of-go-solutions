package day3

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2015,
		Day:        3,
		Part:       1,
		Calculator: pt1,
	}
}

var offsets = map[rune]utils.Point{
	'^': {X: 0, Y: 1},
	'>': {X: 1, Y: 0},
	'v': {X: 0, Y: -1},
	'<': {X: -1, Y: 0},
}

func pt1(input string) (string, error) {
	cur := utils.Point{X: 0, Y: 0}
	visited := map[utils.Point]bool{}
	for _, char := range input {
		visited[cur] = true
		cur = cur.Add(offsets[char])
	}
	visited[cur] = true
	return strconv.Itoa(len(visited)), nil
}
