package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func partition(s string) [][]string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	result := [][]string{{}}
	recursive(s, dp, 0, len(s)-1, &result)
	return result
}

func recursive(s string, dp [][]bool, start, end int, result *[][]string) {
	prefix := (*result)[len(*result)-1]
	*result = (*result)[:len(*result)-1]
	for idx := start; idx <= end; idx++ {
		sEnd := idx
		isPrevPalindrome := (start+1 >= sEnd-1 || dp[start+1][sEnd-1] || dp[start][sEnd]) && s[start] == s[sEnd]
		// fmt.Printf("s: %v\n", s[start:sEnd+1])
		// fmt.Printf("isPrevPalindrome: %v\n", isPrevPalindrome)
		if isPrevPalindrome {
			dp[start][sEnd] = true
			if sEnd == end {
				*result = append(*result, append(append([]string{}, prefix...), s[start:sEnd+1]))
			} else {
				*result = append(*result, append(prefix, s[start:sEnd+1]))
				recursive(s, dp, sEnd+1, end, result)
			}
		}
	}
}

func TestPartition(t *testing.T) {
	assert := require.New(t)
	assert.ElementsMatch([][]string{
		{"a", "a", "b"},
		{"aa", "b"},
	}, partition("aab"))
}
