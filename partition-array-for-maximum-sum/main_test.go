package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxSumAfterPartitioning(arr []int, k int) int {
	cache := make([]int, k)
	for idx, val := range arr {
		maxNum := val
		maxVal := 0
		for i := 0; i < k && idx-i >= 0; i++ {
			if maxNum < arr[idx-i] {
				maxNum = arr[idx-i]
			}
			max := cache[k-i-1] + maxNum*(i+1)
			if maxVal < max {
				maxVal = max
			}
		}
		for i := 1; i < k; i++ {
			cache[i-1] = cache[i]
		}
		cache[k-1] = maxVal
	}
	return cache[k-1]
}

func TestPartitioning(t *testing.T) {
	assert.Equal(t, 84, maxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3))
	assert.Equal(t, 83, maxSumAfterPartitioning([]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4))
	//                                                1, 7, 7, 7, 7, 9, 9, 9, 9, 9, 9
}
