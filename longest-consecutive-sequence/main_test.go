package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, val := range nums {
		set[val] = true
	}
	max := 0
	for _, val := range nums {
		if set[val-1] {
			continue
		}
		i := 0
		for set[val] {
			i++
			val += 1
		}
		if i > max {
			max = i
		}
	}
	return max
}

func TestFunc(t *testing.T) {
	assert.Equal(t, 4, longestConsecutive([]int{
		100, 4, 200, 1, 3, 2,
	}))
	assert.Equal(t, 2, longestConsecutive([]int{
		0, 0, 0, 0, 1, 0, 0, 0, 1,
	}))
	assert.Equal(t, 9, longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
