package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func computeCurrentValue(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min + 1
}

func maximalSquare(matrix [][]byte) (maxVal int) {
	cache := make([][]int, len(matrix))
	for colIdx, column := range matrix {
		cache[colIdx] = make([]int, len(column))
		for rowIdx, currentVal := range column {
			if colIdx == 0 || rowIdx == 0 {
				if maxVal < int(matrix[colIdx][rowIdx]-'0') {
					maxVal = int(matrix[colIdx][rowIdx] - '0')
				}
				cache[colIdx][rowIdx] = int(matrix[colIdx][rowIdx] - '0')
				continue
			}
			if currentVal != '0' {
				cache[colIdx][rowIdx] =
					computeCurrentValue(cache[colIdx][rowIdx-1], cache[colIdx-1][rowIdx], cache[colIdx-1][rowIdx-1])
				// fmt.Printf("cache: %v\n", cache)
				if maxVal < cache[colIdx][rowIdx] {
					maxVal = cache[colIdx][rowIdx]
				}
			}
		}
	}
	maxVal *= maxVal
	return
}

func TestSquare(t *testing.T) {
	assert.Equal(t, 4, maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'},
	}))

	assert.Equal(t, 1, maximalSquare([][]byte{
		{'0', '1'}, {'1', '0'},
	}))
	assert.Equal(t, 0, maximalSquare([][]byte{
		{'0'},
	}))
	assert.Equal(t, 1, maximalSquare([][]byte{
		{'1'},
	}))
}
