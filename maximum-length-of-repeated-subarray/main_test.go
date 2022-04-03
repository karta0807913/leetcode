package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLength(nums1 []int, nums2 []int) int {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	dp := make([]int, len(nums2)+1)
	next := make([]int, len(nums2)+1)
	result := 0
	for _, num := range nums1 {
		for idx := range nums2 {
			if nums2[idx] == num {
				next[idx+1] = dp[idx] + 1
			}
			dp[idx] = 0
			result = max(next[idx+1], result)
		}
		dp[len(nums2)] = 0
		next, dp = dp, next
		// fmt.Printf("dp: %v\n", dp)
	}

	return result
}

func TestFindLength(t *testing.T) {
	assert := require.New(t)
	assert.Equal(3, findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
	assert.Equal(4, findLength([]int{1, 1, 2, 3, 4}, []int{1, 2, 3, 4}))
	assert.Equal(4, findLength([]int{0, 0, 0, 0}, []int{0, 0, 0, 0}))
	assert.Equal(4, findLength([]int{0, 0, 0, 0}, []int{0, 0, 0, 0}))
	assert.Equal(2, findLength([]int{0, 1, 1, 1, 1}, []int{1, 0, 1, 0, 1}))
	assert.Equal(3, findLength([]int{1, 0, 0, 0, 1}, []int{1, 0, 0, 1, 1}))
}
