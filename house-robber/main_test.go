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
	prev, pprev, ppprev := 0, 0, 0
	for idx := 0; idx < len(nums); idx++ {
		prev, pprev, ppprev = max(pprev+nums[idx], ppprev+nums[idx]), prev, pprev
	}
	return max(pprev, prev)
}

func TestRob(t *testing.T) {
	assert.Equal(t, 4, rob([]int{1, 2, 3, 1}))
	assert.Equal(t, 101, rob([]int{1, 0, 0, 100}))
	assert.Equal(t, 12, rob([]int{2, 7, 9, 3, 1}))
	assert.Equal(t, 16, rob([]int{1, 2, 3, 0, 1, 2, 7, 9, 3, 1}))
}
