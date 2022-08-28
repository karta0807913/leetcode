package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func numberOfArithmeticSlices(nums []int) int {
	ans := 0
	gap := -999999999
	combo := 0
	for idx := 1; idx < len(nums); idx++ {
		prev := nums[idx-1]
		if nums[idx]-prev != gap {
			combo -= 2
			if combo > 0 {
				combo = ((combo + 1) * combo) / 2
				ans += combo
			}
			combo = 1
		}
		combo += 1
		gap = nums[idx] - prev
	}
	combo -= 2
	if combo > 0 {
		combo = ((combo + 1) * combo) / 2
		ans += combo
	}
	return ans
}

func TestSlices(t *testing.T) {
	assert := require.New(t)

	assert.Equal(3, numberOfArithmeticSlices([]int{1, 2, 3, 4}))
}
