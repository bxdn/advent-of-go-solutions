package day13

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
		Day: 13,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	ops, e := utils.StringsToInts(strings.Split(input, ","))
	if e != nil {
		return "", fmt.Errorf("Error converting program to ints: %w", e)
	}
	grid := utils.NewInfGrid[int]()
	_, e = runGame(ops, grid)
	return strconv.Itoa(len(grid.FindAll(2))), e
}

func runGame(program []int, grid utils.InfGrid[int]) (int, error) {
	count, x, y := 0, 0, 0
	score := 0
	out := func (n int) {
		switch count {
			case 0: x = n
			case 1: y = n
			case 2: {
				if x == -1 && y == 0 {
					score = n
				} else {
					grid.Set(x, y, n)
				}
			}
		}
		count = (count + 1) % 3
	}
	in := func() int {
		paddleX := grid.Find(3).Or(utils.Point{}).X
		ballX := grid.Find(4).Or(utils.Point{}).X
		if paddleX < ballX {
			return 1
		} else if ballX < paddleX {
			return -1
		}
		return 0
	}
	e := intcode.Run(program, in, out)
	return score, e
}