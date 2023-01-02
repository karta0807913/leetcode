package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func recursive(arr []string, strLen, strMask int) int {
	maxLength := strLen
	for i, str := range arr {
		mask := 0
		for _, char := range str {
			mask |= 1 << (char - 'a')
		}
		if strMask&mask == 0 {
			if length := recursive(
				arr[i+1:],
				strLen+len(str),
				strMask|mask,
			); length > maxLength {
				maxLength = length
			}
		}
	}
	return maxLength
}

func maxLength(arr []string) int {
	left := 0
	for _, str := range arr {
		mask := 0
		keepThis := true
		for _, char := range str {
			c := 1 << (char - 'a')
			keepThis = keepThis && mask&c == 0
			mask |= c
		}
		if keepThis {
			arr[left] = str
			left++
		}
	}
	return recursive(arr[:left], 0, 0)
}

func TestLength(t *testing.T) {
	assert := require.New(t)
	assert.Equal(6, maxLength([]string{"cha", "r", "act", "ers"}))
	assert.Equal(0, maxLength([]string{"aa", "bb"}))
}
