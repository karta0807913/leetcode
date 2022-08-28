package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func fillCups(amount []int) int {
	seconds := 1
	for ; ; seconds++ {
		filled := 0
		sort.Slice(amount, func(i, j int) bool {
			return amount[j] < amount[i]
		})
		for i := range amount {
			if amount[i] != 0 && filled != 2 {
				amount[i]--
				filled += 1
			}
		}
		if filled == 0 {
			break
		}
	}
	return seconds - 1
}

func TestFill(t *testing.T) {
	assert := require.New(t)
	assert.Equal(7, fillCups([]int{
		5, 4, 4,
	}))
	assert.Equal(4, fillCups([]int{
		1, 4, 2,
	}))
}
