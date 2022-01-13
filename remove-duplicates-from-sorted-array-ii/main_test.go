package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeDuplicates(nums []int) int {
	left := 0
	k := 0
	prev := nums[0]
	for idx := range nums {
		if prev == nums[idx] {
			k++
		} else {
			k = 1
			prev = nums[idx]
		}
		nums[left] = nums[idx]
		if k < 3 {
			left++
		}
	}
	return left
}

func TestRemove(t *testing.T) {
	var target []int
	var k int
	target = []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	k = removeDuplicates(target)
	assert.Equal(t, []int{1, 1, 2, 2, 3, 3}, target[:k])

	target = []int{1, 1, 1, 2, 2, 2, 3, 3}
	k = removeDuplicates(target)
	assert.Equal(t, []int{1, 1, 2, 2, 3, 3}, target[:k])

	target = []int{1}
	k = removeDuplicates(target)
	assert.Equal(t, []int{1}, target[:k])
}
