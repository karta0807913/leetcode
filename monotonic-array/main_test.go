package monotonicarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isIncreaseMonotonic(nums []int) bool {
	before := nums[0]
	for idx := range nums[1:] {
		if before > nums[idx+1] {
			return false
		}
		before = nums[idx+1]
	}
	return true
}

func isDecreaseMonotonic(nums []int) bool {
	before := nums[0]
	for idx := range nums[1:] {
		if before < nums[idx+1] {
			return false
		}
		before = nums[idx+1]
	}
	return true
}

func isMonotonic(nums []int) bool {
	before := nums[0]
	for idx := range nums[1:] {
		if nums[idx+1] < before {
			return isDecreaseMonotonic(nums[idx+1:])
		} else if nums[idx+1] > before {
			return isIncreaseMonotonic(nums[idx+1:])
		}
	}
	return true
}

func TestMonotonic(t *testing.T) {
	assert.True(t, isMonotonic([]int{1, 2, 3, 4}))
	assert.True(t, isMonotonic([]int{1, 1, 1, 1}))
	assert.True(t, isMonotonic([]int{1, 1, 1, 3}))
	assert.False(t, isMonotonic([]int{1, 1, 4, 3}))
}
