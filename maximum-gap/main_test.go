package main

import "math"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	bucketMax := make([]int, len(nums))
	bucketMin := make([]int, len(nums))
	for i := range bucketMax {
		bucketMax[i] = math.MinInt64
		bucketMin[i] = math.MaxInt64
	}
	minVal := nums[0]
	maxVal := nums[0]
	for _, n := range nums {
		minVal = min(minVal, n)
		maxVal = max(maxVal, n)
	}
	if minVal == maxVal {
		return 0
	}

	gap := float64(maxVal-minVal) / float64(len(nums)-1)
	for _, n := range nums {
		idx := int(float64(n-minVal) / gap)
		bucketMax[idx] = max(bucketMax[idx], n)
		bucketMin[idx] = min(bucketMin[idx], n)
	}

	// fmt.Println(bucketMax, bucketMin)
	maxGap := 0
	previous := minVal
	for idx := range bucketMax {
		if bucketMax[idx] == math.MinInt64 || bucketMin[idx] == math.MaxInt64 {
			continue
		}
		maxGap = max(bucketMin[idx]-previous, maxGap)
		previous = bucketMax[idx]
	}
	return maxGap
}
