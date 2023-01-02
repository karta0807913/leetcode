package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Pair struct {
	Left  int
	Right int
}

func isPalindrome(s string, left, right, k int) (p Pair, ok bool) {
	for 0 <= left && right < len(s) {
		if s[left] != s[right] {
			return Pair{}, false
		}
		if right-left+1 >= k {
			return Pair{Left: left, Right: right}, true
		}
		left--
		right++
	}
	return Pair{}, false
}

func maxPalindromes(s string, k int) int {
	end := -1
	result := 0
	if k%2 == 0 {
		for i := range s {
			if p, ok := isPalindrome(s, i, i+1, k); ok {
				if p.Left > end {
					end = p.Right
					result++
				}
				continue
			}
			if p, ok := isPalindrome(s, i, i, k); ok {
				if p.Left > end {
					end = p.Right
					result++
				}
			}
		}
	} else {
		for i := range s {
			if p, ok := isPalindrome(s, i, i, k); ok {
				if p.Left > end {
					end = p.Right
					result++
				}
				continue
			}
			if p, ok := isPalindrome(s, i, i+1, k); ok {
				if p.Left > end {
					end = p.Right
					result++
				}
			}
		}
	}
	return result
}

func TestMax(t *testing.T) {
	assert := require.New(t)
	assert.Equal(12, maxPalindromes("zqzogfurlfmrnlffuipuupidkfhkggkhdrzezghwziopoinnsdkwkymhygonbiizmmmmzjhmyczzlz", 2))
}
