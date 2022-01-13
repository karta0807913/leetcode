package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)-1
	for left < len(matrix[0]) && right >= 0 {
		if matrix[right][left] > target {
			right -= 1
		} else if matrix[right][left] < target {
			left += 1
		} else {
			return true
		}
	}
	return false
}

func TestSearch(t *testing.T) {
	assert.True(t, searchMatrix([][]int{
		{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25},
	}, 15))
	assert.False(t, searchMatrix([][]int{
		{1, 1},
	}, 15))
}
