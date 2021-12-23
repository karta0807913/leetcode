package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minPathSum(grid [][]int) int {
	for colIdx, column := range grid {
		for rowIdx := range column {
			left, top := math.MaxInt64, math.MaxInt64
			if rowIdx == 0 && colIdx == 0 {
				left = 0
			} else {
				if rowIdx-1 != -1 {
					left = grid[colIdx][rowIdx-1]
				}
				if colIdx-1 != -1 {
					top = grid[colIdx-1][rowIdx]
				}
			}
			grid[colIdx][rowIdx] += min(left, top)
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func TestPath(t *testing.T) {
	assert.Equal(t, 7, minPathSum([][]int{
		{1, 3, 1}, {1, 5, 1}, {4, 2, 1},
	}))
	assert.Equal(t, 12, minPathSum([][]int{
		{1, 2, 3}, {4, 5, 6},
	}))
}
