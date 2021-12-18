package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findRotatePoint(nums []int) int {
	start, end := 0, len(nums)-1
	if len(nums) == 0 {
		return -1
	}
	for start+1 < end {
		// fmt.Println(start, end)
		tmp := 0
		if nums[0] > nums[end] {
			tmp = (end + start) / 2
		} else {
			tmp = (start + end) / 2
		}
		if nums[tmp] < nums[0] {
			end = tmp
		} else {
			start = tmp
		}
	}
	if end == len(nums)-1 {
		if nums[start] > nums[end] {
			return end
		}
		return -1
	}
	return end
}

func BinarySearch(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		idx := (start + end) / 2
		if nums[idx] < target {
			start = idx + 1
		} else if nums[idx] > target {
			end = idx - 1
		} else if nums[idx] == target {
			return idx
		}
	}
	return -1
}

func search(nums []int, target int) int {
	k := findRotatePoint(nums)
	if k == -1 {
		return BinarySearch(nums, target)
	}
	if nums[0] <= target {
		return BinarySearch(nums[:k], target)
	} else {
		p := BinarySearch(nums[k:], target)
		if p == -1 {
			return -1
		}
		return k + p
	}
}

func TestSearch(t *testing.T) {
	assert.Equal(t, 3, findRotatePoint([]int{5, 6, 7, 0, 1, 2, 4}), nil)
	assert.Equal(t, -1, findRotatePoint([]int{0, 1, 2, 4}), nil)
	assert.Equal(t, 4, findRotatePoint([]int{5, 6, 7, 8, 0, 1, 2, 4}), nil)
	assert.Equal(t, -1, findRotatePoint([]int{0, 1, 2, 4, 5, 6, 7}), nil)

	assert.Equal(t, 2, BinarySearch([]int{0, 1, 2, 4, 5, 6, 7}, 2), nil)
	assert.Equal(t, 6, BinarySearch([]int{0, 1, 2, 4, 5, 6, 7}, 7), nil)

	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0), nil)
	assert.Equal(t, -1, search([]int{4, 5, 6, 7, 0, 1, 2}, 3), nil)
	assert.Equal(t, -1, search([]int{1}, 0), nil)
	assert.Equal(t, -1, search([]int{}, 0), nil)
	assert.Equal(t, 1, search([]int{3, 1}, 1), nil)
	assert.Equal(t, -1, search([]int{3, 5, 1}, 0), nil)
	assert.Equal(t, -1, search([]int{3, 5, 1, 2}, 4), nil)
}
