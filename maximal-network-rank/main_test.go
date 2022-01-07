package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maximalNetworkRank(n int, roads [][]int) int {
	lineCount := make([]int, n)
	directionCount := make([][]int, n)
	for idx := range directionCount {
		directionCount[idx] = make([]int, n)
	}
	for _, road := range roads {
		from, to := road[0], road[1]
		if from > to {
			from, to = to, from
		}
		directionCount[from][to] += 1
		lineCount[from]++
		lineCount[to]++
	}
	max := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			rank := lineCount[i] + lineCount[j] - directionCount[i][j]
			// fmt.Printf("lineCount[i]: %v\n", lineCount[i])
			// fmt.Printf("lineCount[j]: %v\n", lineCount[j])
			// fmt.Printf("directionCount[i][j]: %v\n", directionCount[i][j])
			if rank > max {
				max = rank
			}
		}
	}
	return max
}

func TestRank(t *testing.T) {
	assert.Equal(t, 1, maximalNetworkRank(2, [][]int{
		{1, 0},
	}))
	assert.Equal(t, 4, maximalNetworkRank(4, [][]int{
		{0, 1}, {0, 3}, {1, 2}, {1, 3},
	}))
	assert.Equal(t, 5, maximalNetworkRank(5, [][]int{
		{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4},
	}))
}
