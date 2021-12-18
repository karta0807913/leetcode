package findfirstandlastpositionofelementinsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func binarySearch(nums []int, target int) (int, int, int) {
	for left, right := 0, len(nums)-1; left <= right; {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid, left, right
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1, 0, 0
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if nums[0] == target && nums[len(nums)-1] == target {
		return []int{0, len(nums) - 1}
	}
	idx, left, right := binarySearch(nums, target)
	if idx == -1 {
		return []int{-1, -1}
	}
	leftTarget := searchRange(nums[left:idx], target)
	if leftTarget[0] == -1 {
		left = idx
	} else {
		left = leftTarget[0] + left
	}
	rightTarget := searchRange(nums[idx+1:right+1], target)
	if rightTarget[1] == -1 {
		right = idx
	} else {
		right = rightTarget[1] + idx + 1
	}
	return []int{left, right}
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
	idx, _, _ := binarySearch([]int{8, 10}, 8)
	assert.Equal(t, 0, idx)
}
