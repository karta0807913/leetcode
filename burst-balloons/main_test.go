package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxCoins(nums []int) int {
	nums = append(nums, 1)
	nums = append([]int{1}, nums...)
	dp := make([][]int, len(nums))
	// fmt.Printf("nums: %v\n", nums)
	for idx := range nums {
		dp[idx] = make([]int, len(nums))
	}

	for length := 0; length < len(nums)-2; length++ {
		// [start, end]
		for start := 1; start < len(nums)-1-length; start++ {
			end := start + length
			// fmt.Printf("start: %v, end: %v\n", start, end)
			for i := start; i <= end; i++ {
				res := dp[start][i-1] + nums[start-1]*nums[i]*nums[end+1] + dp[i+1][end]
				// fmt.Printf("res: %v\n", res)
				dp[start][end] = max(res, dp[start][end])
			}
			// fmt.Printf("result: %v\n", dp[start][end])
		}
	}
	return dp[1][len(nums)-2]
}

func TestCoins(t *testing.T) {
	assert := require.New(t)

	assert.Equal(3, maxCoins([]int{3}))
	assert.Equal(20, maxCoins([]int{3, 5}))
	assert.Equal(133, maxCoins([]int{3, 5, 7}))
}
