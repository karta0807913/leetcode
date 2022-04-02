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

func calNext(current, previous int) int {
	previous -= current
	if previous <= 0 {
		return 1
	}
	return previous
}

func calculateMinimumHP(dungeon [][]int) int {
	lastColumn := len(dungeon) - 1
	lastRow := len(dungeon[0]) - 1
	dungeon[lastColumn][lastRow] = calNext(dungeon[lastColumn][lastRow], 1)
	for idx := lastRow - 1; idx >= 0; idx-- {
		dungeon[lastColumn][idx] = calNext(
			dungeon[lastColumn][idx],
			dungeon[lastColumn][idx+1],
		)
	}
	// fmt.Printf("dungeon: %v\n", dungeon)

	for columnIdx := lastColumn - 1; columnIdx >= 0; columnIdx-- {
		for rowIdx := lastRow; rowIdx >= 0; rowIdx-- {
			if rowIdx == lastRow {
				dungeon[columnIdx][rowIdx] = calNext(
					dungeon[columnIdx][rowIdx],
					dungeon[columnIdx+1][rowIdx],
				)
			} else {
				dungeon[columnIdx][rowIdx] = min(
					calNext(
						dungeon[columnIdx][rowIdx],
						dungeon[columnIdx][rowIdx+1],
					),
					calNext(
						dungeon[columnIdx][rowIdx],
						dungeon[columnIdx+1][rowIdx],
					),
				)
			}
		}
	}
	// fmt.Printf("dungeon: %v\n", dungeon)
	return dungeon[0][0]
}

func TestCalculateHealth(t *testing.T) {
	assert.Equal(t, 7, calculateMinimumHP([][]int{
		{-1, -2, -3},
	}))

	assert.Equal(t, 1, calculateMinimumHP([][]int{
		{100},
	}))

	assert.Equal(t, 3, calculateMinimumHP([][]int{
		{-1, -2, -3},
		{-1, 2, 3},
	}))

	assert.Equal(t, 7, calculateMinimumHP([][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}))

	assert.Equal(t, 3, calculateMinimumHP([][]int{
		{1, -3, 3},
		{0, -2, 0},
		{-3, -3, -3},
	}))
}
