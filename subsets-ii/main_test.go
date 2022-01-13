package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func idx2subset(nums, current []int) []int {
	subset := []int{}
	for _, idx := range current {
		subset = append(subset, nums[idx])
	}
	return subset
}
func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	current := make([]int, 1, len(nums))
	current[0] = -1

	flag := false
	for len(current) > 0 {
		n := len(current) - 1
		current[n] += 1
		if !flag && current[n] != 0 && current[n] < len(nums) && nums[current[n]] == nums[current[n]-1] {
			continue
		}
		flag = false
		if current[n] >= len(nums) {
			current = current[:n]
			ans = append(ans, idx2subset(nums, current))
		} else if len(current) != len(nums) {
			current = append(current, current[n])
			flag = true
		} else {
			ans = append(ans, idx2subset(nums, current))
		}
	}
	return
}

func TestSubset(t *testing.T) {
	assert.ElementsMatch(t, [][]int{
		{},
		{0},
	}, subsetsWithDup([]int{0}))
	assert.ElementsMatch(t, [][]int{
		{},
		{1, 2, 2},
		{1, 2},
		{1},
		{2},
		{2, 2},
	}, subsetsWithDup([]int{1, 2, 2}))
}
