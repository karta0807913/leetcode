package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] > envelopes[j][0] {
			return false
		} else {
			if envelopes[i][1] > envelopes[j][1] {
				return true
			}
		}
		return false
	})
	top := []int{}
	for _, env := range envelopes {
		h := env[1]
		left := 0
		right := len(top)
		for left < right {
			mid := (left + right) / 2
			if top[mid] < h {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if right == len(top) {
			top = append(top, h)
		} else {
			top[right] = env[1]
		}
	}
	return len(top)
}

func TestEnvelopes(t *testing.T) {
	assert.Equal(t, 2, maxEnvelopes([][]int{
		{1, 8}, {2, 3}, {5, 4}, {5, 2}, {6, 7}, {6, 4},
	}))
}
