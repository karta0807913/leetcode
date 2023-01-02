package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func countSubarrays(nums []int, k int) int {
	result := 0
	sum := 0
	startIndex := len(nums)
	dp := make(map[int]int)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			startIndex = i
			break
		}
		if nums[i] < k {
			sum--
		} else if nums[i] > k {
			sum++
		}
		dp[sum]++
	}
	for ; startIndex < len(nums); startIndex++ {
		if nums[startIndex] < k {
			sum--
		} else if nums[startIndex] > k {
			sum++
		}
		result += dp[sum]
		result += dp[sum-1]
	}
	return result
}

func TestSubarrays(t *testing.T) {
	assert := require.New(t)
	// assert.Equal(6, countSubarrays([]int{3, 3, 3}, 3))
	assert.Equal(3, countSubarrays([]int{3, 2, 1, 4, 5}, 4))
	assert.Equal(1, countSubarrays([]int{2, 3, 1}, 3))
}
