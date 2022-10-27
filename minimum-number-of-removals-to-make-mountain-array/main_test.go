package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func binarySearch(lis []int, n int) int {
	left, right := 0, len(lis)
	for left < right {
		mid := (left + right) / 2
		if lis[mid] < n {
			left = mid + 1
		} else if lis[mid] > n {
			right = mid
		} else {
			return mid
		}
	}
	return left
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumMountainRemovals(nums []int) int {
	dp := make([]int, len(nums))
	lis := make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		n := binarySearch(lis, nums[i])
		if n == len(lis) {
			lis = append(lis, nums[i])
		} else {
			lis[n] = min(nums[i], lis[n])
		}
		dp[i] = (len(nums) - 1 - i) - n
		if dp[i] == len(nums)-1-i {
			dp[i] = len(nums)
		}
	}
	result := len(nums)
	lis = lis[:0]
	lis = append(lis, nums[0])
	for i := 1; i < len(nums); i++ {
		n := binarySearch(lis, nums[i])
		if n == len(lis) {
			lis = append(lis, nums[i])
			result = min(result, dp[i]+(i-n))
		} else {
			lis[n] = min(nums[i], lis[n])
		}
	}
	return result
}

func TestResult(t *testing.T) {
	assert := require.New(t)
	assert.Equal(10, minimumMountainRemovals([]int{4, 5, 13, 17, 1, 7, 6, 11, 2, 8, 10, 15, 3, 9, 12, 14, 16}))
	assert.Equal(6, minimumMountainRemovals([]int{100, 92, 89, 77, 74, 66, 64, 66, 64}))
	assert.Equal(3, minimumMountainRemovals([]int{2, 1, 1, 5, 6, 2, 3, 1}))
	assert.Equal(2, minimumMountainRemovals([]int{9, 8, 1, 7, 6, 5, 4, 3, 2, 1}))
}
