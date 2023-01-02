package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func goodIndices(nums []int, k int) []int {
	lr := make([]int, len(nums))
	lr[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			lr[i] = 1
			continue
		}
		lr[i] += lr[i-1] + 1
	}

	rl := make([]int, len(nums))
	rl[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			rl[i] = 1
			continue
		}
		rl[i] += rl[i+1] + 1
	}
	result := []int{}
	for i := k; i < len(nums)-k; i++ {
		if lr[i-1] >= k && rl[i+1] >= k {
			result = append(result, i)
		}
	}
	return result
}

func TestA(t *testing.T) {
	assert := require.New(t)
	assert.Equal([]int{2, 3}, goodIndices([]int{
		2, 1, 1, 1, 3, 4, 1,
	}, 2))
}
