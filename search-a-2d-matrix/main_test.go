package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func searchMatrix(matrix [][]int, target int) bool {
	length := len(matrix) * len(matrix[0])
	left, right := 0, length-1
	for left <= right {
		i := (left + right) / 2
		current := matrix[i/len(matrix[0])][i%len(matrix[0])]
		if target == current {
			return true
		} else if target < current {
			right = i - 1
		} else {
			left = i + 1
		}
	}
	return false
}

func TestSearch(t *testing.T) {
	assert.True(t, searchMatrix([][]int{
		{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
	}, 3))
}
