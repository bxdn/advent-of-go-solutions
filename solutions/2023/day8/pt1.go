package day8

import (
	"advent-of-go/utils"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var regex = regexp.MustCompile(`[A-Z]{3}`)

type node struct {
	left string
	right string
}

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 8,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
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
	return strconv.Itoa(traverse("AAA", directions, nodes, isZZZ)), nil
}

func getNodes(nodeLines []string) (map[string]node, error) {
	nodes := map[string]node{}
	for _, line := range nodeLines {
		matches := regex.FindAllString(line, -1)
		if len(matches) != 3 {
			return nil, errors.New("Malformed input: Every line should have 3 node labels!")
		}
		nodes[matches[0]] = node{matches[1], matches[2]}
	}
	return nodes, nil
}

type endValidator = func (string) bool

func isZZZ(label string) bool {return label == "ZZZ"}

func traverse(startingLabel string, directions string, nodes map[string]node, endSensor endValidator ) int {
	steps := 0
	curLabel := startingLabel
	for {
		for _, dir := range directions {
			steps++
			nextNode := nodes[curLabel]
			if dir == 'R' {
				curLabel = nextNode.right
			} else {
				curLabel = nextNode.left
			}
			if endSensor(curLabel) {
				return steps
			}
		}
	}
}