package utils

import (
	"errors"
	"fmt"
	"slices"
)

type FinGrid[T comparable] struct {
	items  []T
	width  int
	height int
}

func GridFromString(s string) FinGrid[rune] {
	return GridFromLines(GetLines(s))
}

func GridFromSlice[T comparable](s []T, width int) FinGrid[T] {
	height := len(s) / width
	if len(s)%width != 0 {
		height++
	}
	return FinGrid[T]{s, width, height}
}

func GridFromLines(lines []string) FinGrid[rune] {
	runes := StringsToRunes(lines)
	return GridFrom2dSlice(runes)
}

func GridFrom2dSlice[T comparable](slices [][]T) FinGrid[T] {
	width := 0
	if len(slices) != 0 {
		width = len(slices[0])
	}
	return FinGrid[T]{Flatten(slices), width, len(slices)}
}

func BlankGrid[T comparable](width, height int) FinGrid[T] {
	items := make([]T, width*height)
	return FinGrid[T]{items, width, height}
}

func (g *FinGrid[T]) Dims() (int, int) {
	return g.width, g.height
}

func (g *FinGrid[T]) At(x int, y int) CmpOption[T] {
	if x < 0 || y < 0 || x >= g.width || y >= g.height {
		return CNone[T]()
	}
	idx := y*g.width + x
	if idx >= len(g.items) {
		return CNone[T]()
	}
	return CSome(g.items[idx])
}

func (g *FinGrid[T]) Set(x, y int, value T) error {
	if x < 0 || y < 0 || x >= g.width || y >= g.height {
		return errors.New("coordinates out of bounds")
	}
	idx := y*g.width + x
	if idx >= len(g.items) {
		return errors.New("cordinates out of bounds")
	}
	g.items[idx] = value
	return nil
}

func (g *FinGrid[T]) AtP(p Point) CmpOption[T] {
	if p.X < 0 || p.Y < 0 || p.X >= g.width || p.Y >= g.height {
		return CNone[T]()
	}
	idx := p.Y*g.width + p.X
	if idx >= len(g.items) {
		return CNone[T]()
	}
	return CSome(g.items[idx])
}

func (g *FinGrid[T]) SetP(p Point, value T) error {
	if p.X < 0 || p.Y < 0 || p.X >= g.width || p.Y >= g.height {
		return errors.New("cordinates out of bounds")
	}
	idx := p.Y*g.width + p.X
	if idx >= len(g.items) {
		return errors.New("cordinates out of bounds")
	}
	g.items[idx] = value
	return nil
}

type Point struct {
	X, Y int
}

func (p *Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

func (p *Point) Sub(other Point) Point {
	return Point{p.X - other.X, p.Y - other.Y}
}

func (g *FinGrid[T]) Find(item T) CmpOption[Point] {
	i := slices.Index(g.items, item)
	if i >= 0 {
		return CSome(Point{i % g.width, i / g.width})
	}
	return CNone[Point]()
}

func (g *FinGrid[T]) FindAll(item T) []Point {
	all := []Point{}
	for i, v := range g.items {
		if v == item {
			all = append(all, Point{i % g.width, i / g.width})
		}
	}
	return all
}

func (g *FinGrid[T]) Points() func(func(Point) bool) {
	return func(yield func(Point) bool) {
		for y := range g.height {
			for x := range g.width {
				if !yield(Point{x, y}) {
					return
				}
			}
		}
	}
}

var adjOffsets = [4]Point{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

var adjCOffsets = [8]Point{{-1, -1}, {-1, 1}, {-1, 0}, {1, -1}, {1, 1}, {1, 0}, {0, 1}, {0, -1}}

func (g *FinGrid[T]) Adj(p Point) func(func(CmpOption[T]) bool) {
	return func(yield func(CmpOption[T]) bool) {
		for _, off := range adjOffsets {
			if !yield(g.AtP(p.Add(off))) {
				return
			}
		}
	}
}

func (g *FinGrid[T]) AdjC(p Point) func(func(CmpOption[T]) bool) {
	return func(yield func(CmpOption[T]) bool) {
		for _, off := range adjCOffsets {
			if !yield(g.AtP(p.Add(off))) {
				return
			}
		}
	}
}

func PrintGrid(g FinGrid[rune]) {
	for i, char := range g.items {
		if i%g.width == 0 {
			fmt.Print("\n")
		}
		fmt.Printf("%c", char)
	}
	println()
}
