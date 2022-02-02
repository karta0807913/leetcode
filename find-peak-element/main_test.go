package main

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	// fmt.Println(nums)
	mid := (len(nums)-1)/2 + 1
	if nums[mid-1] < nums[mid] {
		return mid + findPeakElement(nums[mid:])
	}
	return findPeakElement(nums[:mid])
}
