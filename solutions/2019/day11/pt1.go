package day11

import (
	"advent-of-go/solutions/2019/intcode"
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

type color int
const (
	Black color = iota
	White
)

type direction int
const (
	Up direction = iota
	Right
	Down
	Left
)

var offsets = [4]utils.Point{{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}}

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 11,
		Part: 1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	ops, e := utils.StringsToInts(strings.Split(input, ","))
	if e != nil {
		return "", fmt.Errorf("Error parsing input: %w", e)
	}
	grid := utils.NewInfGrid[color]()
	runBot(ops, grid)
	numWhite := len(grid.FindAll(White))
	numBlack := len(grid.FindAll(Black))
	return strconv.Itoa(numWhite + numBlack), nil
}

func runBot(ops []int, grid utils.InfGrid[color]) {
	botX, botY := 0, 0
	botOrientation := Up
	painting := true
	out := func (n int) {
		if painting {
			grid.Set(botX, botY, color(n))
		} else {
			if n == 0 {
				botOrientation = botOrientation - 1
				if botOrientation == -1 {
					botOrientation = Left
				}
			} else {
				botOrientation = (botOrientation + 1) % 4
			}
			botX += offsets[botOrientation].X
			botY += offsets[botOrientation].Y
		}
		painting = !painting
	}
	in := func () int {
		return int(grid.At(botX, botY).Or(Black))
	}
	intcode.Run(ops, in, out)
}