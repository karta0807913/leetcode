package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func subarraySum(nums []int, k int) int {
	result := 0
	hashMap := make(map[int]int, len(nums))
	hashMap[0] = 1
	total := 0
	for _, val := range nums {
		total += val
		result += hashMap[total-k]
		hashMap[total] += 1
	}
	return result
}

func TestSum(t *testing.T) {
	assert := require.New(t)
	assert.Equal(1, subarraySum([]int{1, 1, 2, -2}, 0))
	assert.Equal(0, subarraySum([]int{1}, 0))
	assert.Equal(1, subarraySum([]int{1, 2, 1}, 2))
}
