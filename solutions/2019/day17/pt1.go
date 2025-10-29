package day17

import (
	"advent-of-go/solutions/2019/intcode"
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 17,
		Part: 1,
		Calculator: pt1,
	}
}

var OFFSETS = [4]utils.Point{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 0, Y: -1}}

func pt1(input string) (string, error) {
	grid, e := getGrid(input)
	if e != nil {
		return "", fmt.Errorf("Error getting the scaffold grid: %w", e)
	}
	return strconv.Itoa(getTotal(grid)), nil
}

func getGrid(input string) (utils.FinGrid[rune], error) {
	outStr := strings.Builder{}
	out := func(n int) {
		outStr.WriteRune(rune(n))
	}
	in := func() int {
		println("Input called?")
		return -1
	}
	var grid utils.FinGrid[rune]
	if e := intcode.RunString(input, in, out); e != nil {
		return grid, fmt.Errorf("Error running intcode: %w", e)
	}
	return utils.GridFromString(outStr.String()), nil
}

func getTotal(grid utils.FinGrid[rune]) int {
	scaffoldPoints := grid.FindAll('#')
	total := 0
	for _, p := range scaffoldPoints {
		adjScaff := 0
		for _, offset := range OFFSETS {
			adj := p.Add(offset)
			if grid.At(adj.X, adj.Y).Or('.') == '#' {
				adjScaff++
			}
		}
		if adjScaff == 4 {
			total += p.X * p.Y
		}
	}
	return total
}