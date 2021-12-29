package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	prev, pprev, ppprev := 0, 0, 0
	for idx := 1; idx < len(nums); idx++ {
		prev, pprev, ppprev = max(pprev+nums[idx], ppprev+nums[idx]), prev, pprev
	}
	round1 := max(pprev, prev)

	// fmt.Printf("round1: %v\n", round1)
	prev, pprev, ppprev = 0, 0, 0
	for idx := 0; idx < len(nums)-1; idx++ {
		prev, pprev, ppprev = max(pprev+nums[idx], ppprev+nums[idx]), prev, pprev
		// fmt.Printf("prev: %v, pprev: %v, ppprev: %v\n", prev, pprev, ppprev)
	}
	// fmt.Printf("prev: %v, pprev: %v\n", prev, pprev)
	return max(round1, max(prev, pprev))
}

func TestRob(t *testing.T) {
	assert.Equal(t, 1, rob([]int{1}))
	assert.Equal(t, 4, rob([]int{1, 2, 3, 1}))
	assert.Equal(t, 100, rob([]int{1, 0, 0, 100}))
	assert.Equal(t, 11, rob([]int{2, 7, 9, 3, 1}))
	assert.Equal(t, 3, rob([]int{2, 3, 2}))
	assert.Equal(t, 15, rob([]int{1, 2, 3, 0, 1, 2, 7, 9, 3, 1}))
	assert.Equal(t, 340, rob([]int{200, 3, 140, 20, 10}))
	assert.Equal(t, 340, rob([]int{200, 1, 3, 140, 20, 10}))
}
