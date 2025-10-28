package utils

import (
	"errors"
	"fmt"
	"slices"
)

type FinGrid[T comparable] struct {
	items []T
	width int
	height int
}

func GridFromString(s string) FinGrid[rune] {
	return GridFromLines(GetLines(s))
}

func GridFromSlice[T comparable](s []T, width int) FinGrid[T] {
	height := len(s) / width
	if len(s) % width != 0 {
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

func (g *FinGrid[T]) At(x int, y int) Option[T] {
	if x < 0 || y < 0 {
		return None[T]()
	}
	idx := y*g.width + x
	if idx > len(g.items) {
		return None[T]()
	}
	return Some(g.items[idx])
}

func (g *FinGrid[T]) Set(x, y int, value T) error {
	if x < 0 || y < 0 {
		return errors.New("Coordinates out of bounds!")
	}
	idx := y*g.width + x
	if idx > len(g.items) {
		return errors.New("Coordinates out of bounds!")
	}
	g.items[y*g.width + x] = value
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

func (g *FinGrid[T]) Find(item T) Option[Point] {
	i := slices.Index(g.items, item)
	if i >= 0 {
		return Some(Point{i % g.width, i / g.width})
	}
	return None[Point]()
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

func PrintGrid(g FinGrid[rune]) {
	for i, char := range g.items {
		if i % g.width == 0 {
			fmt.Print("\n")
		}
		fmt.Printf("%c", char)
	}
	println()
}