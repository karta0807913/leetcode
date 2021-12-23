package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 || obstacleGrid[0][0] == 1 {
		return 0
	}
	obstacleGrid[0][0] = -1
	for colIdx, row := range obstacleGrid {
		for rowIdx, val := range row {
			if val == 1 {
				continue
			}
			if rowIdx-1 != -1 && obstacleGrid[colIdx][rowIdx-1] != 1 {
				obstacleGrid[colIdx][rowIdx] += obstacleGrid[colIdx][rowIdx-1]
			}
			if colIdx-1 != -1 && obstacleGrid[colIdx-1][rowIdx] != 1 {
				obstacleGrid[colIdx][rowIdx] += obstacleGrid[colIdx-1][rowIdx]
			}
		}
	}
	return -obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

func TestPath(t *testing.T) {
	assert.Equal(t, 3, uniquePathsWithObstacles([][]int{
		{0, 1, 0}, {0, 0, 0}, {0, 0, 0},
	}))
	assert.Equal(t, 2, uniquePathsWithObstacles([][]int{
		{0, 0, 0}, {0, 1, 0}, {0, 0, 0},
	}))
	assert.Equal(t, 1, uniquePathsWithObstacles([][]int{
		{0, 1}, {0, 0},
	}))
	assert.Equal(t, 0, uniquePathsWithObstacles([][]int{
		{0, 1}, {0, 1},
	}))
	assert.Equal(t, 0, uniquePathsWithObstacles([][]int{
		{0, 1, 0}, {0, 1, 0}, {0, 1, 0},
	}))
	assert.Equal(t, 0, uniquePathsWithObstacles([][]int{
		{1, 1, 0}, {0, 1, 0}, {0, 1, 0},
	}))
	assert.Equal(t, 0, uniquePathsWithObstacles([][]int{
		{1, 0},
	}))
	assert.Equal(t, 0, uniquePathsWithObstacles([][]int{
		{1, 0},
	}))
}
