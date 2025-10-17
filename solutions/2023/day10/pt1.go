package day10

import (
	"advent-of-go/utils"
	"errors"
	"fmt"
	"slices"
	"strconv"
)

type Direction int

const (
	North = iota
	South
	East
	West
)

var directions []Direction = []Direction{North, South, East, West}

var transformations map[rune]map[Direction]Direction = map[rune]map[Direction]Direction{
	'|': {North: North, South: South},
	'-': {East: East, West: West},
	'7': {East: South, North: West},
	'L': {South: East, West: North},
	'F': {North: East, West: South},
	'J': {East: North, South: West},
}

var vertexChars []rune = []rune{'J', 'F', 'L', '7'}

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 10,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	grid := utils.GridFromString(input)
	startingPoint, e := grid.Find('S').OrErr("Input does not have starting point!")
	if e != nil {
		return "", fmt.Errorf("Issue getting starting point: %w", e)
	}
	totalSteps, _, e := FindLoopCountSteps(grid, startingPoint)
	if e != nil {
		return "", fmt.Errorf("Error getting total steps in loop: %w", e)
	}
	return strconv.Itoa(totalSteps / 2), nil
}

func FindLoopCountSteps(grid utils.Grid[rune], startingPoint utils.Point) (int, []utils.Point, error) {
	for _, direction := range directions {
		nextX, nextY := transform(direction, startingPoint.X, startingPoint.Y)
		nextRune := grid.At(nextX, nextY).Or('.')
		transformation, ok := transformations[nextRune]
		if ok {
			_, ok = transformation[direction]
			if ok {
				steps, vertices, finalDir, e := CountStepsInRestOfLoop(nextX, nextY, direction, grid, []utils.Point{})
				if e != nil {
					return 0, nil, fmt.Errorf("Error getting steps: %w", e)
				}
				if finalDir != direction {
					vertices = append(vertices, startingPoint)
				}
				return 1 + steps, vertices, nil
			}
		}
		
	}
	return 0, nil, errors.New("No connecting pipes to starting point!")
}

func CountStepsInRestOfLoop(x, y int, direction Direction, grid utils.Grid[rune], vertices []utils.Point) (int, []utils.Point, Direction, error) {
	currentChar := grid.At(x, y).Or('.')
	if currentChar == 'S' {
		return 0, vertices, direction, nil
	}
	curPoint := utils.Point{X: x, Y: y}
	if slices.Contains(vertexChars, currentChar) {
		vertices = append(vertices, curPoint)
	}
	transformation, ok := transformations[currentChar]
	if !ok {
		return 0, nil, 0, errors.New("Not a loop: Ran into ground or border!")
	}
	nextDir, ok := transformation[direction]
	if !ok {
		return 0, nil, 0, errors.New("Not a loop: Ran into non-joining pipe!")
	}
	nextX, nextY := transform(nextDir, x, y)
	steps, vertices, finalDir, e := CountStepsInRestOfLoop(nextX, nextY, nextDir, grid, vertices)
	if e != nil {
		return 0, nil, 0, fmt.Errorf("Error Recursively getting steps: %w", e)
	}
	return 1 + steps, vertices, finalDir, nil
}

func transform(dir Direction, x, y int) (int, int) {
	switch dir {
		case North: return x, y - 1
		case South: return x, y + 1
		case East: return x + 1, y
		case West: return x - 1, y
	}
	return 0, 0
}