package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func calRemoveTime(times []int) int {
	if len(times) == 1 {
		return 0
	}
	max := 0
	sum := 0
	for _, time := range times {
		sum += time
		if max < time {
			max = time
		}
	}
	return sum - max
}

func minCost(colors string, neededTime []int) int {
	beforeIdx := 0
	answer := 0
	sum := neededTime[0]
	maximum := neededTime[0]
	for idx := 1; idx < len(colors); idx++ {
		if colors[idx] != colors[beforeIdx] {
			answer += sum - maximum
			beforeIdx = idx
			sum = 0
			maximum = 0
		}
		if neededTime[idx] > maximum {
			maximum = neededTime[idx]
		}
		sum += neededTime[idx]
	}
	answer += sum - maximum
	return answer
}

func TestCost(t *testing.T) {
	assert.Equal(t, 3, minCost("abaac", []int{
		1, 2, 3, 4, 5,
	}))
	assert.Equal(t, 0, minCost("abc", []int{
		1, 2, 3,
	}))
	assert.Equal(t, 1, minCost("bbababa", []int{
		1, 99, 2, 3, 4, 5, 1,
	}))
	assert.Equal(t, 2, minCost("aabaa", []int{
		1, 2, 3, 4, 1,
	}))
	assert.Equal(t, 8, minCost("cddcdcae", []int{
		4, 8, 8, 4, 4, 5, 4, 2,
	}))
}
