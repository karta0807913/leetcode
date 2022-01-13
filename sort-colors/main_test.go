package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sortColors(nums []int) {
	cache := []int{0, 0, 0}
	for _, val := range nums {
		cache[val] += 1
	}
	for idx := 0; idx < cache[0]; idx++ {
		nums[idx] = 0
	}
	for idx := cache[0]; idx < cache[0]+cache[1]; idx++ {
		nums[idx] = 1
	}
	for idx := cache[0] + cache[1]; idx < len(nums); idx++ {
		nums[idx] = 2
	}
}

// usual version, using heap sort
// func sortColors(nums []int) {
// 	for idx := len(nums)/2 - 1; idx >= 0; idx-- {
// 		down(nums, idx)
// 	}
// 	n := len(nums) - 1
// 	for n != 0 {
// 		nums[n], nums[0] = nums[0], nums[n]
// 		down(nums[:n], 0)
// 		n = n - 1
// 	}
// }

func down(nums []int, i int) {
	n := len(nums)
	for j := i*2 + 1; j < n && j >= 0; j = i*2 + 1 {
		j1 := j + 1
		if j1 >= 0 && j1 != n && nums[j1] > nums[j] {
			j = j1
		}
		if nums[i] >= nums[j] {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i = j
	}
}

func TestSort(t *testing.T) {
	m := []int{5, 4, 3}
	down(m, 0)
	fmt.Printf("m: %v\n", m)
	var target []int
	// target = []int{5, 4, 3, 2, 1}
	// sortColors(target)
	// assert.Equal(t, []int{1, 2, 3, 4, 5}, target)

	target = []int{1, 2, 2, 2, 2, 0, 0, 0, 1, 1}
	sortColors(target)
	assert.Equal(t, []int{0, 0, 0, 1, 1, 1, 2, 2, 2, 2}, target)
}

//heap sort
