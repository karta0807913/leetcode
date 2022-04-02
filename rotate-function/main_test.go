package main

func maxRotateFunction(nums []int) int {
	sum := 0
	fSum := 0
	for idx := 1; idx < len(nums); idx++ {
		sum += nums[idx]
		fSum += nums[idx] * idx
	}
	maxVal := fSum
	for idx := len(nums) - 1; idx > 0; idx-- {
		sum += nums[(idx+1)%len(nums)]
		fSum += sum
		sum -= nums[idx]
		fSum -= len(nums) * nums[idx]
		if fSum > maxVal {
			maxVal = fSum
		}
	}
	return maxVal
}
