package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lengthOfLIS(nums []int) int {
	cache := make([]int, len(nums))
	cache[0] = 1
	for idx, val := range nums {
		maxLength := 0
		for i := 0; i < idx; i++ {
			if val > nums[i] && cache[i] > maxLength {
				maxLength = cache[i]
			}
		}
		cache[idx] = maxLength + 1
	}
	max := 0
	for _, val := range cache {
		if val > max {
			max = val
		}
	}
	return max
}

func TestLength(t *testing.T) {
	assert.Equal(t, 4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	assert.Equal(t, 4, lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	assert.Equal(t, 1, lengthOfLIS([]int{7}))
	assert.Equal(t, 1, lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7}))
	assert.Equal(t, 6, lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}
