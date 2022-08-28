package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func maxResult(nums []int, k int) int {
	monoStack := make([]int, 0, len(nums))
	monoStack = append(monoStack, 0)
	for i := 1; i < len(nums); i++ {
		if monoStack[0] < i-k {
			monoStack = monoStack[1:]
		}
		nums[i] = nums[i] + nums[monoStack[0]]
		j := len(monoStack) - 1
		for ; j >= 0 && nums[monoStack[j]] <= nums[i]; j-- {
		}
		monoStack = monoStack[:j+1]
		monoStack = append(monoStack, i)
		// fmt.Printf("monoStack: %v\n", monoStack)
	}
	// fmt.Printf("nums: %v\n", nums)
	return nums[len(nums)-1]
}

func TestResult(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		17,
		maxResult([]int{
			10, -5, -2, 4, 0, 3,
		}, 3),
	)

	assert.Equal(
		-3,
		maxResult([]int{
			1, -1, -2, 4, -7,
		}, 2),
	)

	assert.Equal(
		7,
		maxResult([]int{
			1, -1, -2, 4, -7, 3,
		}, 2),
	)

}
