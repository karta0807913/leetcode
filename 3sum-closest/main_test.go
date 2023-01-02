package main

import (
	"math"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closet := math.MaxInt
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		remains := target - nums[i]
		for left < right {
			tmp := nums[left] + nums[right]
			if tmp < remains {
				left += 1
			} else if tmp > remains {
				right -= 1
			} else {
				return target
			}
			if abs(remains-tmp) < abs(target-closet) {
				closet = tmp + nums[i]
			}
		}
	}
	return closet
}
