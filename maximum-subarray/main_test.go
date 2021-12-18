package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxSubArray(nums []int) (maxSub int) {
	maxSub = nums[0]
	currentMax := nums[0]
	for _, arrayValue := range nums[1:] {
		currentMax += arrayValue
		if currentMax < arrayValue {
			currentMax = arrayValue
		}
		if maxSub < currentMax {
			maxSub = currentMax
		}
	}
	return
}

func max(result int, others ...int) int {
	for _, value := range others {
		if value > result {
			result = value
		}
	}
	return result
}

func maxSubArraySlow(nums []int) (maxSub int) {
	maxSub = math.MinInt64
	devideStorage := make([]int, len(nums))
	for arrayIdx, arrayValue := range nums {
		for idx := range devideStorage[:arrayIdx+1] {
			devideStorage[idx] += arrayValue
		}
		maxSub = max(maxSub, devideStorage[:arrayIdx+1]...)
	}
	return
}

func TestResult(t *testing.T) {
	assert.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	assert.Equal(t, 23, maxSubArray([]int{5, 4, -1, 7, 8}))
	assert.Equal(t, 1, maxSubArray([]int{1}))
	assert.Equal(t, -1, maxSubArray([]int{-1}))
}
