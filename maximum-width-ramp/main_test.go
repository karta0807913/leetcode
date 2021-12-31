package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxWidthRamp(nums []int) int {
	stack := make([]int, 1, len(nums))
	stack[0] = 0
	for idx, num := range nums {
		if nums[stack[len(stack)-1]] > num {
			stack = append(stack, idx)
		}
	}
	// fmt.Printf("stack: %v\n", stack)
	answer := 0
	for idx := len(nums) - 1; idx >= 0; idx-- {
		for len(stack) != 0 && nums[stack[len(stack)-1]] <= nums[idx] {
			tmp := idx - stack[len(stack)-1]
			if answer < tmp {
				answer = tmp
			}
			stack = stack[:len(stack)-1]
		}
	}
	return answer
}

func TestRamp(t *testing.T) {
	assert.Equal(t, 4, maxWidthRamp([]int{6, 0, 8, 2, 1, 5}))
	assert.Equal(t, 4, maxWidthRamp([]int{0, 8, 2, 1, 5}))
	assert.Equal(t, 7, maxWidthRamp([]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}))
}
