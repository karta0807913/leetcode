package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func countSubarrays(nums []int, minK int, maxK int) int64 {
	minStart, maxStart := 0, 0
	satisfyMin, satisfyMax := false, false
	result := int64(0)
	start := 0
	for i, num := range nums {
		if num < minK || maxK < num {
			satisfyMax = false
			satisfyMin = false
			start = i + 1
			continue
		}
		if num == minK {
			minStart = i
			satisfyMin = true
		}
		if num == maxK {
			maxStart = i
			satisfyMax = true
		}
		if satisfyMax && satisfyMin {
			result += int64(min(minStart, maxStart) - start + 1)
		}
	}
	return result
}

func TestArr(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		int64(81),
		countSubarrays([]int{
			1, 2, 4, 4, 2, 4, 1, 4, 2, 4, 1, 2, 2, 2, 2,
		},
			1,
			4,
		),
	)
}
