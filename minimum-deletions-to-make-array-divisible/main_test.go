package main

import "sort"

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func minOperations(nums []int, numsDivide []int) int {
	gcd := numsDivide[0]
	for _, n := range numsDivide {
		gcd = GCD(gcd, n)
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i, val := range nums {
		if gcd%val == 0 {
			return i
		}
		if val > gcd {
			break
		}
	}
	return -1
}
