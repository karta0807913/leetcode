package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func recursive(n int, start, k, totalK int) [][]int {
	if n > (((9 + (9 - k + 1)) * k) / 2) {
		return nil
	}
	if k == 0 {
		return nil
	}
	if k == 1 {
		if n >= start && n < 10 {
			result := make([]int, 0, totalK)
			result = append(result, n)
			return [][]int{result}
		}
		return nil
	}
	ans := [][]int{}
	for i := start; i < 10 && i < n; i++ {
		for _, val := range recursive(n-i, i+1, k-1, totalK) {
			ans = append(ans, append(val, i))
		}
	}
	return ans
}

func combinationSum3(k int, n int) [][]int {
	if n > (((9 + (9 - k + 1)) * k) / 2) {
		return nil
	}
	return recursive(n, 1, k, k)
}

func TestSum(t *testing.T) {
	assert := require.New(t)
	assert.ElementsMatch(
		[][]int{{4, 2, 1}},
		combinationSum3(3, 7),
	)
	assert.ElementsMatch(
		[][]int{{9, 8, 7, 6, 5, 4, 3, 2, 1}},
		combinationSum3(9, 45),
	)
	assert.ElementsMatch(
		[][]int{},
		combinationSum3(2, 18),
	)
}
