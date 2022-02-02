package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func nextGreaterElements(nums []int) []int {
	ans := make([]int, len(nums))
	stack := []int{}
	for idx := len(nums)*2 - 1; idx >= 0; idx-- {
		i := idx % len(nums)
		ans[i] = -1
		for len(stack) != 0 {
			n := len(stack) - 1
			if stack[n] > nums[i] {
				ans[i] = stack[n]
				break
			} else {
				stack = stack[:n]
			}
		}
		stack = append(stack, nums[i])
	}

	return ans
}

func TestElements(t *testing.T) {
	assert.Equal(t, []int{
		2, 3, -1,
	}, nextGreaterElements([]int{
		1, 2, 3,
	}))
	assert.Equal(t, []int{
		2, -1, 2,
	}, nextGreaterElements([]int{
		1, 2, 1,
	}))
}
