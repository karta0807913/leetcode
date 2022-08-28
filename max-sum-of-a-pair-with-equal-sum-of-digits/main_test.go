package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func maximumSum(nums []int) int {
	m := make(map[int][]int)
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] < nums[i]
	})
	for _, val := range nums {
		t := val
		c := 0
		for val != 0 {
			c += val % 10
			val /= 10
		}
		if len(m[c]) == 2 {
			continue
		}
		m[c] = append(m[c], t)
	}
	max := -1
	for _, arr := range m {
		if len(arr) < 2 {
			continue
		}
		t := arr[0] + arr[1]
		if t > max {
			max = t
		}
	}
	return max
}

func TestSum(t *testing.T) {
	assert := require.New(t)
	assert.Equal(54, maximumSum([]int{18, 43, 36, 13, 7}))
}
