package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findDuplicates(nums []int) (ans []int) {
	for _, n := range nums {
		n = abs(n)
		if nums[n-1] < 0 {
			ans = append(ans, n)
		} else {
			nums[n-1] *= -1
		}
	}
	return ans
}

func TestDuplicates(t *testing.T) {
	assert.Equal(t, []int{
		2,
	}, findDuplicates([]int{
		2, 2,
	}))
	assert.Equal(t, []int{
		1,
	}, findDuplicates([]int{
		1, 1, 2,
	}))
	assert.Equal(t, []int{
		2, 3,
	}, findDuplicates([]int{
		4, 3, 2, 7, 8, 2, 3, 1,
	}))
}
