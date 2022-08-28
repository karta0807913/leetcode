package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func checkPossibility(nums []int) bool {
	prev := math.MinInt
	modified := false
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if modified {
				return false
			}
			modified = true
			if prev <= nums[i+1] {
				nums[i] = prev
			} else {
				nums[i+1] = nums[i]
			}
		}
		prev = nums[i]
	}
	return true
}

func TestCheck(t *testing.T) {
	assert := require.New(t)
	assert.True(checkPossibility([]int{
		1, 4, 1, 2,
	}))
	assert.False(checkPossibility([]int{
		3, 4, 2, 3,
	}))
	assert.True(checkPossibility([]int{
		5, 7, 1, 8,
	}))
	assert.True(checkPossibility([]int{
		1, 2, 3, 4, 5, 6, 7,
	}))
	assert.True(checkPossibility([]int{
		1, 2, 1, 4, 5, 6, 7,
	}))
	assert.False(checkPossibility([]int{
		1, 2, 1, 4, 1, 6, 7,
	}))
	assert.True(checkPossibility([]int{
		4, 2, 3,
	}))
	assert.False(checkPossibility([]int{
		3, 2, 1,
	}))
}
