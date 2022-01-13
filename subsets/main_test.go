package main

import (
	"fmt"
	"testing"
)

// using bit logic
func subsets(nums []int) (ans [][]int) {
	var bits uint

	max := uint(1)
	for range nums {
		max *= 2
	}

	for ; bits != max; bits++ {
		v := []int{}
		idx := 0
		for idx < len(nums) {
			if bits&(1<<idx) != 0 {
				v = append(v, nums[idx])
			}
			idx += 1
		}
		// fmt.Printf("bits: %v\n", bits)
		// fmt.Printf("v: %v\n", v)
		ans = append(ans, v)
	}
	return
}

// func idx2subset(nums, current []int) []int {
// 	subset := []int{}
// 	for _, idx := range current {
// 		subset = append(subset, nums[idx])
// 	}
// 	return subset
// }

// // the same logic with 77.Combinations
// func subsets(nums []int) (ans [][]int) {
// 	current := make([]int, 1, len(nums))
// 	current[0] = -1
// 	for len(current) > 0 {
// 		n := len(current) - 1
// 		current[n] += 1
// 		if current[n] == len(nums) {
// 			current = current[:n]
// 			ans = append(ans, idx2subset(nums, current))
// 			// fmt.Printf("current: %v\n", current)
// 		} else if len(current) != len(nums) {
// 			current = append(current, current[n])
// 		} else {
// 			ans = append(ans, idx2subset(nums, current))
// 			// fmt.Printf("current: %v\n", current)
// 		}
// 	}
// 	return ans
// }

func TestSubset(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
	t.Fail()
}
