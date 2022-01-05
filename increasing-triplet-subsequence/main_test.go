package main

import "math"

func increasingTriplet(nums []int) bool {
	a, b := math.MaxInt64, math.MaxInt64
	for _, val := range nums {
		if val < a {
			a, b = val, a
		} else if val < b {
			b = val
		} else {
			return true
		}
	}
	return false
}
