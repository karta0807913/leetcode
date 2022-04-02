package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setTrue(dp []bool) {
	for d := range dp {
		dp[d] = true
	}
}

func isMatch(s string, p string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	prefix := 0
	for i := range p {
		c := p[i]
		switch c {
		case '?':
			for idx := len(s); idx > prefix; idx-- {
				dp[idx] = dp[idx-1]
				dp[idx-1] = false
			}
			prefix += 1
		case '*':
			for idx := prefix; idx < len(s); idx++ {
				if dp[idx] {
					setTrue(dp[idx:])
					prefix = idx
					break
				}
				dp[len(dp)-1] = dp[len(dp)-1] || dp[len(dp)-2]
			}
		default:
			for idx := len(s); idx > prefix; idx-- {
				dp[idx] = dp[idx-1] && s[idx-1] == c
				dp[idx-1] = false
			}
			prefix += 1
		}
		// fmt.Printf("dp: %v, c: %c\n", dp, c)
	}
	return dp[len(dp)-1] && prefix < len(dp)
}

func TestMatch(t *testing.T) {
	assert.False(t, isMatch("a", "aa"))
	assert.False(t, isMatch("ab", "*a"))
	assert.True(t, isMatch("ab", "*a?"))
	assert.False(t, isMatch("aa", "a"))
	assert.True(t, isMatch("aa", "a?"))
	assert.False(t, isMatch("acdcb", "a*c?b"))
	assert.False(t, isMatch("mississippi", "m??*ss*?i*pi"))
	assert.False(t, isMatch("baabba", "?*?a??"))
}
