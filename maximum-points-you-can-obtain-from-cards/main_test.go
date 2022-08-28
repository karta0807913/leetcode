package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func maxScore(cardPoints []int, k int) int {
	sum := 0
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
	}
	if k == len(cardPoints) {
		return sum
	}
	max := sum
	for i := k - 1; i >= 0; i-- {
		removed := cardPoints[i]
		added := cardPoints[len(cardPoints)-(k-i)]
		sum = sum - removed + added
		if sum > max {
			max = sum
		}
	}
	return max
}

func TestScore(t *testing.T) {
	assert := require.New(t)
	assert.Equal(202, maxScore([]int{
		1, 79, 80, 1, 1, 1, 200, 1,
	}, 3))
	assert.Equal(12, maxScore([]int{
		1, 2, 3, 4, 5, 6, 1,
	}, 3))
}
