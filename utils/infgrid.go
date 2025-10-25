package utils

import (
	"fmt"
	"math"
)

type InfGrid[T comparable] struct {
	items                map[Point]T
	bottomLeft, topRight *Point
}

func NewInfGrid[T comparable]() InfGrid[T] {
	return InfGrid[T]{map[Point]T{}, &Point{math.MaxInt, math.MaxInt}, &Point{math.MinInt, math.MinInt}}
}

func (g *InfGrid[T]) At(x int, y int) Option[T] {
	if v, ok := g.items[Point{x, y}]; ok {
		return Some(v)
	}
	return None[T]()
}

func (g *InfGrid[T]) Set(x, y int, value T) {
	g.bottomLeft.X = min(g.bottomLeft.X, x)
	g.bottomLeft.Y = min(g.bottomLeft.Y, y)
	g.topRight.X = max(g.topRight.X, x)
	g.topRight.Y = max(g.topRight.Y, y)
	g.items[Point{x, y}] = value
}

func (g *InfGrid[T]) Find(item T) Option[Point] {
	for k, v := range g.items {
		if v == item {
			return Some(k)
		}
	}
	return None[Point]()
}

func (g *InfGrid[T]) FindAll(item T) []Point {
	all := []Point{}
	for k, v := range g.items {
		if v == item {
			all = append(all, k)
		}
	}
	return all
}

func (g *InfGrid[T]) Bounds() (Point, Point) {
	return *g.bottomLeft, *g.topRight
}

func (g *InfGrid[T]) ToFinGrid() (FinGrid[T], error) {
	newGrid := BlankGrid[T](g.topRight.X - g.bottomLeft.X + 1, g.topRight.Y - g.bottomLeft.Y + 1)
	for k, v := range g.items {
		if e := newGrid.Set(k.X - g.bottomLeft.X, k.Y - g.bottomLeft.Y, v); e != nil {
			return newGrid, fmt.Errorf("Error converting to finite grid: %w", e)
		}
	}
	return newGrid, nil
}