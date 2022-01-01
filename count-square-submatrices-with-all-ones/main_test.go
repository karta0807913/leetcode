package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minSquare(a, b, c int) int {
	min := a

	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min + 1
}

func countSquares(matrix [][]int) int {
	sum := 0
	for colIdx, column := range matrix {
		for rowIdx, val := range column {
			if val == 0 || colIdx == 0 || rowIdx == 0 {
				sum += matrix[colIdx][rowIdx]
				continue
			}
			matrix[colIdx][rowIdx] = minSquare(matrix[colIdx-1][rowIdx-1], matrix[colIdx][rowIdx-1], matrix[colIdx-1][rowIdx])
			sum += matrix[colIdx][rowIdx]
		}
	}
	return sum
}

func TestSquares(t *testing.T) {
	assert.Equal(t, 15, countSquares([][]int{
		{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1},
	}))
}
