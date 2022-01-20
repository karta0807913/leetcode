package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func binarySearch(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			left = left + 1
		} else {
			right = right - 1
		}
	}
	return false
}

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		// fmt.Printf("nums[%v:%v]: %v\n", left, right, nums[left:right+1])
		if nums[mid] == target {
			return true
		}
		if nums[mid] < nums[right] {
			if binarySearch(nums[mid+1:right+1], target) {
				return true
			}
			right = mid - 1
			continue
		}
		if nums[left] < nums[mid] {
			if binarySearch(nums[left:mid], target) {
				return true
			}
			left = mid + 1
			continue
		}

		if nums[left] > nums[mid] {
			if binarySearch(nums[mid+1:right+1], target) {
				return true
			}
			right = mid - 1
			continue
		}

		if nums[mid] > nums[right] {
			if binarySearch(nums[left:mid], target) {
				return true
			}
			left = mid + 1
			continue
		}
		left++
	}
	return false
}

func TestSearch(t *testing.T) {
	assert.True(t, search([]int{1, 1, 1, 1, 1, 1}, 1))
	assert.True(t, search([]int{4, 5, 6, 7, 1, 2, 3, 4}, 1))
	assert.True(t, search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	assert.False(t, search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	assert.True(t, search([]int{1, 0, 1, 1, 1}, 0))
	assert.True(t, search([]int{1, 1, 1, 0, 1}, 0))
	assert.True(t, search([]int{
		1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	}, 13))
	// t.Fail()
}
