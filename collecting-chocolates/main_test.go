package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCost(nums []int, x int) int64 {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1e9 + 1
	}
	bestSum := int64(1e9+1) * int64(len(nums))
	for shift := range nums {
		sum := int64(0)
		for i := range nums {
			shiftedIndex := (i + shift) % len(nums)
			dp[shiftedIndex] = min(nums[i], dp[shiftedIndex])
			sum += int64(dp[shiftedIndex])
		}
		sum += int64(x) * int64(shift)
		if sum < bestSum {
			bestSum = sum
		}
	}
	return bestSum
}

func TestCost(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		int64(186),
		minCost([]int{169, 37, 58, 175, 3, 9}, 47),
	)
	assert.Equal(
		int64(119),
		minCost([]int{31, 25, 18, 59}, 27),
	)
	assert.Equal(
		int64(13),
		minCost([]int{20, 1, 15}, 5),
	)
	assert.Equal(
		int64(6),
		minCost([]int{1, 2, 3}, 4),
	)
}
