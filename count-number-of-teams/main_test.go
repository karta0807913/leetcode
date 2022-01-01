package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func numTeams(rating []int) (answer int) {
	bigger := make([]int, len(rating))
	for target := len(rating) - 1; target >= 0; target-- {
		for current := target + 1; current < len(rating); current++ {
			if rating[current] > rating[target] {
				bigger[target] += 1
			}
		}
	}
	for left := 0; left < len(bigger); left++ {
		for right := left + 1; right < len(bigger); right++ {
			rightVal := len(rating) - 1 - right - bigger[right]
			if rating[left] > rating[right] {
				answer += rightVal
			} else {
				answer += bigger[right]
			}
		}
	}
	return
}

func TestTeams(t *testing.T) {
	assert.Equal(t, 3, numTeams([]int{2, 5, 3, 4, 1}))
	assert.Equal(t, 0, numTeams([]int{2, 1, 3}))
	assert.Equal(t, 4, numTeams([]int{1, 2, 3, 4}))
	assert.Equal(t, 9, numTeams([]int{74, 739, 867, 704, 16, 910, 415}))
}
