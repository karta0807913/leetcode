package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func BinarySearch(nums []int, distant int) (splitPoint int) {
	leftPtr, rightPtr := 0, len(nums)-1
	for leftPtr <= rightPtr {
		if nums[leftPtr] > distant {
			leftPtr++
		} else {
			nums[leftPtr], nums[rightPtr] =
				nums[rightPtr], nums[leftPtr]
			rightPtr--
		}
	}
	return leftPtr
}

func findKthLargest(nums []int, k int) int {
	min, max := -10000, 10000
	for min < max && k != 0 {
		mid := (max-min)/2 + min
		p := BinarySearch(nums, mid)
		// fmt.Printf("nums: %
		// v, distant: %v, p: %v, k: %v\n", nums, mid, p, k)
		if p < k {
			k -= p
			nums = nums[p:]
			max = mid - 1
		} else {
			nums = nums[:p]
			min = mid + 1
		}
		// fmt.Printf("nums: %v, min: %v, max: %v\n", nums, min, max)
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return nums[k-1]
}

func TestLargest(t *testing.T) {
	assert := require.New(t)
	assert.Equal(BinarySearch([]int{
		6, 5, 5, 4, 2, 1, 3, 3, 2,
	}, 4), 3)
	assert.Equal(5, findKthLargest([]int{
		3, 2, 1, 5, 6, 4,
	}, 2))
	assert.Equal(4, findKthLargest([]int{
		3, 2, 3, 1, 2, 4, 5, 5, 6,
	}, 4))
}
