package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func iterate(numbers map[int]int) (ans [][]int) {
	for key := range numbers {
		if numbers[key] == 0 {
			continue
		}
		numbers[key] -= 1
		result := iterate(numbers)
		numbers[key] += 1
		if result == nil {
			ans = append(ans, []int{key})
			continue
		}
		for _, res := range result {
			ans = append(ans, append(res, key))
		}
	}
	return
}

func permuteUnique(nums []int) [][]int {
	numbers := make(map[int]int)
	for idx := range nums {
		numbers[nums[idx]] += 1
	}
	return iterate(numbers)
}

func TestUnique(t *testing.T) {
	assert.ElementsMatch(t, [][]int{{1, 2}, {2, 1}}, permuteUnique([]int{1, 2}))
	assert.ElementsMatch(t, [][]int{
		{1, 1, 2}, {1, 2, 1}, {2, 1, 1},
	}, permuteUnique([]int{1, 1, 2}))
}
