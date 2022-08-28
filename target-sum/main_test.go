package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func findTargetSumWays(nums []int, target int) int {
	dp := map[int]int{target: 1}
	var next map[int]int

	for _, num := range nums {
		next = make(map[int]int, len(dp)*2)
		for val, times := range dp {
			next[val+num] += times
			next[val-num] += times
		}
		dp, next = next, dp
	}
	return dp[0]
}

// // recursive solution
// func recursive(nums []int, target int, cache []map[int]int) int {
// 	if len(nums) == 0 {
// 		if target == 0 {
// 			return 1
// 		}
// 		return 0
// 	}
// 	if cache[len(nums)-1] == nil {
// 		cache[len(nums)-1] = make(map[int]int)
// 	}
// 	localCache := cache[len(nums)-1]
// 	if v, ok := localCache[target]; ok {
// 		return v
// 	}
// 	localCache[target] += recursive(nums[1:], target+nums[0], cache)
// 	localCache[target] += recursive(nums[1:], target-nums[0], cache)
// 	return localCache[target]
// }

// func findTargetSumWays(nums []int, target int) int {
// 	cache := make([]map[int]int, len(nums))
// 	return recursive(nums, target, cache)
// }

func TestTargetSum(t *testing.T) {
	assert := require.New(t)
	assert.Equal(5, findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	assert.Equal(64, findTargetSumWays([]int{0, 0, 0, 0, 0, 0}, 0))
	assert.Equal(6, findTargetSumWays([]int{5, 4, 3, 2, 1, 0}, 3))
	assert.Equal(31076, findTargetSumWays([]int{
		6, 4, 4, 1, 10, 7, 5, 0, 3, 5, 4, 8, 8, 3, 7, 10, 2, 8, 3, 6,
	}, 0))
}
