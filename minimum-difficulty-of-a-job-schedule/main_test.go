package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calMinDifficulty(index int, jobDifficulty, dp []int) int {
	difficulty := jobDifficulty[index]
	result := math.MaxInt
	for ; index < len(dp)-1; index++ {
		difficulty = max(jobDifficulty[index], difficulty)
		result = min(
			result,
			difficulty+dp[index+1],
		)
	}
	return result
}

func minDifficulty(jobDifficulty []int, d int) int {
	if d > len(jobDifficulty) {
		return -1
	}
	dp := make([]int, len(jobDifficulty))
	next := make([]int, len(jobDifficulty))
	maxVal := jobDifficulty[len(jobDifficulty)-1]
	for i := len(jobDifficulty) - 1; i >= 0; i-- {
		maxVal = max(jobDifficulty[i], maxVal)
		dp[i] = maxVal
	}
	for ; d != 1; d-- {
		for i := range dp {
			next[i] = calMinDifficulty(i, jobDifficulty, dp)
		}
		next, dp = dp[:len(dp)-1], next[:len(dp)-1]
	}
	return dp[0]
}

func TestDifficulty(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		7,
		minDifficulty([]int{6, 5, 4, 3, 2, 1}, 2),
	)
}
