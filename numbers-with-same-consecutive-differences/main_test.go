package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func recursive(arr []int, n, k int) []int {
	newArr := []int{}
	for _, val := range arr {
		lastDigits := val % 10
		nextUppperDigits := lastDigits + k
		if nextUppperDigits < 10 {
			newArr = append(newArr, val*10+nextUppperDigits)
		}
		nextLowerDigits := lastDigits - k
		if nextUppperDigits != nextLowerDigits &&
			nextLowerDigits >= 0 {
			newArr = append(newArr, val*10+nextLowerDigits)
		}
	}
	if n == 1 {
		return newArr
	}
	return recursive(newArr, n-1, k)
}

func numsSameConsecDiff(n int, k int) (result []int) {
	for i := 1; i < 10; i++ {
		result = append(result, recursive([]int{i}, n-1, k)...)
	}
	return result
}

func TestDiff(t *testing.T) {
	assert := require.New(t)
	assert.ElementsMatch(
		[]int{181, 292, 707, 818, 929},
		numsSameConsecDiff(3, 7),
	)
	assert.ElementsMatch(
		[]int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98},
		numsSameConsecDiff(2, 1),
	)
}
