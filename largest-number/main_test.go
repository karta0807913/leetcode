package main

import (
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func grater(left, right string) bool {
	return left+right > right+left
}

func largestNumber(nums []int) (result string) {
	strNums := make([]string, len(nums))
	for idx := range strNums {
		strNums[idx] = strconv.Itoa(nums[idx])
	}
	sort.Slice(strNums, func(i, j int) bool {
		return grater(strNums[i], strNums[j])
	})
	if strNums[0] == "0" {
		return "0"
	}
	for _, val := range strNums {
		result += val
	}
	return result
}

func TestGreater(t *testing.T) {
	assert.True(t, grater("34", "30"))
	assert.True(t, grater("34", "3"))
	assert.False(t, grater("3", "34"))
	assert.True(t, grater("3", "30"))
	assert.False(t, grater("30", "3"))
}
