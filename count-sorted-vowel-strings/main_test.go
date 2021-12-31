package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func countVowelStrings(n int) int {
	answer := []int{5, 4, 3, 2, 1}
	for i := 1; i < n; i++ {
		remain := sum(answer)
		sub := 0
		for idx := range answer {
			tmp := remain - sub
			sub = answer[idx]
			answer[idx] = tmp
			remain = answer[idx]
		}
	}
	return answer[0]
}

func TestCount(t *testing.T) {
	assert.Equal(t, 5, countVowelStrings(1))
	assert.Equal(t, 15, countVowelStrings(2))
	assert.Equal(t, 35, countVowelStrings(3))
	assert.Equal(t, 70, countVowelStrings(4))
}
