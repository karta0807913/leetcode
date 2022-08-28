package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	prev := nums[0] - 1
	for i, target := range nums {
		if target == prev {
			continue
		}
		tmp := make(map[int]int)
		for _, num := range nums[i+1:] {
			if tmp[num] > 1 {
				continue
			}
			n := -target - num
			if tmp[n] != 0 {
				tmp[num] += 1
				result = append(result, []int{
					target,
					n,
					num,
				})
			}
			tmp[num] += 1
		}
		prev = target
	}
	return result
}

func Test3Sum(t *testing.T) {
	assert := require.New(t)
	assert.ElementsMatch(
		[][]int{
			{-2, 1, 1},
			{-1, -1, 2},
		},
		threeSum([]int{2, -1, -1, 2, 2, 1, 1, -2}),
	)
	assert.ElementsMatch(
		[][]int{{-1, -1, 2}, {-1, 0, 1}},
		threeSum([]int{-1, 0, 1, 2, -1, -4}),
	)
}
