package utils

import (
	"errors"
	"slices"
)

type Grid[T comparable] struct {
	items []T
	width int
	height int
}

func GridFromString(s string) Grid[rune] {
	return GridFromLines(GetLines(s))
}

func GridFromLines(lines []string) Grid[rune] {
	runes := StringsToRunes(lines)
	return GridFrom2dSlice(runes)
}

func GridFrom2dSlice[T comparable](slices [][]T) Grid[T] {
	width := 0
	if len(slices) != 0 {
		width = len(slices[0])
	}
	return Grid[T]{Flatten(slices), width, len(slices)}
}

func BlankGrid[T comparable](width, height int) Grid[T] {
	items := make([]T, width*height)
	return Grid[T]{items, width, height}
}

func (g *Grid[T]) Dims() (int, int) {
	return g.width, g.height
}

func (g *Grid[T]) At(x int, y int) Option[T] {
	if x < 0 || y < 0 {
		return None[T]()
	}
	idx := y*g.width + x
	if idx > len(g.items) {
		return None[T]()
	}
	return Some(g.items[idx])
}

func (g *Grid[T]) Set(x, y int, value T) error {
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

func (g *Grid[T]) Find(item T) Option[Point] {
	i := slices.Index(g.items, item)
	if i >= 0 {
		return Some(Point{i % g.width, i / g.width})
	}
	return None[Point]()
}

func (g *Grid[T]) FindAll(item T) []Point {
	all := []Point{}
	for i, v := range g.items {
		if v == item {
			all = append(all, Point{i % g.width, i / g.width})
		}
	}
	return all
}