package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func markCan(dp []bool, start int) {
	for ; start < len(dp); start++ {
		dp[start] = true
	}
}

func isMatch(s string, p string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	preLen := 0
	for idx := 0; idx < len(p); idx++ {
		if p[idx] == '.' {
			if idx != len(p)-1 && p[idx+1] == '*' {
				for i := preLen; i < len(dp); i++ {
					if dp[i] {
						preLen = i
						markCan(dp, i)
						break
					}
				}
				idx++
			} else {
				dp[len(dp)-1] = false
				for workIdx := len(dp) - 2; workIdx >= preLen; workIdx-- {
					if dp[workIdx] {
						dp[workIdx] = false
						dp[workIdx+1] = true
					}
				}
				preLen += 1
			}
		} else {
			if idx != len(p)-1 && p[idx+1] == '*' {
				for i := preLen; i < len(dp)-1; i++ {
					if dp[i] && s[i] == p[idx] {
						i++
						for ; i < len(dp)-1; i++ {
							dp[i] = true
							if s[i] != p[idx] {
								break
							}
						}
					}
				}
				// fmt.Printf("dp: %v\n", dp)
				dp[len(dp)-1] = dp[len(dp)-1] || (dp[len(dp)-2] && s[len(dp)-2] == p[idx])
				idx++
			} else {
				dp[len(dp)-1] = false
				for workIdx := len(dp) - 2; workIdx >= preLen; workIdx-- {
					dp[workIdx+1] = dp[workIdx] && s[workIdx] == p[idx]
					dp[workIdx] = false
					// fmt.Printf("dp: %v\n", dp)
				}
				preLen += 1
			}
		}
	}
	return dp[len(dp)-1]
}

func TestRegex(t *testing.T) {
	assert.False(t, isMatch("abcd", "d*"))
	assert.False(t, isMatch("aal", "a*"))
	assert.True(t, isMatch("a", "ab*"))
	assert.True(t, isMatch("aa", "a*"))
	assert.False(t, isMatch("aal", "a*b"))
	assert.False(t, isMatch("aa", "a"))
	assert.False(t, isMatch("aa", "a*c"))
	assert.True(t, isMatch("aac", "a*c"))
	assert.True(t, isMatch("aac", "..."))
	assert.False(t, isMatch("aac", ".."))
	assert.False(t, isMatch("aaaaad", "a*c"))
	assert.True(t, isMatch("aaaaad", "a*c*d"))
	assert.True(t, isMatch("aaaaacd", "a*c*d"))
	assert.True(t, isMatch("aaaaacd", "a*.d"))
	assert.True(t, isMatch("c", "c*c"))
	assert.False(t, isMatch("abc", "c*c"))
	assert.True(t, isMatch("aac", "..*"))
	assert.True(t, isMatch("aac", "..c"))
	assert.False(t, isMatch("bac", "a.*c"))
	assert.True(t, isMatch("aac", "..*c"))
	assert.True(t, isMatch("aac", "a.*c"))
	assert.True(t, isMatch("a", "."))
	assert.True(t, isMatch("a", ".*"))
	assert.True(t, isMatch("", ".*"))
	assert.False(t, isMatch("ippi", "p.*"))
	assert.False(t, isMatch("ssippi", "s*p*."))
	assert.False(t, isMatch("mississippi", "mis*is*p*."))
}
