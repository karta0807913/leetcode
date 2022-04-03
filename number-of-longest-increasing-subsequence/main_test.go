package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	numCount := make([]int, len(nums))

	maxLength := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		numCount[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					numCount[i] = numCount[j]
				} else if dp[i] == dp[j]+1 {
					numCount[i] += numCount[j]
				}
			}
		}
		if dp[i] > maxLength {
			maxLength = dp[i]
		}
	}
	// fmt.Printf("dp: %v\n", dp)
	// fmt.Printf("numCount: %v\n", numCount)
	result := 0
	for i := range dp {
		if dp[i] == maxLength {
			result += numCount[i]
		}
	}
	return result
}

func TestNumber(t *testing.T) {
	assert.Equal(t, 8, findNumberOfLIS([]int{2, 2, 2, 2, 2, 2, 2, 2}))
	assert.Equal(t, 1, findNumberOfLIS([]int{2, 3, 2, 2, 2, 2, 2, 2}))
	assert.Equal(t, 2, findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	assert.Equal(t, 3, findNumberOfLIS([]int{1, 2, 4, 3, 5, 4, 7, 2}))
}
