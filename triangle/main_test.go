package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotal(triangle [][]int) int {
	for colIdx := len(triangle) - 2; colIdx >= 0; colIdx-- {
		for rowIdx := range triangle[colIdx] {
			triangle[colIdx][rowIdx] += min(triangle[colIdx+1][rowIdx], triangle[colIdx+1][rowIdx+1])
		}
	}
	return triangle[0][0]
}

func TestTotal(t *testing.T) {
	assert.Equal(t, 1, minimumTotal([][]int{
		{1},
	}))
	assert.Equal(t, 3, minimumTotal([][]int{
		{1}, {2, 3},
	}))
	assert.Equal(t, 11, minimumTotal([][]int{
		{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3},
	}))
	assert.Equal(t, 16, minimumTotal([][]int{
		{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}, {2, 5, 7, 1, 2},
	}))
}
