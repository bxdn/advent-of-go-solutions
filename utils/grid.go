package utils

import "errors"

type Grid[T any] struct {
	items []T
	width int
	height int
}

func GridFromLines(lines []string) Grid[rune] {
	runes := StringsToRunes(lines)
	return GridFrom2dSlice(runes)
}

func GridFrom2dSlice[T any](slices [][]T) Grid[T] {
	width := 0
	if len(slices) != 0 {
		width = len(slices[0])
	}
	return Grid[T]{Flatten(slices), width, len(slices)}
}

func BlankGrid[T any](width, height int) Grid[T] {
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