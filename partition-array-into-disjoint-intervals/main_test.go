package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func partitionDisjoint(nums []int) int {
	smallest := make([]int, len(nums))
	smallest[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i > 0; i-- {
		if nums[i] < smallest[i+1] {
			smallest[i] = nums[i]
		} else {
			smallest[i] = smallest[i+1]
		}
	}
	biggest := 0
	for idx := 0; idx < len(nums)-1; idx++ {
		if biggest < nums[idx] {
			biggest = nums[idx]
		}
		if biggest <= smallest[idx+1] {
			return idx + 1
		}
	}
	return -1
}

func TestDisjoint(t *testing.T) {
	assert.Equal(t, 3, partitionDisjoint([]int{5, 0, 3, 8, 6}))
	assert.Equal(t, 4, partitionDisjoint([]int{1, 1, 1, 0, 6, 12}))
	assert.Equal(t, 1, partitionDisjoint([]int{1, 1}))
}
