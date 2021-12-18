package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func numPairsDivisibleBy60(time []int) int {
	bucket := make([]int, 60)
	for _, val := range time {
		bucket[val%60] += 1
	}
	result := 0
	result += (bucket[0]*(bucket[0]-1))/2 + (bucket[30]*(bucket[30]-1))/2
	for idx := 1; idx < 30; idx++ {
		result += bucket[idx] * bucket[60-idx]
	}
	return result
}

func slowDivisibleBy60(time []int) int {
	result := 0
	for idxA := 0; idxA < len(time); idxA++ {
		for idxB := idxA + 1; idxB < len(time); idxB++ {
			if (time[idxA]+time[idxB])%60 == 0 {
				result += 1
			}
		}
	}
	return result
}

func TestNumber(t *testing.T) {
	var time []int
	time = []int{30, 20, 150, 100, 40}
	assert.Equal(t, slowDivisibleBy60(time), numPairsDivisibleBy60(time), nil)

	time = []int{30, 20, 20, 150, 100, 40}
	assert.Equal(t, slowDivisibleBy60(time), numPairsDivisibleBy60(time), nil)

	// time = []int{60, 60, 60}
	// assert.Equal(t, slowDivisibleBy60(time), numPairsDivisibleBy60(time), nil)
}
