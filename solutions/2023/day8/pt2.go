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
	return lcm(individualSteps)
}

func endsWithZ(label string) bool {return label[2] == 'Z'}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	result := numbers[0]
	for _, num := range numbers[1:] {
		result = (result * num) / gcd(result, num)
	}
	return result
}
