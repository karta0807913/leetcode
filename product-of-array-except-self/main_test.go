package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	product := nums[0]
	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = product
		product *= nums[i]
	}

	product = nums[len(nums)-1]
	for i := len(nums) - 2; i > 0; i-- {
		result[i] *= product
		product *= nums[i]
	}
	result[0] = product
	return result
}

func TestProduct(t *testing.T) {
	assert := require.New(t)
	assert.Equal([]int{0, 0, 9, 0, 0}, productExceptSelf([]int{-1, 1, 0, -3, 3}))
	assert.Equal([]int{24, 12, 8, 6}, productExceptSelf([]int{1, 2, 3, 4}))
	assert.Equal([]int{0, 0}, productExceptSelf([]int{0, 0}))
	assert.Equal([]int{0, 1}, productExceptSelf([]int{1, 0}))
	assert.Equal([]int{1, 0}, productExceptSelf([]int{0, 1}))
	assert.Equal([]int{0, -1}, productExceptSelf([]int{-1, 0}))
	assert.Equal([]int{-1, 0}, productExceptSelf([]int{0, -1}))
}
