package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func largestDivisibleSubset(nums []int) []int {
	dp := make([]int, len(nums))
	sort.Ints(nums)
	maxIdx := 0
	maxVal := 1
	for idx := range dp {
		dp[idx] = 1
		for compareIdx := idx - 1; compareIdx >= 0; compareIdx-- {
			if nums[idx]%nums[compareIdx] == 0 {
				if dp[compareIdx] >= dp[idx] {
					dp[idx] = dp[compareIdx] + 1
				}
			}
			if dp[idx] > maxVal {
				maxIdx = idx
				maxVal = dp[idx]
			}
		}
	}
	// fmt.Printf("dp: %v\n", dp)
	insertIdx := len(dp) - 1
	dp[insertIdx] = nums[maxIdx]
	for idx := maxIdx - 1; idx >= 0; idx-- {
		if dp[idx] == maxVal-1 && (nums[maxIdx]%nums[idx] == 0) {
			insertIdx -= 1
			maxVal = dp[idx]
			maxIdx = idx
			dp[insertIdx] = nums[idx]
		}
	}
	return dp[insertIdx:]
}

func TestSubset(t *testing.T) {
	assert.Equal(t, []int{9, 18, 90, 180, 360, 720}, largestDivisibleSubset([]int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720}))
}
