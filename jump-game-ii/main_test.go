package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func min(nums []int) int {
	m := -1
	for _, n := range nums {
		if n != -1 && (n < m || m == -1) {
			m = n
		}
	}
	return m
}

func MinInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var steps []int = make([]int, len(nums))
	for index := len(nums) - 2; index >= 0; index-- {
		left := index + 1
		right := MinInt(index+nums[index]+1, len(steps))
		if left > right {
			steps[index] = -1
			continue
		}
		m := min(steps[left:right])
		if m == -1 {
			steps[index] = -1
		} else {
			steps[index] = m + 1
		}
	}
	return steps[0]
}

func TestJump(t *testing.T) {
	assert.Equal(t, 2, jump([]int{2, 3, 1, 1, 4}), nil)
	assert.Equal(t, 2, jump([]int{2, 3, 0, 1, 4}), nil)
	assert.Equal(t, 1, jump([]int{2, 1}), nil)
}
