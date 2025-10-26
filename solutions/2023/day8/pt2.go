package day8

import (
	"advent-of-go/utils"
	"errors"
	"fmt"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 8,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	if len(lines) == 0 {
		return "", errors.New("Empty input!")
	}
	directions := lines[0]
	nodeLines := lines[2:]
	nodes, e := getNodes(nodeLines)
	if e != nil {
		return "", fmt.Errorf("Error getting nodes: %w", e)
	}
	return strconv.Itoa(traverseGhosts(directions, nodes)), nil
}

func traverseGhosts(directions string, nodes map[string]node) int {
	individualSteps := []int{}
	for k := range nodes {
		if k[2] == 'A' {
			individualSteps = append(individualSteps, traverse(k, directions, nodes, endsWithZ))
		}
	}
	return utils.Lcm(individualSteps)
}

func endsWithZ(label string) bool {return label[2] == 'Z'}
