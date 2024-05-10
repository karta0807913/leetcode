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

func continuousSubarrays(nums []int) int64 {
	var maxStack, minStack []int
	var ans int64
	// fmt.Printf("nums: %v\n", nums)
	var left int
	for i := range nums {
		for len(maxStack) != 0 {
			lastIndex := maxStack[len(maxStack)-1]
			if nums[lastIndex] >= nums[i] {
				break
			}
			maxStack = maxStack[:len(maxStack)-1]
		}
		maxStack = append(maxStack, i)
		for len(minStack) != 0 {
			lastIndex := minStack[len(minStack)-1]
			if nums[lastIndex] <= nums[i] {
				break
			}
			minStack = minStack[:len(minStack)-1]
		}
		minStack = append(minStack, i)
		// fmt.Printf("minStack: %v, \tmaxStack: %v\n", minStack, maxStack)
		for nums[maxStack[0]]-nums[minStack[0]] > 2 {
			if maxStack[0] < minStack[0] {
				maxStack = maxStack[1:]
			} else if maxStack[0] > minStack[0] {
				minStack = minStack[1:]
				// } else {
				// 	maxStack = maxStack[1:]
				// 	minStack = minStack[1:]
			}
			left = min(minStack[0], maxStack[0])
		}
		// fmt.Printf("minStack: %v, \tmaxStack: %v\n", minStack, maxStack)
		ans += int64(i-left) + 1
		// fmt.Printf("ans: %v, i: %v\n", ans, i)
	}
	return ans
}

func TestSubarr(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		int64(10),
		continuousSubarrays([]int{31, 30, 31, 32}),
	)
	assert.Equal(
		int64(8),
		continuousSubarrays([]int{5, 4, 2, 4}),
	)
}
