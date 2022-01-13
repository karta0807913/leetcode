package findfirstandlastpositionofelementinsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	ans[0] = FindFirst(nums, target)
	ans[1] = FindLast(nums, target)
	return ans
}

func FindLast(nums []int, target int) (ans int) {
	left, right := 0, len(nums)-1
	ans = -1
	for left <= right {
		i := (left + right) / 2
		if nums[i] == target {
			ans = i
		}
		if nums[i] <= target {
			left = i + 1
		} else {
			right = i - 1
		}
	}
	return ans
}

func FindFirst(nums []int, target int) (ans int) {
	ans = -1
	left, right := 0, len(nums)-1
	for left <= right {
		i := (left + right) / 2
		if nums[i] == target {
			ans = i
		}
		if target <= nums[i] {
			right = i - 1
		} else {
			left = i + 1
		}
	}
	return ans
}

func TestSearchRange(t *testing.T) {
	assert.Equal(t, []int{1, 2}, searchRange([]int{1, 2, 2, 3}, 2))
	assert.Equal(t, []int{1, 3}, searchRange([]int{1, 2, 2, 2}, 2))
	assert.Equal(t, []int{0, 2}, searchRange([]int{2, 2, 2, 1}, 2))
	assert.Equal(t, []int{2, 4}, searchRange([]int{1, 1, 2, 2, 2}, 2))
	assert.Equal(t, []int{0, 2}, searchRange([]int{2, 2, 2, 1, 1}, 2))
	assert.Equal(t, []int{2, 2}, searchRange([]int{1, 1, 2, 3}, 2))
	assert.Equal(t, []int{0, 2}, searchRange([]int{2, 2, 2, 3}, 2))
	assert.Equal(t, []int{0, 3}, searchRange([]int{2, 2, 2, 2}, 2))
	assert.Equal(t, []int{-1, -1}, searchRange([]int{2, 2, 2, 2}, 3))
	assert.Equal(t, []int{-1, -1}, searchRange([]int{}, 3))
	assert.Equal(t, []int{3, 4}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	assert.Equal(t, []int{1, 1}, searchRange([]int{1, 2}, 2))
}
