package main

func findMin(nums []int) int {
	left, right, lastVal := 0, len(nums)-1, nums[len(nums)-1]
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < lastVal {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left == len(nums) {
		return nums[right]
	}
	if right == -1 {
		return nums[left]
	}
	if nums[left] < nums[right] {
		return nums[left]
	}
	return nums[right]
}
