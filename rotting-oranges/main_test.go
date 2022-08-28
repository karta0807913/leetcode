package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Point struct {
	X, Y int
}

func getOrange(grid [][]int, x, y int) int {
	if y < 0 || len(grid) <= y {
		return 0
	}
	if x < 0 || len(grid[y]) <= x {
		return 0
	}
	return grid[y][x]
}

func orangesRotting(grid [][]int) int {
	prev := []Point{}
	next := []Point{}
	goodOranges := 0
	for y, col := range grid {
		for x := range col {
			switch grid[y][x] {
			case 1:
				goodOranges += 1
			case 2:
				prev = append(prev, Point{
					X: x, Y: y,
				})
			}
		}
	}
	times := 0
	for len(prev) != 0 && goodOranges != 0 {
		times += 1
		for _, point := range prev {
			for _, delta := range []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				if getOrange(grid, point.X+delta.X, point.Y+delta.Y) == 1 {
					grid[point.Y+delta.Y][point.X+delta.X] = 2
					next = append(next, Point{
						X: point.X + delta.X, Y: point.Y + delta.Y,
					})
					goodOranges -= 1
				}
			}
		}
		next, prev = prev, next
		next = next[:0]
	}
	if goodOranges != 0 {
		return -1
	}
	return times
}

func TestRotten(t *testing.T) {
	assert := require.New(t)
	assert.Equal(2, orangesRotting([][]int{
		{2, 2, 2}, {2, 2, 0}, {0, 1, 1},
	}))

	assert.Equal(4, orangesRotting([][]int{
		{2, 1, 1}, {1, 1, 0}, {0, 1, 1},
	}))

	assert.Equal(-1, orangesRotting([][]int{
		{2, 1, 1}, {1, 1, 0}, {0, 0, 1},
	}))
}
