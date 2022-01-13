package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	n := len(nums) / 2
	for idx := 0; idx < n; idx++ {
		r := len(nums) - idx - 1
		nums[idx], nums[r] = nums[r], nums[idx]
	}
}

func TestRoute(t *testing.T) {
	target := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(target, 1)
	assert.Equal(t, []int{7, 1, 2, 3, 4, 5, 6}, target)

	target = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(target, 2)
	assert.Equal(t, []int{6, 7, 1, 2, 3, 4, 5}, target)
}
