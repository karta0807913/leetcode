package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
	} else {
		if b > c {
			return b
		}
	}
	return c
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	answer := nums[0]
	maxVal, minVal := 0, 0

	for _, num := range nums {
		maxVal, minVal = maxVal*num, minVal*num
		maxVal, minVal = max(maxVal, num, minVal), min(maxVal, num, minVal)
		if maxVal > answer {
			answer = maxVal
		}
	}
	return answer
}

func TestProduct(t *testing.T) {
	assert.Equal(t, 6, maxProduct([]int{2, 3, -2, 4}))
	assert.Equal(t, 0, maxProduct([]int{-2, 0, -1}))
	assert.Equal(t, -2, maxProduct([]int{-2}))
	assert.Equal(t, 6, maxProduct([]int{-2, -3}))
	assert.Equal(t, 6, maxProduct([]int{-2, -3}))
	assert.Equal(t, 24, maxProduct([]int{-2, 2, 3, 4}))
	assert.Equal(t, 48, maxProduct([]int{-2, 2, 3, 4, -1}))
	assert.Equal(t, 24, maxProduct([]int{2, -5, -2, -4, 3}))

}
