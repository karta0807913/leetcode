package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Point int

type UnionSet struct {
	islandMap   map[int]int
	islandCount int
}

func (set *UnionSet) Find(x int) int {
	parent, ok := set.islandMap[x]
	if !ok {
		set.islandMap[x] = x
		set.islandCount += 1
		return x
	}
	if parent != x {
		set.islandMap[x] = set.Find(parent)
		return set.islandMap[x]
	}
	return x
}

func (set *UnionSet) Union(x, y int) {
	x, y = set.Find(x), set.Find(y)
	if x != y {
		set.islandMap[x] = y
		set.islandCount -= 1
	}
}

func regionsBySlashes(grid []string) int {
	unionSet := UnionSet{
		islandMap:   make(map[int]int),
		islandCount: 0,
	}
	colLengnth := len(grid)
	for colIdx, str := range grid {
		for rowIdx, c := range str {
			currentPoint := Point(colIdx*colLengnth + rowIdx)
			if colIdx != 0 {
				unionSet.Union(
					Point((colIdx-1)*colLengnth+rowIdx).Down(),
					currentPoint.Top(),
				)
			}

			if rowIdx != 0 {
				unionSet.Union(
					Point(colIdx*colLengnth+rowIdx-1).Right(),
					currentPoint.Left(),
				)
			}

			switch c {
			case ' ':
				unionSet.Union(currentPoint.Left(), currentPoint.Right())
				unionSet.Union(currentPoint.Left(), currentPoint.Top())
				unionSet.Union(currentPoint.Left(), currentPoint.Down())
			case '/':
				unionSet.Union(currentPoint.Left(), currentPoint.Top())
				unionSet.Union(currentPoint.Right(), currentPoint.Down())
			case '\\':
				unionSet.Union(currentPoint.Left(), currentPoint.Down())
				unionSet.Union(currentPoint.Right(), currentPoint.Top())
			}
		}
	}
	return unionSet.islandCount
}

func (p Point) Top() int {
	return int(p) << 2
}

func (p Point) Down() int {
	return int(p)<<2 | 1
}

func (p Point) Left() int {
	return int(p)<<2 | 2
}

func (p Point) Right() int {
	return int(p)<<2 | 3
}

func TestSlashes(t *testing.T) {
	assert.Equal(t, 2, regionsBySlashes([]string{" /", "/ "}))
	assert.Equal(t, 1, regionsBySlashes([]string{" "}))
	assert.Equal(t, 2, regionsBySlashes([]string{" /"}))
}
