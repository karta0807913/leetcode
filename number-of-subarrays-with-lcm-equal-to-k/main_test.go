package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func subarrayLCM(nums []int, k int) int {
	cache := make([]int, 0, len(nums))
	result := 0
	for _, num := range nums {
		for i, c := range cache {
			if c > k {
				continue
			}
			gcd := GCD(c, num)
			cache[i] = (c / gcd) * (num / gcd) * gcd
			if cache[i] == k {
				result++
			}
		}
		if num == k {
			result++
		}
		cache = append(cache, num)
	}
	return result
}

func TestLCM(t *testing.T) {
	assert := require.New(t)
	assert.Equal(4, subarrayLCM([]int{3, 6, 2, 7, 1}, 6))
}
