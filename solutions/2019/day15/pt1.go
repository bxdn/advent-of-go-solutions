package day15

import (
	"advent-of-go/solutions/2019/intcode"
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2019, 
		Day: 15,
		Part: 1,
		Calculator: pt1,
	}
}

var OFFSETS = [5]utils.Point{{}, {X: 0, Y: -1}, {X: 0, Y: 1}, {X: 1, Y: 0}, {X: -1, Y: 0}}

func pt1(input string) (string, error) {
	grid, e := dfs(input)
	if e != nil {
		return "", fmt.Errorf("Error Traversing grid with robot: %w", e)
	}
	steps := bfs(grid, utils.Point{})
	goal, e := grid.Find(2).OrErr("Oxygen system not found")
	if e != nil {
		return "", e
	}
	return strconv.Itoa(steps[goal]), e
}

func dfs(program string) (utils.InfGrid[int], error) {
	grid := utils.NewInfGrid[int]()
	grid.Set(0, 0, 1)

	// shared state between in/out functions

	next := utils.Point{}
	visitStack := []utils.Point{{}}
	visited := map[utils.Point]bool{{}: true}
	backward := false


	in := func() int {
		current := utils.Last(visitStack).OrPanic("Guaranteed stack isn't empty above")
		// Try to go forward
		dir := findForwardDir(current, visited)
		if dir != 0 {
			backward = false
			next = current.Add(OFFSETS[dir])
			return dir
		}

		// Go back
		visitStack, _ = utils.Pop(visitStack)
		if len(visitStack) == 0 {
			return intcode.HALT_CODE
		}
		backward = true
		next = utils.Last(visitStack).OrPanic("Guaranteed stack isn't empty above")
		return findBackwardDir(current, next)
	}

	out := func(n int) {
		if backward {
			return
		}
		visited[next] = true
		grid.Set(next.X, next.Y, n)
		switch n {
			case 1: fallthrough
			case 2: visitStack = append(visitStack, next)
		}
	}

	e := intcode.RunString(program, in, out)
	return grid, e
}

func findBackwardDir(current, next utils.Point) int {
	offset := next.Sub(current)
	for i := range 4 {
		dir := i + 1
		if OFFSETS[dir] == offset {
			return dir
		}
	}
	// Shouldn't be reachable
	return intcode.HALT_CODE
}

func findForwardDir(current utils.Point, visited map[utils.Point]bool) int {
	for i := range 4 {
		dir := i + 1
		potential := OFFSETS[dir].Add(current)
		if _, ok := visited[potential]; !ok {
			return dir
		}
	}
	return 0
}


func bfs(g utils.InfGrid[int], from utils.Point) map[utils.Point]int {
	queue := []utils.Point{from}
	visited := map[utils.Point]int{from: 0}
	for len(queue) > 0 {
		var opt utils.Option[utils.Point]
		queue, opt = utils.Dequeue(queue)
		current := opt.OrPanic("Guaranteed not empty")
		for i := range 4 {
			offset := OFFSETS[i + 1]
			next := current.Add(offset)
			val, exists := g.At(next.X, next.Y).Get()
			if _, found := visited[next]; !found && exists && val != 0 {
				visited[next] = visited[current] + 1
				queue = append(queue, next)
			}
		}
	}
	return visited
}